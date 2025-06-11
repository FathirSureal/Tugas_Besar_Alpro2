package main

import "fmt"

const MAX = 200

type Produk struct {
	ID             string
	Nama           string
	Kategori       string
	Stok           int
	UmurSimpanHari int
	LokasiRak      string
	TanggalMasuk   string
}

var daftarProduk [MAX]Produk
var jumlahProduk int = 0
var kategoriList = [10]string{"Buah", "Sayur", "Daging", "Ikan", "Susu", "Roti", "Minuman", "Bumbu", "Frozen", "Snack"}
var logMasuk [100]string
var jumlahLog int = 0
var lokasiRakList = [10]string{"A1", "A2", "B1", "B2", "C1", "C2", "D1", "D2", "E1", "E2"}
var tanggalMasukList [MAX]string
var namaProdukList [MAX]string

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
	fmt.Print("Masukkan ID produk: ")
	fmt.Scan(&p.ID)
	fmt.Print("Masukkan nama produk: ")
	fmt.Scan(&p.Nama)
	TampilkanPilihan("Kategori", kategoriList)
	var pilihanKategori int
	fmt.Print("Masukkan nomor kategori: ")
	fmt.Scan(&pilihanKategori)
	switch pilihanKategori {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		p.Kategori = kategoriList[pilihanKategori-1]
	default:
		fmt.Println("Kategori tidak valid. Produk batal ditambahkan.")
		return
	}
	fmt.Print("Masukkan stok: ")
	fmt.Scan(&p.Stok)
	fmt.Print("Masukkan umur simpan (hari): ")
	fmt.Scan(&p.UmurSimpanHari)
	TampilkanPilihan("Rak Penyimpanan", lokasiRakList)
	var pilihanRak int
	fmt.Print("Masukkan nomor rak: ")
	fmt.Scan(&pilihanRak)

	switch pilihanRak {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		p.LokasiRak = lokasiRakList[pilihanRak-1]
	default:
		fmt.Println("Lokasi rak tidak valid. Produk batal ditambahkan.")
		return
	}
	fmt.Print("Masukkan tanggal masuk (DD-MM-YYYY): ")
	fmt.Scan(&p.TanggalMasuk)
	daftarProduk[jumlahProduk] = p
	tanggalMasukList[jumlahProduk] = p.TanggalMasuk
	namaProdukList[jumlahProduk] = p.Nama
	jumlahProduk++
	if jumlahLog < 100 {
		logMasuk[jumlahLog] = "Tambah produk " + p.Nama + " (" + p.ID + ")"
		jumlahLog++
	}
	fmt.Println("Produk berhasil ditambahkan!")
}

