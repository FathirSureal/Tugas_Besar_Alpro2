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
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&p.Kategori)
	fmt.Print("Masukkan stok: ")
	fmt.Scan(&p.Stok)
	fmt.Print("Masukkan umur simpan (hari): ")
	fmt.Scan(&p.UmurSimpanHari)
	fmt.Print("Masukkan lokasi rak: ")
	fmt.Scan(&p.LokasiRak)
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

func CariProdukByID(id string) int {
	for i := 0; i < jumlahProduk; i++ {
		if daftarProduk[i].ID == id {
			return i
		}
	}
	return -1
}

func BinarySearchByNama(nama string) int {
	kiri := 0
	k := jumlahProduk - 1
	for kiri <= k {
		tengah := (kiri + k) / 2
		if daftarProduk[tengah].Nama == nama {
			return tengah
		} else if daftarProduk[tengah].Nama < nama {
			kiri = tengah + 1
		} else {
			k = tengah - 1
		}
	}
	return -1
}

func UrutkanByUmurSimpan(ascending bool) {
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

func UrutkanByKategori(ascending bool) {
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
	idx := CariProdukByID(id)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Printf("Stok lama: %d\n", daftarProduk[idx].Stok)
	fmt.Print("Masukkan stok baru: ")
	fmt.Scan(&daftarProduk[idx].Stok)
	fmt.Println("Stok berhasil diubah.")
}

func HapusProduk() {
	var id string
	fmt.Print("Masukkan ID produk yang ingin dihapus: ")
	fmt.Scan(&id)
	idx := CariProdukByID(id)
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
	for pilihan != 10 {
		fmt.Println("\n--- Menu Gudang FreshMart ---")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Tampilkan Semua Produk")
		fmt.Println("3. Cari Produk by ID")
		fmt.Println("4. Cari Produk by Nama (Binary Search)")
		fmt.Println("5. Urutkan berdasarkan Umur Simpan")
		fmt.Println("6. Urutkan berdasarkan Kategori")
		fmt.Println("7. Lihat Log Transaksi")
		fmt.Println("8. Edit Stok Produk")
		fmt.Println("9. Hapus Produk")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			TambahProduk()
		} else if pilihan == 2 {
			TampilkanProduk()
		} else if pilihan == 3 {
			var id string
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			idx := CariProdukByID(id)
			if idx != -1 {
				fmt.Println("Ditemukan:", daftarProduk[idx])
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
		} else if pilihan == 4 {
			UrutkanByKategori(true)
			var nama string
			fmt.Print("Masukkan Nama Produk: ")
			fmt.Scan(&nama)
			idx := BinarySearchByNama(nama)
			if idx != -1 {
				fmt.Println("Ditemukan:", daftarProduk[idx])
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
		} else if pilihan == 5 {
			var asc int
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			UrutkanByUmurSimpan(asc == 1)
			fmt.Println("Berhasil diurutkan.")
		} else if pilihan == 6 {
			var asc int
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			UrutkanByKategori(asc == 1)
			fmt.Println("Berhasil diurutkan.")
		} else if pilihan == 7 {
			CetakLog()
		} else if pilihan == 8 {
			EditStokProduk()
		} else if pilihan == 9 {
			HapusProduk()
		} else if pilihan == 10 {
			fmt.Println("Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
