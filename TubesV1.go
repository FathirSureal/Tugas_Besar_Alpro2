package main

import "fmt"

const MAX = 200

type Produk struct {
	ID       string
	Nama     string
	Kategori string
	Stok     int
}

type Pemasok struct {
	Nama   string
	Kontak string
	Alamat string
}

type Storage struct {
	LokasiRak      string
	UmurSimpanHari int
	TanggalMasuk   string
}

var daftarProduk [MAX]Produk
var daftarStorage [MAX]Storage
var daftarPemasok [MAX]Pemasok
var jumlahProduk int
var jumlahLog int
var logMasuk [100]string

var kategoriList = [10]string{"Buah", "Sayur", "Daging", "Ikan", "Susu", "Roti", "Minuman", "Bumbu", "Frozen", "Snack"}
var lokasiRakList = [10]string{"A1", "A2", "B1", "B2", "C1", "C2", "D1", "D2", "E1", "E2"}

func TampilkanPilihan(label string, daftar [10]string) {
	fmt.Println("Pilihan", label, ":")
	for i, item := range daftar {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func TambahProduk() {
	if jumlahProduk >= MAX {
		fmt.Println("Gudang penuh!")
		return
	}
	var p Produk
	var s Storage
	var ps Pemasok

	fmt.Print("Masukkan ID produk: ")
	fmt.Scan(&p.ID)
	fmt.Print("Masukkan nama produk: ")
	fmt.Scan(&p.Nama)

	TampilkanPilihan("Kategori", kategoriList)
	var pilihanKategori int
	fmt.Scan(&pilihanKategori)
	if pilihanKategori < 1 || pilihanKategori > 10 {
		fmt.Println("Kategori tidak valid.")
		return
	}
	p.Kategori = kategoriList[pilihanKategori-1]

	fmt.Print("Masukkan stok: ")
	fmt.Scan(&p.Stok)

	fmt.Print("Masukkan umur simpan (hari): ")
	fmt.Scan(&s.UmurSimpanHari)

	TampilkanPilihan("Rak Penyimpanan", lokasiRakList)
	var pilihanRak int
	fmt.Scan(&pilihanRak)
	if pilihanRak < 1 || pilihanRak > 10 {
		fmt.Println("Lokasi rak tidak valid.")
		return
	}
	s.LokasiRak = lokasiRakList[pilihanRak-1]

	fmt.Print("Masukkan tanggal masuk (DD-MM-YYYY): ")
	fmt.Scan(&s.TanggalMasuk)

	fmt.Print("Masukkan nama pemasok: ")
	fmt.Scan(&ps.Nama)
	fmt.Print("Masukkan kontak pemasok: ")
	fmt.Scan(&ps.Kontak)
	fmt.Print("Masukkan alamat pemasok: ")
	fmt.Scan(&ps.Alamat)

	daftarProduk[jumlahProduk] = p
	daftarStorage[jumlahProduk] = s
	daftarPemasok[jumlahProduk] = ps

	if jumlahLog < len(logMasuk) {
		logMasuk[jumlahLog] = fmt.Sprintf("Tambah produk %s (%s)", p.Nama, p.ID)
		jumlahLog++
	}

	jumlahProduk++
	fmt.Println("Produk dan pemasok berhasil ditambahkan!")
}

func TampilkanProduk(i int) {
	p := daftarProduk[i]
	s := daftarStorage[i]
	ps := daftarPemasok[i]
	fmt.Printf("\n[%s] %s - %s\n", p.ID, p.Nama, p.Kategori)
	fmt.Printf("Stok         : %d\n", p.Stok)
	fmt.Printf("Umur Simpan  : %d hari\n", s.UmurSimpanHari)
	fmt.Printf("Lokasi Rak   : %s\n", s.LokasiRak)
	fmt.Printf("Tanggal Masuk: %s\n", s.TanggalMasuk)
	fmt.Printf("Pemasok      : %s | %s | %s\n", ps.Nama, ps.Kontak, ps.Alamat)
}
func MenuPencarian() {
	var pilihan int
	for pilihan != 8 {
		fmt.Println("\n--- Menu Pencarian Produk ---")
		fmt.Println("1. Cari by ID (Sequential Search)")
		fmt.Println("2. Cari by Nama (Binary Search)")
		fmt.Println("3. Cari by Prefix Nama (Binary Search)")
		fmt.Println("4. Cari by Nama Pemasok")
		fmt.Println("5. Cari by Kategori")
		fmt.Println("6. Cari by Tanggal Masuk")
		fmt.Println("7. Cari by Lokasi Rak")
		fmt.Println("8. Kembali ke Menu Utama")
		fmt.Print("Pilih jenis pencarian: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id string
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			if idx := SeqSearchCariProdukByID(id); idx != -1 {
				TampilkanProduk(idx)
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
		case 2:
			InsertionSortByNama(true)
			var nama string
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&nama)
			if idx := BinarySearchByNama(nama); idx != -1 {
				TampilkanProduk(idx)
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
		case 3:
			fmt.Print("Masukkan awalan nama: ")
			var pref string
			fmt.Scan(&pref)
			ids := BinarySearchByPrefixNama(pref)
			if len(ids) == 0 {
				fmt.Println("Tidak ditemukan.")
			} else {
				for _, i := range ids {
					TampilkanProduk(i)
				}
			}
		case 4:
			CariProdukByPemasok()
		case 5:
			CariProdukByKategori()
		case 6:
			CariProdukByTanggalMasuk()
		case 7:
			CariProdukByLokasiRak()
		case 8:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func SeqSearchCariProdukByID(id string) int {
	for i := 0; i < jumlahProduk; i++ {
		if daftarProduk[i].ID == id {
			return i
		}
	}
	return -1
}

func InsertionSortByNama(asc bool) {
	for i := 1; i < jumlahProduk; i++ {
		keyP := daftarProduk[i]
		keyS := daftarStorage[i]
		j := i - 1
		for j >= 0 && ((asc && daftarProduk[j].Nama > keyP.Nama) || (!asc && daftarProduk[j].Nama < keyP.Nama)) {
			daftarProduk[j+1] = daftarProduk[j]
			daftarStorage[j+1] = daftarStorage[j]
			j--
		}
		daftarProduk[j+1] = keyP
		daftarStorage[j+1] = keyS
	}
}

func BinarySearchByNama(nama string) int {
	InsertionSortByNama(true)
	low, high := 0, jumlahProduk-1
	for low <= high {
		mid := (low + high) / 2
		if daftarProduk[mid].Nama == nama {
			return mid
		} else if daftarProduk[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchByPrefixNama(prefix string) []int {
	InsertionSortByNama(true)
	hasilIdx := []int{}
	low, high := 0, jumlahProduk-1
	var mid int
	found := false
	for low <= high && !found {
		mid = (low + high) / 2
		nm := daftarProduk[mid].Nama
		if len(nm) >= len(prefix) && nm[:len(prefix)] == prefix {
			found = true
		} else if nm < prefix {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		return hasilIdx
	}
	i := mid
	for i >= 0 && len(daftarProduk[i].Nama) >= len(prefix) && daftarProduk[i].Nama[:len(prefix)] == prefix {
		i--
	}
	i++
	for i < jumlahProduk && len(daftarProduk[i].Nama) >= len(prefix) && daftarProduk[i].Nama[:len(prefix)] == prefix {
		hasilIdx = append(hasilIdx, i)
		i++
	}
	return hasilIdx
}

func SelectionSortByUmur(asc bool) {
	for i := 0; i < jumlahProduk-1; i++ {
		idx := i
		for j := i + 1; j < jumlahProduk; j++ {
			if (asc && daftarStorage[j].UmurSimpanHari < daftarStorage[idx].UmurSimpanHari) ||
				(!asc && daftarStorage[j].UmurSimpanHari > daftarStorage[idx].UmurSimpanHari) {
				idx = j
			}
		}
		daftarProduk[i], daftarProduk[idx] = daftarProduk[idx], daftarProduk[i]
		daftarStorage[i], daftarStorage[idx] = daftarStorage[idx], daftarStorage[i]
	}
}

func InsertionSortByKategori(asc bool) {
	for i := 1; i < jumlahProduk; i++ {
		keyP := daftarProduk[i]
		keyS := daftarStorage[i]
		j := i - 1
		for j >= 0 && ((asc && daftarProduk[j].Kategori > keyP.Kategori) || (!asc && daftarProduk[j].Kategori < keyP.Kategori)) {
			daftarProduk[j+1] = daftarProduk[j]
			daftarStorage[j+1] = daftarStorage[j]
			j--
		}
		daftarProduk[j+1] = keyP
		daftarStorage[j+1] = keyS
	}
}
func InsertionSortByID(asc bool) {
	for i := 1; i < jumlahProduk; i++ {
		keyP := daftarProduk[i]
		keyS := daftarStorage[i]
		keyPS := daftarPemasok[i]
		j := i - 1
		for j >= 0 && ((asc && daftarProduk[j].ID > keyP.ID) || (!asc && daftarProduk[j].ID < keyP.ID)) {
			daftarProduk[j+1] = daftarProduk[j]
			daftarStorage[j+1] = daftarStorage[j]
			daftarPemasok[j+1] = daftarPemasok[j]
			j--
		}
		daftarProduk[j+1] = keyP
		daftarStorage[j+1] = keyS
		daftarPemasok[j+1] = keyPS
	}
}

func InsertionSortByRak(asc bool) {
	for i := 1; i < jumlahProduk; i++ {
		keyP := daftarProduk[i]
		keyS := daftarStorage[i]
		keyPS := daftarPemasok[i]
		j := i - 1
		for j >= 0 && ((asc && daftarStorage[j].LokasiRak > keyS.LokasiRak) || (!asc && daftarStorage[j].LokasiRak < keyS.LokasiRak)) {
			daftarProduk[j+1] = daftarProduk[j]
			daftarStorage[j+1] = daftarStorage[j]
			daftarPemasok[j+1] = daftarPemasok[j]
			j--
		}
		daftarProduk[j+1] = keyP
		daftarStorage[j+1] = keyS
		daftarPemasok[j+1] = keyPS
	}
}

func InsertionSortByTanggalMasuk(asc bool) {
	for i := 1; i < jumlahProduk; i++ {
		keyP := daftarProduk[i]
		keyS := daftarStorage[i]
		keyPS := daftarPemasok[i]
		j := i - 1
		for j >= 0 && ((asc && daftarStorage[j].TanggalMasuk > keyS.TanggalMasuk) || (!asc && daftarStorage[j].TanggalMasuk < keyS.TanggalMasuk)) {
			daftarProduk[j+1] = daftarProduk[j]
			daftarStorage[j+1] = daftarStorage[j]
			daftarPemasok[j+1] = daftarPemasok[j]
			j--
		}
		daftarProduk[j+1] = keyP
		daftarStorage[j+1] = keyS
		daftarPemasok[j+1] = keyPS
	}
}

func CetakLog() {
	fmt.Println("Riwayat Transaksi:")
	for i := 0; i < jumlahLog; i++ {
		fmt.Printf("%d. %s\n", i+1, logMasuk[i])
	}
}

func EditProduk() {
	var id string
	fmt.Print("Masukkan ID produk yang ingin diedit: ")
	fmt.Scan(&id)
	idx := SeqSearchCariProdukByID(id)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}

	for {
		fmt.Println("\nEdit Data Produk:")
		fmt.Println("1. Nama Produk")
		fmt.Println("2. Kategori")
		fmt.Println("3. Stok")
		fmt.Println("4. Umur Simpan")
		fmt.Println("5. Lokasi Rak")
		fmt.Println("6. Tanggal Masuk")
		fmt.Println("7. Nama Pemasok")
		fmt.Println("8. Kontak Pemasok")
		fmt.Println("9. Alamat Pemasok")
		fmt.Println("0. Selesai Edit")
		fmt.Print("Pilih yang ingin diedit: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Nama baru: ")
			fmt.Scan(&daftarProduk[idx].Nama)
		case 2:
			TampilkanPilihan("Kategori", kategoriList)
			var pilihanKategori int
			fmt.Scan(&pilihanKategori)
			if pilihanKategori >= 1 && pilihanKategori <= 10 {
				daftarProduk[idx].Kategori = kategoriList[pilihanKategori-1]
			}
		case 3:
			fmt.Print("Stok baru: ")
			fmt.Scan(&daftarProduk[idx].Stok)
		case 4:
			fmt.Print("Umur simpan baru: ")
			fmt.Scan(&daftarStorage[idx].UmurSimpanHari)
		case 5:
			TampilkanPilihan("Rak", lokasiRakList)
			var pilihanRak int
			fmt.Scan(&pilihanRak)
			if pilihanRak >= 1 && pilihanRak <= 10 {
				daftarStorage[idx].LokasiRak = lokasiRakList[pilihanRak-1]
			}
		case 6:
			fmt.Print("Tanggal masuk baru (DD-MM-YYYY): ")
			fmt.Scan(&daftarStorage[idx].TanggalMasuk)
		case 7:
			fmt.Print("Nama pemasok baru: ")
			fmt.Scan(&daftarPemasok[idx].Nama)
		case 8:
			fmt.Print("Kontak pemasok baru: ")
			fmt.Scan(&daftarPemasok[idx].Kontak)
		case 9:
			fmt.Print("Alamat pemasok baru: ")
			fmt.Scan(&daftarPemasok[idx].Alamat)
		case 0:
			fmt.Println("Edit selesai.")
			logMasuk[jumlahLog] = fmt.Sprintf("Edit produk %s (%s)", daftarProduk[idx].Nama, id)
			jumlahLog++
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func HapusProduk() {
	var id string
	fmt.Print("Masukkan ID produk yang ingin dihapus: ")
	fmt.Scan(&id)
	idx := SeqSearchCariProdukByID(id)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahProduk-1; i++ {
		daftarProduk[i] = daftarProduk[i+1]
		daftarStorage[i] = daftarStorage[i+1]
	}
	jumlahProduk--
	logMasuk[jumlahLog] = fmt.Sprintf("Hapus produk %s (%s)", id, id)
	jumlahLog++
	fmt.Println("Produk berhasil dihapus.")
}

func IsiDataDummy() {
	for i := 1; i <= 40; i++ {
		p := Produk{
			ID:       fmt.Sprintf("P%03d", i),
			Nama:     fmt.Sprintf("Produk%02d", i),
			Kategori: kategoriList[i%len(kategoriList)],
			Stok:     10 + i,
		}
		s := Storage{
			LokasiRak:      lokasiRakList[i%len(lokasiRakList)],
			UmurSimpanHari: 5 + (i % 10),
			TanggalMasuk:   fmt.Sprintf("%02d-06-2025", (i%30)+1),
		}
		ps := Pemasok{
			Nama:   fmt.Sprintf("Pemasok%02d", (i%10)+1),
			Kontak: fmt.Sprintf("081234567%02d", i),
			Alamat: fmt.Sprintf("Jalan Pemasok No.%d", i),
		}

		daftarProduk[jumlahProduk] = p
		daftarStorage[jumlahProduk] = s
		daftarPemasok[jumlahProduk] = ps

		logMasuk[jumlahLog] = fmt.Sprintf("Tambah produk %s (%s)", p.Nama, p.ID)
		jumlahLog++
		jumlahProduk++
	}
	fmt.Println("40 produk dummy beserta pemasok berhasil dimuat.")
}

func CariProdukByPemasok() {
	var namaPemasok string
	fmt.Print("Masukkan nama pemasok: ")
	fmt.Scan(&namaPemasok)

	ditemukan := false
	for i := 0; i < jumlahProduk; i++ {
		if daftarPemasok[i].Nama == namaPemasok {
			TampilkanProduk(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada produk dari pemasok tersebut.")
	}
}

func CariProdukByKategori() {
	TampilkanPilihan("Kategori", kategoriList)
	var pilihan int
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&pilihan)
	if pilihan < 1 || pilihan > len(kategoriList) {
		fmt.Println("Kategori tidak valid.")
		return
	}
	kategori := kategoriList[pilihan-1]
	ditemukan := false
	for i := 0; i < jumlahProduk; i++ {
		if daftarProduk[i].Kategori == kategori {
			TampilkanProduk(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada produk dalam kategori ini.")
	}
}

func CariProdukByTanggalMasuk() {
	var tanggal string
	fmt.Print("Masukkan tanggal masuk (DD-MM-YYYY): ")
	fmt.Scan(&tanggal)

	ditemukan := false
	for i := 0; i < jumlahProduk; i++ {
		if daftarStorage[i].TanggalMasuk == tanggal {
			TampilkanProduk(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada produk dengan tanggal masuk tersebut.")
	}
}

func CariProdukByLokasiRak() {
	TampilkanPilihan("Rak Penyimpanan", lokasiRakList)
	var pilihan int
	fmt.Print("Pilih rak: ")
	fmt.Scan(&pilihan)
	if pilihan < 1 || pilihan > len(lokasiRakList) {
		fmt.Println("Lokasi rak tidak valid.")
		return
	}
	rak := lokasiRakList[pilihan-1]
	ditemukan := false
	for i := 0; i < jumlahProduk; i++ {
		if daftarStorage[i].LokasiRak == rak {
			TampilkanProduk(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada produk di rak tersebut.")
	}
}

func main() {
	var pilihan int
	for pilihan != 16 {
		fmt.Println("\n--- Menu Gudang FreshMart ---")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Tampilkan Semua Produk")
		fmt.Println("3. Menu Search")
		fmt.Println("4. Urutkan Umur Simpan")
		fmt.Println("5. Urutkan Kategori")
		fmt.Println("6. Cetak Log")
		fmt.Println("7. Edit Produk")
		fmt.Println("8. Hapus Produk")
		fmt.Println("9. Muat Data Dummy (40 Produk)")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahProduk()
		case 2:
			if jumlahProduk == 0 {
				fmt.Println("Belum ada produk yang tersedia di gudang.")
			} else {
				for i := 0; i < jumlahProduk; i++ {
					TampilkanProduk(i)
				}
			}
		case 3:
			MenuPencarian()
		case 4:
			fmt.Println("\n--- Menu Pengurutan Produk ---")
			fmt.Println("1. Berdasarkan ID")
			fmt.Println("2. Berdasarkan Nama Produk")
			fmt.Println("3. Berdasarkan Kategori")
			fmt.Println("4. Berdasarkan Lokasi Rak")
			fmt.Println("5. Berdasarkan Umur Simpan")
			fmt.Println("6. Berdasarkan Tanggal Masuk")
			fmt.Print("Pilih metode pengurutan: ")
			var metode int
			fmt.Scan(&metode)
			fmt.Print("1 = Ascending, 2 = Descending: ")
			var arah int
			fmt.Scan(&arah)
			asc := arah == 1

			switch metode {
			case 1:
				InsertionSortByID(asc)
				fmt.Println("Produk diurutkan berdasarkan ID.")
			case 2:
				InsertionSortByNama(asc)
				fmt.Println("Produk diurutkan berdasarkan Nama.")
			case 3:
				InsertionSortByKategori(asc)
				fmt.Println("Produk diurutkan berdasarkan Kategori.")
			case 4:
				InsertionSortByRak(asc)
				fmt.Println("Produk diurutkan berdasarkan Lokasi Rak.")
			case 5:
				SelectionSortByUmur(asc)
				fmt.Println("Produk diurutkan berdasarkan Umur Simpan.")
			case 6:
				InsertionSortByTanggalMasuk(asc)
				fmt.Println("Produk diurutkan berdasarkan Tanggal Masuk.")
			default:
				fmt.Println("Pilihan tidak valid.")
			}

		case 5:
			CetakLog()
		case 6:
			EditProduk()
		case 7:
			HapusProduk()
		case 8:
			IsiDataDummy()
		case 9:
			fmt.Println("Terima kasih telah menggunakan sistem.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
