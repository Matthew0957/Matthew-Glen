package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta untuk kapasitas array statis
const MAX_FILM int = 100
const MAX_JADWAL int = 200

// Tipe Bentukan (Struct) untuk data Film dan Jadwal
type Film struct {
	Judul          string
	Genre          string
	Durasi         int
	JumlahPenonton int
}

type Jadwal struct {
	IDJadwal  int
	JudulFilm string
	JamTayang string
	SisaKursi int
}

// Variabel Global untuk penyimpanan data utama
var daftarFilm [MAX_FILM]Film
var daftarJadwal [MAX_JADWAL]Jadwal
var nFilm, nJadwal int

// Scanner global untuk membaca input dengan spasi
var scanner = bufio.NewScanner(os.Stdin)

// Fungsi helper: membaca satu baris input sebagai string (mendukung spasi)
func bacaString(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Fungsi helper: membaca satu integer dari input
func bacaInt(prompt string) int {
	var n int
	fmt.Print(prompt)
	fmt.Scan(&n)
	scanner.Scan() // buang newline sisa di buffer
	return n
}

func main() {
	var pilihan int
	nFilm = 0
	nJadwal = 0

	pilihan = -1
	for pilihan != 0 {
		fmt.Println("\033[33m╔══════════════════════════════════════╗\033[0m")
		fmt.Println("\033[33m║ ✮ APLIKASI PEMESANAN TIKET BIOSKOP ✮ ║\033[0m")
		fmt.Println("\033[33m╚══════════════════════════════════════╝\033[0m")
		fmt.Println("1. Kelola Data Film (Tambah/Ubah/Hapus)")
		fmt.Println("2. Kelola Jadwal Tayang")
		fmt.Println("3. Pesan Tiket")
		fmt.Println("4. Cari Film (Judul/Genre)")
		fmt.Println("5. Laporan Film Terlaris (Sorting)")
		fmt.Println("6. Urutkan Film berdasarkan Judul (Insertion Sort)")
		fmt.Println("7. Cari Film berdasarkan Judul (Binary Search)")
		fmt.Println("0. Keluar")
		pilihan = bacaInt("Pilih menu: ")

		if pilihan == 1 {
			menuFilm()
		} else if pilihan == 2 {
			menuJadwal()
		} else if pilihan == 3 {
			pesanTiket()
		} else if pilihan == 4 {
			menuCari()
		} else if pilihan == 5 {
			laporanSorting()
		} else if pilihan == 6 {
			insertionSortJudul()
		} else if pilihan == 7 {
			cariBinarySearch()
		} else if pilihan != 0 {
			fmt.Println("\033[31mPilihan tidak valid.\033[0m")
		}
	}
	fmt.Println("\033[32mTerima kasih telah menggunakan aplikasi ini.\033[0m")
}

// --- FUNGSI & PROSEDUR MODULAR ---

func menuFilm() {
	fmt.Println("\n1. Tambah Film\n2. Ubah Film\n3. Hapus Film")
	aksi := bacaInt("Aksi: ")

	if aksi == 1 {
		// FIX: cek kapasitas sebelum tambah
		if nFilm >= MAX_FILM {
			fmt.Println("\033[31mData film sudah penuh.\033[0m")
			return
		}
		// FIX: pakai bacaString agar judul/genre bisa mengandung spasi
		daftarFilm[nFilm].Judul = bacaString("Judul: ")
		daftarFilm[nFilm].Genre = bacaString("Genre: ")
		daftarFilm[nFilm].Durasi = bacaInt("Durasi (menit): ")
		daftarFilm[nFilm].JumlahPenonton = 0
		nFilm++
		fmt.Println("\033[32mFilm berhasil ditambahkan.\033[0m")

	} else if aksi == 2 {
		target := bacaString("Judul film yang diubah: ")
		idx := cariFilmByJudul(target)
		if idx != -1 {
			daftarFilm[idx].Genre = bacaString("Genre Baru: ")
			daftarFilm[idx].Durasi = bacaInt("Durasi Baru (menit): ")
			fmt.Println("\033[32mFilm berhasil diubah.\033[0m")
		} else {
			fmt.Println("\033[31mFilm tidak ditemukan.\033[0m")
		}

	} else if aksi == 3 {
		target := bacaString("Judul film yang dihapus: ")
		idx := cariFilmByJudul(target)
		if idx != -1 {
			for i := idx; i < nFilm-1; i++ {
				daftarFilm[i] = daftarFilm[i+1]
			}
			nFilm--
			fmt.Println(" \033[32mFilm berhasil dihapus.\033[0m")
		} else {
			fmt.Println("\033[31mFilm tidak ditemukan.\033[0m")
		}
	} else {
		fmt.Println("\033[31mAksi tidak valid.\033[0m")
	}
}

func menuJadwal() {
	fmt.Println("\n1. Tambah Jadwal\n2. Ubah Jadwal\n3. Hapus Jadwal")
	aksi := bacaInt("Aksi: ")

	if aksi == 1 {
		if nJadwal >= MAX_JADWAL {
			fmt.Println("\033[31mData jadwal sudah penuh.\033[0m")
			return
		}

		// Ambil input ID Jadwal dulu
		idJadwal := bacaInt("ID Jadwal: ")

		// Ambil input Judul Film dan langsung divalidasi
		judulFilm := bacaString("Judul Film: ")

		// VALIDASI: Cek apakah film sudah terdaftar di data film
		idxF := cariFilmByJudul(judulFilm)
		if idxF == -1 {
			fmt.Println("\033[31mGagal: Film belum terdaftar! Silakan tambahkan film terlebih dahulu di menu 1.\033[0m")
			return // Keluar dari fungsi, jadwal tidak jadi ditambahkan
		}

		// Jika film ditemukan, lanjutkan pengisian data jadwal
		daftarJadwal[nJadwal].IDJadwal = idJadwal
		daftarJadwal[nJadwal].JudulFilm = daftarFilm[idxF].Judul // Menggunakan nama resmi dari daftarFilm agar case-nya konsisten
		daftarJadwal[nJadwal].JamTayang = bacaString("Jam Tayang (contoh: 14:00): ")
		daftarJadwal[nJadwal].SisaKursi = bacaInt("Kapasitas Kursi: ")
		nJadwal++
		fmt.Println("\033[32mJadwal berhasil ditambahkan.\033[0m")

	} else if aksi == 2 {
		id := bacaInt("ID Jadwal yang diubah: ")
		idxJ := cariJadwalByID(id)
		if idxJ != -1 {
			judulBaru := bacaString("Judul Film Baru: ")

			// VALIDASI: Cek juga saat mengubah jadwal film
			idxF := cariFilmByJudul(judulBaru)
			if idxF == -1 {
				fmt.Println("\033[31mGagal: Film baru belum terdaftar!\033[0m")
				return
			}

			daftarJadwal[idxJ].JudulFilm = daftarFilm[idxF].Judul
			daftarJadwal[idxJ].JamTayang = bacaString("Jam Tayang Baru: ")
			daftarJadwal[idxJ].SisaKursi = bacaInt("Sisa Kursi Baru: ")
			fmt.Println("\033[32mJadwal berhasil diubah.\033[0m")
		} else {
			fmt.Println("\033[31mJadwal tidak ditemukan.\033[0m")
		}

	} else if aksi == 3 {
		id := bacaInt("ID Jadwal yang dihapus: ")
		idxJ := cariJadwalByID(id)
		if idxJ != -1 {
			for i := idxJ; i < nJadwal-1; i++ {
				daftarJadwal[i] = daftarJadwal[i+1]
			}
			nJadwal--
			fmt.Println("\033[32mJadwal berhasil dihapus.\033[0m")
		} else {
			fmt.Println("\033[31mJadwal tidak ditemukan.\033[0m")
		}
	} else {
		fmt.Println("\033[31mAksi tidak valid.\033[0m")
	}
}

func pesanTiket() {
	id := bacaInt("Masukkan ID Jadwal: ")

	idxJ := cariJadwalByID(id)
	if idxJ == -1 {
		fmt.Println("\033[31mJadwal tidak ditemukan.\033[0m")
		return
	}

	fmt.Printf("Film: %s | Jam: %s | Sisa Kursi: %d\n",
		daftarJadwal[idxJ].JudulFilm,
		daftarJadwal[idxJ].JamTayang,
		daftarJadwal[idxJ].SisaKursi)

	jml := bacaInt("Jumlah Kursi dipesan: ")
	if jml <= 0 {
		fmt.Println("\033[31mJumlah kursi harus lebih dari 0.\033[0m")
		return
	}

	if daftarJadwal[idxJ].SisaKursi >= jml {
		daftarJadwal[idxJ].SisaKursi -= jml
		idxF := cariFilmByJudul(daftarJadwal[idxJ].JudulFilm)
		if idxF != -1 {
			daftarFilm[idxF].JumlahPenonton += jml
		}
		fmt.Println("\033[32mPemesanan Berhasil!\033[0m")
	} else {
		fmt.Printf("\033[31mKursi tidak mencukupi. Sisa kursi: %d\n\033[0m", daftarJadwal[idxJ].SisaKursi)
	}
}

func menuCari() {
	fmt.Println("\n1. Cari berdasarkan Judul\n2. Cari berdasarkan Genre")
	// FIX: bacaInt menangani newline sisa di buffer dengan benar
	kriteria := bacaInt("Pilih kriteria: ")
	kataKunci := bacaString("Kata kunci: ")

	if kriteria == 1 {
		idx := cariFilmByJudul(kataKunci)
		if idx != -1 {
			fmt.Printf("Ditemukan: %s (%s) - %d mnt | Penonton: %d\n",
				daftarFilm[idx].Judul,
				daftarFilm[idx].Genre,
				daftarFilm[idx].Durasi,
				daftarFilm[idx].JumlahPenonton)
		} else {
			fmt.Println("\033[31mFilm tidak ditemukan.\033[0m")
		}
	} else if kriteria == 2 {
		ditemukan := false
		for i := 0; i < nFilm; i++ {
			if strings.EqualFold(daftarFilm[i].Genre, kataKunci) {
				fmt.Printf("- %s (%d mnt)\n", daftarFilm[i].Judul, daftarFilm[i].Durasi)
				ditemukan = true
			}
		}
		if !ditemukan {
			fmt.Println("\033[31mTidak ada film dengan genre tersebut.\033[0m")
		}
	} else {
		fmt.Println("\033[31mKriteria tidak valid.\033[0m")
	}
}

func laporanSorting() {
	if nFilm == 0 {
		fmt.Println("\033[31mBelum ada data film.\033[0m")
		return
	}

	// FIX: sorting dilakukan pada salinan array, bukan data asli
	salinan := daftarFilm
	for i := 0; i < nFilm-1; i++ {
		maxIdx := i
		for j := i + 1; j < nFilm; j++ {
			if salinan[j].JumlahPenonton > salinan[maxIdx].JumlahPenonton {
				maxIdx = j
			}
		}
		temp := salinan[i]
		salinan[i] = salinan[maxIdx]
		salinan[maxIdx] = temp
	}

	fmt.Println("\n--- DAFTAR FILM TERLARIS ---")
	for i := 0; i < nFilm; i++ {
		fmt.Printf("%d. %-30s - %d Penonton\n", i+1, salinan[i].Judul, salinan[i].JumlahPenonton)
	}
}

// --- FUNGSI BARU: INSERTION SORT ---
// Mengurutkan daftarFilm berdasarkan Judul secara Ascending (A-Z)
// menggunakan algoritma Insertion Sort.
// FIX: pengurutan langsung mengubah daftarFilm asli, agar data konsisten
// untuk kebutuhan Binary Search selanjutnya.
func insertionSortJudul() {
	if nFilm == 0 {
		fmt.Println("\033[31mBelum ada data film.\033[0m")
		return
	}

	for i := 1; i < nFilm; i++ {
		current := daftarFilm[i]
		j := i - 1

		// FIX: bandingkan secara case-insensitive memakai ToLower
		for j >= 0 && strings.ToLower(daftarFilm[j].Judul) > strings.ToLower(current.Judul) {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = current
	}

	fmt.Println("\033[32mData film berhasil diurutkan berdasarkan Judul (A-Z).\033[0m")
	fmt.Println("\n--- DAFTAR FILM TERURUT ---")
	for i := 0; i < nFilm; i++ {
		fmt.Printf("%d. %-30s (%s) - %d mnt\n", i+1, daftarFilm[i].Judul, daftarFilm[i].Genre, daftarFilm[i].Durasi)
	}
}

// --- FUNGSI BARU: BINARY SEARCH ---
// Mencari film berdasarkan Judul menggunakan algoritma Binary Search.
// CATATAN: Binary Search membutuhkan data yang sudah terurut.
// Jika data belum diurutkan, fungsi ini akan mengurutkannya otomatis
// dengan memanggil insertionSortJudul() terlebih dahulu.
func cariBinarySearch() {
	if nFilm == 0 {
		fmt.Println("\033[31mBelum ada data film.\033[0m")
		return
	}

	if !isSortedByJudul() {
		fmt.Println("\033[33mData belum terurut, mengurutkan data terlebih dahulu...\033[0m")
		insertionSortJudul()
	}

	target := bacaString("Masukkan Judul Film yang dicari: ")

	low := 0
	high := nFilm - 1
	ditemukan := -1
	jumlahLangkah := 0

	for low <= high {
		jumlahLangkah++
		mid := (low + high) / 2
		judulMid := strings.ToLower(daftarFilm[mid].Judul)
		judulTarget := strings.ToLower(target)

		if judulMid == judulTarget {
			ditemukan = mid
			break
		} else if judulMid < judulTarget {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if ditemukan != -1 {
		fmt.Printf("\033[32mFilm ditemukan dalam %d langkah pencarian.\033[0m\n", jumlahLangkah)
		fmt.Printf("Judul : %s\nGenre : %s\nDurasi: %d menit\nPenonton: %d\n",
			daftarFilm[ditemukan].Judul,
			daftarFilm[ditemukan].Genre,
			daftarFilm[ditemukan].Durasi,
			daftarFilm[ditemukan].JumlahPenonton)
	} else {
		fmt.Printf("\033[31mFilm tidak ditemukan setelah %d langkah pencarian.\033[0m\n", jumlahLangkah)
	}
}

// Fungsi pembantu: memeriksa apakah daftarFilm sudah terurut berdasarkan Judul
func isSortedByJudul() bool {
	for i := 0; i < nFilm-1; i++ {
		if strings.ToLower(daftarFilm[i].Judul) > strings.ToLower(daftarFilm[i+1].Judul) {
			return false
		}
	}
	return true
}

// Fungsi pembantu: Sequential Search berdasarkan judul
func cariFilmByJudul(judul string) int {
	res := -1
	i := 0
	for i < nFilm && res == -1 {
		// FIX: pakai EqualFold agar pencarian tidak case-sensitive
		if strings.EqualFold(daftarFilm[i].Judul, judul) {
			res = i
		}
		i++
	}
	return res
}

// Fungsi pembantu: Sequential Search berdasarkan ID Jadwal
func cariJadwalByID(id int) int {
	for i := 0; i < nJadwal; i++ {
		if daftarJadwal[i].IDJadwal == id {
			return i
		}
	}
	return -1
}