func TampilkanProduk() {
	if jumlahProduk == 0 {
		fmt.Println("Belum ada produk.")
		return
	}
	fmt.Println("Daftar Produk:")
	for i := 0; i < jumlahProduk; i++ {
		p := daftarProduk[i]
		fmt.Printf("%d. [%s] %s - %s | Stok: %d | Umur Simpan: %d hari | Rak: %s | Masuk: %s\n", i+1, p.ID, p.Nama, p.Kategori, p.Stok, p.UmurSimpanHari, p.LokasiRak, p.TanggalMasuk)
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

func BinarySearchByNama(nama string) int {
	kiri := 0
	kanan := jumlahProduk - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarProduk[tengah].Nama == nama {
			return tengah
		} else if daftarProduk[tengah].Nama < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}
func BinarySearchByPrefixNama(prefix string) []Produk {
	InsertionSortByNama(true)
	hasil := []Produk{}
	kiri := 0
	kanan := jumlahProduk - 1
	tengah := -1
	for kiri <= kanan && tengah == -1 {
		mid := (kiri + kanan) / 2
		nama := daftarProduk[mid].Nama
		if len(nama) >= len(prefix) && nama[:len(prefix)] == prefix {
			tengah = mid
		}
		if nama < prefix {
			kiri = mid + 1
		} else if nama > prefix {
			kanan = mid - 1
		}
	}
	if tengah == -1 {
		return hasil
	}
	i := tengah
	for i >= 0 {
		nama := daftarProduk[i].Nama
		if len(nama) >= len(prefix) && nama[:len(prefix)] == prefix {
			i--
		} else {
			i++
			break
		}
	}
	for i < jumlahProduk {
		nama := daftarProduk[i].Nama
		if len(nama) >= len(prefix) && nama[:len(prefix)] == prefix {
			hasil = append(hasil, daftarProduk[i])
			i++
		} else {
			i = jumlahProduk
		}
	}
	return hasil
}

func InsertionSortByNama(ascending bool) {
	for i := 1; i < jumlahProduk; i++ {
		key := daftarProduk[i]
		j := i - 1
		for j >= 0 && ((ascending && daftarProduk[j].Nama > key.Nama) || (!ascending && daftarProduk[j].Nama < key.Nama)) {
			daftarProduk[j+1] = daftarProduk[j]
			j--
		}
		daftarProduk[j+1] = key
	}
}

func SelectionSortUrutkanByUmurSimpan(ascending bool) {
	for i := 0; i < jumlahProduk-1; i++ {
		idx := i
		for j := i + 1; j < jumlahProduk; j++ {
			if (ascending && daftarProduk[j].UmurSimpanHari < daftarProduk[idx].UmurSimpanHari) ||
				(!ascending && daftarProduk[j].UmurSimpanHari > daftarProduk[idx].UmurSimpanHari) {
				idx = j
			}
		}
		if idx != i {
			daftarProduk[i], daftarProduk[idx] = daftarProduk[idx], daftarProduk[i]
		}
	}
}

func InsertionSortUrutkanByKategori(ascending bool) {
	for i := 1; i < jumlahProduk; i++ {
		key := daftarProduk[i]
		j := i - 1
		for j >= 0 && ((ascending && daftarProduk[j].Kategori > key.Kategori) || (!ascending && daftarProduk[j].Kategori < key.Kategori)) {
			daftarProduk[j+1] = daftarProduk[j]
			j--
		}
		daftarProduk[j+1] = key
	}
}

func CetakLog() {
	fmt.Println("Riwayat Transaksi:")
	for i := 0; i < jumlahLog; i++ {
		fmt.Printf("%d. %s\n", i+1, logMasuk[i])
	}
}

func EditStokProduk() {
	var id string
	fmt.Print("Masukkan ID produk yang ingin diubah stoknya: ")
	fmt.Scan(&id)
	idx := SeqSearchCariProdukByID(id)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Printf("Stok lama: %d\n", daftarProduk[idx].Stok)
	fmt.Print("Masukkan stok baru: ")
	fmt.Scan(&daftarProduk[idx].Stok)
	fmt.Println("Stok berhasil diubah.")
}

func EditDataProduk() {
	var id string
	fmt.Print("Masukkan ID produk yang ingin diedit: ")
	fmt.Scan(&id)
	idx := SeqSearchCariProdukByID(id)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Println("Data saat ini:")
	fmt.Printf("ID: %s\n", daftarProduk[idx].ID)
	fmt.Printf("Nama: %s\n", daftarProduk[idx].Nama)
	fmt.Printf("Kategori: %s\n", daftarProduk[idx].Kategori)

	fmt.Print("Masukkan ID baru (tekan '-' untuk lewati): ")
	var newID string
	fmt.Scan(&newID)
	if newID != "-" {
		daftarProduk[idx].ID = newID
	}
	fmt.Print("Masukkan Nama baru (tekan '-' untuk lewati): ")
	var newNama string
	fmt.Scan(&newNama)
	if newNama != "-" {
		daftarProduk[idx].Nama = newNama
		namaProdukList[idx] = newNama
	}
	TampilkanPilihan("Kategori", kategoriList)
	fmt.Print("Masukkan Kategori baru (tekan '-' untuk lewati): ")
	var newKategori string
	fmt.Scan(&newKategori)
	if newKategori != "-" {
		daftarProduk[idx].Kategori = newKategori
	}
	fmt.Println("Data produk berhasil diperbarui.")
}

func IsiDataDummy() {
	data := []Produk{
		{"P001", "Apel_Fuji", "Buah", 100, 14, "A1", "01-06-2025"},
		{"P002", "Bayam_Segar", "Sayur", 50, 5, "A2", "02-06-2025"},
		{"P003", "Daging_Sapi", "Daging", 30, 7, "B1", "03-06-2025"},
		{"P004", "Ikan_Salmon", "Ikan", 40, 6, "B2", "03-06-2025"},
		{"P005", "Susu_UHT", "Susu", 80, 30, "C1", "04-06-2025"},
		{"P006", "Roti_Tawar", "Roti", 60, 3, "C2", "04-06-2025"},
		{"P007", "Teh_Botol", "Minuman", 120, 60, "D1", "05-06-2025"},
		{"P008", "Kecap_Manis", "Bumbu", 45, 180, "D2", "05-06-2025"},
		{"P009", "Nugget_Ayam", "Frozen", 55, 90, "E1", "06-06-2025"},
		{"P010", "Keripik_Singkong", "Snack", 70, 150, "E2", "06-06-2025"},
		{"P011", "Jeruk_Mandarin", "Buah", 90, 10, "A1", "07-06-2025"},
		{"P012", "Wortel_Organik", "Sayur", 65, 8, "A2", "07-06-2025"},
		{"P013", "Daging_Kambing", "Daging", 20, 5, "B1", "08-06-2025"},
		{"P014", "Ikan_Tuna", "Ikan", 35, 6, "B2", "08-06-2025"},
		{"P015", "Susu_Kedelai", "Susu", 40, 20, "C1", "09-06-2025"},
		{"P016", "Roti_Gandum", "Roti", 30, 4, "C2", "09-06-2025"},
		{"P017", "Jus_Apel", "Minuman", 100, 14, "D1", "10-06-2025"},
		{"P018", "Saus_Tomat", "Bumbu", 38, 120, "D2", "10-06-2025"},
		{"P019", "Sosis_Sapi", "Frozen", 25, 60, "E1", "11-06-2025"},
		{"P020", "Cokelat_Batang", "Snack", 95, 180, "E2", "11-06-2025"},
		{"P021", "Pisang_Cavendish", "Buah", 70, 7, "A1", "12-06-2025"},
		{"P022", "Kangkung", "Sayur", 60, 4, "A2", "12-06-2025"},
		{"P023", "Daging_Cincang", "Daging", 28, 5, "B1", "13-06-2025"},
		{"P024", "Ikan_Nila", "Ikan", 50, 4, "B2", "13-06-2025"},
		{"P025", "Susu_Skimming", "Susu", 55, 25, "C1", "14-06-2025"},
		{"P026", "Roti_Sobek", "Roti", 42, 2, "C2", "14-06-2025"},
		{"P027", "Air_Mineral", "Minuman", 150, 365, "D1", "15-06-2025"},
		{"P028", "Garam_Dapur", "Bumbu", 80, 730, "D2", "15-06-2025"},
		{"P029", "Dimsum_Ayam", "Frozen", 60, 45, "E1", "16-06-2025"},
		{"P030", "Kacang_Garing", "Snack", 85, 180, "E2", "16-06-2025"},
		{"P031", "Mangga_Manalagi", "Buah", 50, 9, "A1", "17-06-2025"},
		{"P032", "Tomat_Merah", "Sayur", 75, 6, "A2", "17-06-2025"},
		{"P033", "Daging_Bebek", "Daging", 18, 5, "B1", "18-06-2025"},
		{"P034", "Ikan_Gurame", "Ikan", 33, 5, "B2", "18-06-2025"},
		{"P035", "Susu_Cokelat", "Susu", 48, 15, "C1", "19-06-2025"},
		{"P036", "Roti_Keju", "Roti", 28, 3, "C2", "19-06-2025"},
		{"P037", "Minuman_Soda", "Minuman", 110, 180, "D1", "20-06-2025"},
		{"P038", "Merica_Bubuk", "Bumbu", 22, 365, "D2", "20-06-2025"},
		{"P039", "Bakso_Ikan", "Frozen", 36, 60, "E1", "21-06-2025"},
		{"P040", "Permen_Karet", "Snack", 200, 365, "E2", "21-06-2025"},
	}
	for _, p := range data {
		if jumlahProduk >= MAX {
			break
		}
		daftarProduk[jumlahProduk] = p
		tanggalMasukList[jumlahProduk] = p.TanggalMasuk
		namaProdukList[jumlahProduk] = p.Nama
		jumlahProduk++
		if jumlahLog < 100 {
			logMasuk[jumlahLog] = "Tambah produk " + p.Nama + " (" + p.ID + ")"
			jumlahLog++
		}
	}
	fmt.Println(">> Data dummy sebanyak", jumlahProduk, "produk berhasil dimuat.")
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
		tanggalMasukList[i] = tanggalMasukList[i+1]
		namaProdukList[i] = namaProdukList[i+1]
	}
	jumlahProduk--
	fmt.Println("Produk berhasil dihapus.")
}

func main() {
	var pilihan int
	for pilihan != 13 {
		fmt.Println("\n--- Menu Gudang FreshMart ---")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Tampilkan Semua Produk")
		fmt.Println("3. Cari Produk by ID")
		fmt.Println("4. Cari Produk by Nama (Binary Search)")
		fmt.Println("5. Cari Produk by Awalan Nama (Binary Prefix Search)")
		fmt.Println("6. Urutkan berdasarkan Umur Simpan")
		fmt.Println("7. Urutkan berdasarkan Kategori")
		fmt.Println("8. Lihat Log Transaksi")
		fmt.Println("9. Edit Stok Produk")
		fmt.Println("10. Hapus Produk")
		fmt.Println("11. Edit ID/Nama/Kategori Produk")
		fmt.Println("12. Muat Data Dummy (40 produk)")
		fmt.Println("13. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahProduk()
		case 2:
			TampilkanProduk()
		case 3:
			var id string
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			idx := SeqSearchCariProdukByID(id)
			if idx != -1 {
				fmt.Println("Ditemukan:", daftarProduk[idx])
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
		case 4:
			InsertionSortByNama(true)
			var nama string
			fmt.Print("Masukkan Nama Produk: ")
			fmt.Scan(&nama)
			idx := BinarySearchByNama(nama)
			if idx != -1 {
				fmt.Println("Ditemukan:", daftarProduk[idx])
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
		case 5:
			InsertionSortByNama(true)
			fmt.Print("Masukkan awalan nama produk: ")
			var prefix string
			fmt.Scan(&prefix)
			hasil := BinarySearchByPrefixNama(prefix)
			if len(hasil) == 0 {
				fmt.Println("Tidak ditemukan produk dengan prefix tersebut.")
			} else {
				fmt.Println("Ditemukan:")
				for i, p := range hasil {
					fmt.Printf("%d. [%s] %s - %s | Stok: %d | Umur Simpan: %d hari | Rak: %s | Masuk: %s\n",
						i+1, p.ID, p.Nama, p.Kategori, p.Stok, p.UmurSimpanHari, p.LokasiRak, p.TanggalMasuk)
				}
			}
		case 6:
			var asc int
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			SelectionSortUrutkanByUmurSimpan(asc == 1)
			fmt.Println("Berhasil diurutkan.")
		case 7:
			var asc int
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			InsertionSortUrutkanByKategori(asc == 1)
			fmt.Println("Berhasil diurutkan.")
		case 8:
			CetakLog()
		case 9:
			EditStokProduk()
		case 10:
			HapusProduk()
		case 11:
			EditDataProduk()
		case 12:
			IsiDataDummy()
		case 13:
			fmt.Println("Terima kasih! Program selesai.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
