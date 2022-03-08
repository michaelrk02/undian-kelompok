# Sistem Undian Kelompok

Berbasis Golang

## Cara Compile

```
$ go build -o build/undian
```

## Cara Penggunaan

```
$ ./undian -b
  -a    bagi berdasarkan jumlah anggota kelompok
  -b    tampilkan bantuan
  -i string
        file masukan berupa CSV (default "-")
  -j int
        jumlah kelompok atau anggota kelompok (wajib diisi)
  -o string
        file keluaran berupa CSV (default "-")
```

```
sample-in.csv
-------------
NIM,Nama
M0520001,Bruce Dickinson
M0520002,Adrian Smith
M0520003,Dave Murray
M0520004,Steve Harris
M0520005,Clive Burr
M0520006,Tony Iommi
M0520007,Geezer Butler
M0520008,Ozzy Osbourne
M0520009,Bill Ward
M0520010,Neil Peart
M0520011,John Petrucci
M0520012,John Myung
M0520013,James LaBrie
M0520014,Jordan Rudess
M0520015,Mike Portnoy
```

```
$ ./undian -i sample-in.csv -j 4 -a -o sample-out.csv
```

```
sample-out.csv
--------------
Kelompok,NIM,Nama
1,M0520010,Neil Peart
1,M0520007,Geezer Butler
1,M0520011,John Petrucci
1,M0520005,Clive Burr
2,M0520014,Jordan Rudess
2,M0520001,Bruce Dickinson
2,M0520013,James LaBrie
2,M0520003,Dave Murray
3,M0520015,Mike Portnoy
3,M0520009,Bill Ward
3,M0520012,John Myung
3,M0520006,Tony Iommi
4,M0520004,Steve Harris
4,M0520008,Ozzy Osbourne
4,M0520002,Adrian Smith
```
