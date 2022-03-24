package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	arg, _ := strconv.Atoi(args[1])

	studentList := []Student{
		{Nama: "Mubin", Alamat: "Dsn Sidorawuh RT 009 RW 002 Kel. Sidogemah Kec. Sayung Kab. Demak", Pekerjaan: "Staff pada PTIPD UIN Walisongo Semarang", Alasan: "Meningkatkan kompetensi dibidang programming"},
		{Nama: "Aan", Alamat: "Purwokerto, Kab. Purwokerto", Pekerjaan: "Staff pada PTIPD UIN Walisongo Semarang", Alasan: "Meningkatkan kompetensi dibidang programming"},
		{Nama: "Satria", Alamat: "Semarang, Kota Semarang", Pekerjaan: "Staff pada PTIPD UIN Walisongo Semarang", Alasan: "Meningkatkan kompetensi dibidang programming"},
		{Nama: "Teo", Alamat: "Semarang, Kota Semarang", Pekerjaan: "Staff pada PTIPD UIN Walisongo Semarang", Alasan: "Meningkatkan kompetensi dibidang programming"},
		{Nama: "Wahyu", Alamat: "Bulung Kulon Jekulo Kudus", Pekerjaan: "Staff pada PTIPD UIN Walisongo Semarang", Alasan: "Meningkatkan kompetensi dibidang programming"},
	}

	fmt.Println("Peserta Training Golang dengan Nomor Urut : ", arg)
	studentList[arg-1].PrintBiodata()

}

type Student struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (s Student) PrintBiodata() {
	fmt.Println("Nama :", s.Nama)
	fmt.Println("Alamat :", s.Alamat)
	fmt.Println("Pekerjaan :", s.Pekerjaan)
	fmt.Println("Alasan :", s.Alasan)
}
