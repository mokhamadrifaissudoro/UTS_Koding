package main

import (
	"fmt"
	"os"
)

type Buku struct {
	Nama   string
	Jumlah int
}

type Peminjaman struct {
	Username string
	Buku     string
	Jumlah   int
}

var bukuList = []Buku{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
}

var peminjamanList []Peminjaman

// Username dan password mahasiswa
var userDB = map[string]string{
	"Rifais": "2406359140",
}

// Fungsi utama
func main() {
	if !login() {
		return
	}
	for {
		fmt.Println("\n--- MENU PERPUSTAKAAN ---")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar Dari Program")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			LihatInformasiPenggunaProgram()
		case 2:
			LihatDaftarBuku()
		case 3:
			TambahDaftarBuku()
		case 4:
			TambahPeminjamanBuku()
		case 5:
			HistoriPeminjamanBuku()
		case 6:
			KeluarDariProgram()
		default:
			fmt.Println("===== Pilihan tidak valid, coba lagi! =====")
		}
	}
}

// Fungsi untuk login
func login() bool {
	var username, password string
	fmt.Println("===== SELAMAT DATANG DI PERPUSTAKAAN VOKASI =====")
	fmt.Print("Silahkan Input Username: ")
	fmt.Scanln(&username)

	fmt.Print("Silahkan Input Password: ")
	fmt.Scanln(&password)

	if pass, exists := userDB[username]; exists && pass == password {
		fmt.Println("===== LOGIN SUKSES! =====")
		return true
	} else {
		fmt.Println("===== LOGIN GAGAL! USERNAME ATAU PASSWORD SALAH! =====")
		return false
	}
}

// Fungsi untuk melihat informasi pengguna
func LihatInformasiPenggunaProgram() {
	fmt.Println("===== INFORMASI PENGGUNA PROGRAM =====")
	fmt.Printf("Username: %s\n", "Rifais")
	fmt.Printf("NPM: %s\n", "2406359140")
	fmt.Printf("TTL: %s\n", "Jakarta, 20 Juni 2006")
	fmt.Printf("Jenis Kelamin: %s\n", "Laki-laki")
	fmt.Printf("MBTI: %s\n", "ESFP")
	fmt.Printf("Hobby: %s\n", "Menyanyi")
	fmt.Printf("Makanan Favorit Favorit: %s\n", "Cumi Goreng Tepung")
	fmt.Printf("Minuman Favorit Favorit: %s\n", "Thai Tea")
	fmt.Printf("Artis Favorit: %s\n", "Giselle Of AESPA")
	fmt.Printf("Lagu Favorit: %s\n", "Supernova by AESPA")
}

// Fungsi untuk melihat daftar buku
func LihatDaftarBuku() {
	fmt.Println("===== DAFTAR BUKU =====")
	for _, buku := range bukuList {
		fmt.Printf("Nama Buku: %s, Jumlah: %d\n", buku.Nama, buku.Jumlah)
	}
}

// Fungsi untuk menambah daftar buku
func TambahDaftarBuku() {
	var namaBuku string
	var jumlah int

	fmt.Println("===== TAMBAH DAFTAR BUKU =====")
	fmt.Print("Masukkan Nama Buku: ")
	fmt.Scanln(&namaBuku)

	fmt.Print("Masukkan Jumlah Buku: ")
	fmt.Scanln(&jumlah)

	bukuList = append(bukuList, Buku{Nama: namaBuku, Jumlah: jumlah})
	fmt.Println("===== BUKU BERHASIL DITAMBAHKAN! =====")
}

// Fungsi untuk menambah peminjaman buku
func TambahPeminjamanBuku() {
	var username string
	var jumlah int

	fmt.Println("===== TAMBAH PEMINJAMAN BUKU =====")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&username)

	// Tampilkan daftar buku sebelum meminta input nama buku
	fmt.Println("\nDaftar Buku Tersedia untuk Dipinjam:")
	for i, buku := range bukuList {
		fmt.Printf("%d. %s (%d)\n", i+1, buku.Nama, buku.Jumlah)
	}

	var nomorBuku int
	fmt.Print("Masukkan Nomor Buku yang Ingin Dipinjam: ")
	fmt.Scanln(&nomorBuku)

	if nomorBuku < 1 || nomorBuku > len(bukuList) {
		fmt.Println("===== NOMOR BUKU TIDAK VALID! =====")
		return
	}

	namaBuku := bukuList[nomorBuku-1].Nama // Ambil nama buku dari daftar
	fmt.Print("Masukkan Jumlah Buku yang Dipinjam: ")
	fmt.Scanln(&jumlah)

	// Validasi jumlah buku
	if jumlah <= 0 || jumlah > bukuList[nomorBuku-1].Jumlah {
		fmt.Println("===== JUMLAH BUKU TIDAK VALID ATAU MELEBIHI STOK YANG ADA! =====")
		return
	}

	// Kurangi jumlah buku
	bukuList[nomorBuku-1].Jumlah -= jumlah
	peminjamanList = append(peminjamanList, Peminjaman{Username: username, Buku: namaBuku, Jumlah: jumlah})
	fmt.Println("===== PEMINJAMAN BUKU BERHASIL! =====")
}

// Fungsi untuk melihat histori peminjaman buku
func HistoriPeminjamanBuku() {
	fmt.Println("===== HISTORI PEMINJAMAN BUKU =====")
	if len(peminjamanList) == 0 {
		fmt.Println("===== BELUM ADA HISTORI PEMINJAMAN BUKU =====")
	} else {
		for _, peminjaman := range peminjamanList {
			fmt.Printf("Username: %s, Buku: %s, Jumlah: %d\n", peminjaman.Username, peminjaman.Buku, peminjaman.Jumlah)
		}
	}
}

// Fungsi untuk keluar dari program
func KeluarDariProgram() {
	fmt.Println("===== TERIMA KASIH TELAH BERKUNJUNG! YUK RAJIN MEMBACA! =====")
	os.Exit(0)
}
