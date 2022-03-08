package main

import (
    "container/list"
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "math/rand"
    "os"
    "time"
)

type Mahasiswa struct {
    NIM string
    Nama string
}

func RandomPick(in *list.List, out *list.List) {
    var walk int = 0
    var index int = rand.Intn(in.Len())
    for iter := in.Front(); iter != nil; iter = iter.Next() {
        if walk == index {
            out.PushBack(iter.Value)
            in.Remove(iter)
            break
        }
        walk++
    }
}

func main() {
    var err error

    rand.Seed(time.Now().UnixNano())

    var input, output string
    var jumlah int
    var anggota, bantuan bool

    flag.StringVar(&input, "i", "-", "file masukan berupa CSV")
    flag.StringVar(&output, "o", "-", "file keluaran berupa CSV")
    flag.IntVar(&jumlah, "j", 0, "jumlah kelompok atau anggota kelompok (wajib diisi)")
    flag.BoolVar(&anggota, "a", false, "bagi berdasarkan jumlah anggota kelompok")
    flag.BoolVar(&bantuan, "b", false, "tampilkan bantuan")

    flag.Parse()

    if bantuan {
        flag.PrintDefaults()
        return
    }

    if jumlah == 0 {
        fmt.Fprintln(os.Stderr, "Jumlah wajib diisi dan tidak boleh nol. Sisipkan -b untuk melihat bantuan")
        return
    }

    var inputFile, outputFile *os.File

    if input == "-" {
        inputFile = os.Stdin
    } else {
        if inputFile, err = os.Open(input); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
            return
        }
        defer inputFile.Close()
    }

    if output == "-" {
        outputFile = os.Stdout
    } else {
        if outputFile, err = os.Create(output); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
            return
        }
        defer outputFile.Close()
    }

    var csvIn *csv.Reader = csv.NewReader(inputFile)

    var line []string
    line, err = csvIn.Read()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
        return
    }

    var indexMap map[string]int = make(map[string]int)
    for i := range line {
        if line[i] == "NIM" {
            indexMap["NIM"] = i
        } else if line[i] == "Nama" {
            indexMap["Nama"] = i
        }
    }

    var mhsList *list.List = list.New()

    for true {
        line, err = csvIn.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
            return
        }

        var mhs *Mahasiswa = new(Mahasiswa)
        mhs.NIM = line[indexMap["NIM"]]
        mhs.Nama = line[indexMap["Nama"]]

        mhsList.PushBack(mhs)
    }

    var jKel int
    if anggota {
        jKel = mhsList.Len() / jumlah
        if mhsList.Len() % jumlah > 0 {
            jKel++
        }
    } else {
        jKel = jumlah
    }

    var kelArr []*list.List = make([]*list.List, jKel)
    for i := range kelArr {
        kelArr[i] = list.New()
    }

    var iKel int = 0
    for mhsList.Len() > 0 {
        RandomPick(mhsList, kelArr[iKel])
        iKel = (iKel + 1) % jKel
    }

    var csvOut *csv.Writer = csv.NewWriter(outputFile)
    csvOut.Write([]string{"Kelompok", "NIM", "Nama"})
    for i := 0; i < jKel; i++ {
        for iter := kelArr[i].Front(); iter != nil; iter = iter.Next() {
            var mhs = iter.Value.(*Mahasiswa)
            csvOut.Write([]string{fmt.Sprintf("%d", i + 1), mhs.NIM, mhs.Nama})
        }
    }
    csvOut.Flush()
}

