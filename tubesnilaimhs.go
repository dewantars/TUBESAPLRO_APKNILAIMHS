package main

import "fmt"

const NMAX int = 1000

type tabNilai struct{
	matakuliah string
	quiz int
	total int
	UTS int
	UAS int
	grade string
}


type mahasiswa struct{
	nama string
	nim string
	kelas string
	j int
	prodi string
	ip float64
	matkul [NMAX]tabNilai
}

type tabMhs [NMAX]mahasiswa

func main() {
	var A tabMhs
	var n int

	inputMenu(&A, &n)
}

func mainMenu() {
	fmt.Println("============ Main Menu ============")
	fmt.Println("1. Mengisi data mahasiswa          ")
	fmt.Println("2. Mengubah data mahasiswa         ")
	fmt.Println("3. Menambah data mahasiswa         ")
	fmt.Println("4. Menghapus data mahasiswa        ")
	fmt.Println("5. Menampilkan data mahasiswa      ")
	fmt.Println("6. Exit                            ")
	fmt.Println("===================================")
}

func inputMenu(A *tabMhs, n *int) {
	// IS. -
	// FS. memanggil procedure berdasarkan inputan user
	mainMenu()
	var menu int
	fmt.Scan(&menu)
	switch menu {
	case 1:
		mengisiData(*&A, *&n)
	case 2:
		mengubahData(*&A, *&n)
	case 3:
		menambahData(*&A, *&n)
	case 4:
		menghapusData(*&A, *&n)
	case 5:
		menampilkanData(*&A, *&n)
	case 6:
		fmt.Println("Terima kasih. Program berhenti.")
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		inputMenu(*&A, *&n)
	}
}

func mengisiData(A *tabMhs, n *int) {
	// IS. -
	// FS. memanggil procedure berdasarkan inputan user
	var pilihan int
	menuIsi()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		dataMhs(*&A, *&n)
	case 2:
		nilaiMhs(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		mengisiData(*&A, *&n)
	}
}

func menuIsi() {
	fmt.Println("======== Isi Data Mahasiswa ========")
	fmt.Println("1. Nama, NIM, Kelas, Prodi, Semester")
	fmt.Println("2. Nilai Mahasiswa                  ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func dataMhs(A *tabMhs, n *int) {
	// IS. -
	// FS. A terisi data mahasiswa sebanyak n
	fmt.Println("Masukkan jumlah data mahasiswa")
	fmt.Scan(n)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	for i := 0; i < *n; i++{
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].kelas, &A[i].prodi)
	}
	fmt.Println("Data berhasil ditambahkan")
	mengisiData(*&A, *&n)
}



func nilaiMhs(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. A terisi nilai mahasiswa
	var index, i int
	var idmhs string

	fmt.Println("Masukkan ID mahasiswa")
	fmt.Scan(&idmhs)	
	index = SearchId(*A, *n, idmhs)
	for index == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&idmhs)
		index = SearchId(*A, *n, idmhs)
	}
	fmt.Println("Masukkan jumlah mata kuliah")
	fmt.Scan(&A[index].j)
	for i = 0; i<A[index].j; i++{
		fmt.Println("Masukkan nama matakuliah, nilai quiz, nilai UTS dan nilai UAS")
		fmt.Scan(&A[index].matkul[i].matakuliah, &A[index].matkul[i].quiz, &A[index].matkul[i].UTS, &A[index].matkul[i].UAS)
		A[index].matkul[i].total = A[index].matkul[i].quiz+A[index].matkul[i].UTS+A[index].matkul[i].UAS
		var rata int
		rata = A[index].matkul[i].total / 3
		if rata > 80{
			A[index].matkul[i].grade = "A"
		}else if rata <= 80 && rata > 70{
			A[index].matkul[i].grade = "AB"
		}else if rata <= 70 && rata > 65{
			A[index].matkul[i].grade = "B"
		}else if rata <= 65 && rata > 60{
			A[index].matkul[i].grade = "BC"
		}else if rata <= 60 && rata > 50{
			A[index].matkul[i].grade = "C"
		}else if rata <= 50 && rata > 40{
			A[index].matkul[i].grade = "D"
		}else{
			A[index].matkul[i].grade = "E"
		}
	}
	fmt.Println("Data berhasil ditambahkan")
	mengisiData(*&A, *&n)
}

func SearchId(A tabMhs, n int, x string) int {
	// mengembalikkan index array data mahasiswa(A)
	var i int
	var ketemu int

	ketemu = -1
	for i = 0; i < n; i++{
		if x == A[i].nim {
			ketemu = i
		}
	}
	return ketemu
}

func mengubahData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa
	// FS. Memanggil Procedure sesuai inputan user
	var pilihan int
	menuUbah()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		ubahData(*&A, *&n)
	case 2:
		ubahNilai(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		mengubahData(*&A, *&n)
	}
}

func menuUbah() {
	fmt.Println("======= Ubah Data Mahasiswa ========")
	fmt.Println("1. Biodata mahasiswa                ")
	fmt.Println("2. Nilai Mahasiswa                  ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func ubahData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. A telah berhasil diedit
	var idx int
	var x string
	fmt.Println("Masukkan id mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	fmt.Println(A[idx].nama, A[idx].nim, A[idx].kelas, A[idx].prodi)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	fmt.Scan(&A[idx].nama, &A[idx].nim, &A[idx].kelas, &A[idx].prodi)
	fmt.Println("Data berhasil diedit")
	mengubahData(*&A, *&n)
}

func ubahNilai(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. A telah berhasil diedit
	var x string
	var idx, i int
	var matkul string
	fmt.Println("Masukkan id mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak ditemukan. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	fmt.Println("Masukkan nama mata kuliah yang nilainya ingin diubah")
	fmt.Scan(&matkul)
	i = searchMatkul(*A, idx, A[idx].j, matkul)
	for i == -1{
		fmt.Println("Masukkan tidak ditemukan. Silahkan masukkan kembali.")
		fmt.Scan(&matkul)
		i = searchMatkul(*A, idx, A[idx].j, matkul)
	}
	fmt.Println(A[idx].matkul[i].matakuliah, A[idx].matkul[i].quiz, A[idx].matkul[i].UTS, A[idx].matkul[i].UAS, A[idx].matkul[i].total, A[idx].matkul[i].grade)
	fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
	fmt.Scan(&A[idx].matkul[i].quiz, &A[idx].matkul[i].UTS, &A[idx].matkul[i].UAS)
	A[idx].matkul[i].total = A[idx].matkul[i].quiz+A[idx].matkul[i].UTS+A[idx].matkul[i].UAS
	var rata int
	rata = A[idx].matkul[i].total / 3
	if rata > 80{
		A[idx].matkul[i].grade = "A"
	}else if rata <= 80 && rata > 70{
		A[idx].matkul[i].grade = "AB"
	}else if rata <= 70 && rata > 65{
		A[idx].matkul[i].grade = "B"
	}else if rata <= 65 && rata > 60{
		A[idx].matkul[i].grade = "BC"
	}else if rata <= 60 && rata > 50{
		A[idx].matkul[i].grade = "C"
	}else if rata <= 50 && rata > 40{
		A[idx].matkul[i].grade = "D"
	}else{
		A[idx].matkul[i].grade = "E"
	}
	fmt.Println(A[idx].matkul[i].matakuliah, A[idx].matkul[i].quiz, A[idx].matkul[i].UTS, A[idx].matkul[i].UAS, A[idx].matkul[i].total, A[idx].matkul[i].grade)
	fmt.Println("Data berhasil diedit")
	mengubahData(*&A, *&n)
}

func searchMatkul(A tabMhs, n int, j int, matkul string) int {
	// mengembalikan index dari array matkul
	var ketemu, i int
	ketemu = -1
	for i = 0; i < j; i++{
		if matkul == A[n].matkul[i].matakuliah{
			ketemu = i
		}
	}
	return ketemu
}

func menghapusData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. memanggil procedure sesuai inputan user
	var pilihan int
	menuHapus()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		hapusData(*&A, *&n)
	case 2:
		hapusNilai(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		menghapusData(*&A, *&n)
	}
}

func menuHapus() {
	fmt.Println("======= Hapus Data Mahasiswa =======")
	fmt.Println("1. Biodata mahasiswa dan nilai      ")
	fmt.Println("2. nilai mata kuliah                ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func hapusData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. salah satu data pada array A berhasil dihapus
	var x string
	var idx int
	fmt.Println("Masukkan ID Mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	for i:= idx; i<*n-1; i++{
		A[i] = A[i+1]
	}
	*n = *n-1
	fmt.Println("Data berhasil dihapus")
	menghapusData(*&A, *&n)
}

func hapusNilai(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. salah satu nilai satu mata kuliah berhasil dihapus
	var x, matkul string
	var idx, i int
	fmt.Println("Masukkan ID Mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak ditemukan. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	fmt.Println("Masukkan Nama Matakuliah")
	fmt.Scan(&matkul)
	i = searchMatkul(*A, idx, A[idx].j, matkul)
	for i == -1{
		fmt.Println("Masukkan tidak ditenukan. Silahkan masukkan kembali.")
		fmt.Scan(&matkul)
		i = searchMatkul(*A, idx, A[idx].j, matkul)
	}
	var j int
	j = A[idx].j
	for a := i; a<j-1; a++{
		A[a].matkul = A[a+1].matkul
	}
	A[idx].j = A[idx].j - 1
	fmt.Println("Data berhasil dihapus")
	menghapusData(*&A, *&n)
}

func menambahData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. memanggil procedure sesuai inputan user
	var pilihan int
	menuTambah()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahData(*&A, *&n)
	case 2:
		tambahMatkul(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		menambahData(*&A, *&n)
	}
}

func menuTambah() {
	fmt.Println("====== Tambah Data Mahasiswa =======")
	fmt.Println("1. Mahasiswa                        ")
	fmt.Println("2. Mata kuliah dan nilai            ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func tambahData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. A dan n bertambah 
	var x int
	fmt.Println("Masukkan jumlah mahasiswa yang ingin ditambahkan")
	fmt.Scan(&x)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	for i := *n; i < *n+x; i++{
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].kelas, &A[i].prodi)
	}
	*n = *n+x
	fmt.Println("Data berhasil ditambahkan")
	menambahData(*&A, *&n)
}

func tambahMatkul(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. A.matkul bertambah
	var index int
	var idmhs string
	var matkul string

	fmt.Println("Masukkan nama mahasiswa")
	fmt.Scan(&idmhs)	
	index = SearchId(*A, *n, idmhs)
	for index == -1{
		fmt.Println("Masukkan tidak ditemukan. Silahkan masukkan kembali.")
		fmt.Scan(&idmhs)
		index = SearchId(*A, *n, idmhs)
	}
	fmt.Println("Masukkan nama mata kuliah")
	fmt.Scan(&matkul)
	var j int
	j = A[index].j
	for matkul != "0"{
		A[index].matkul[j].matakuliah = matkul
		fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
		fmt.Scan(&A[index].matkul[j].quiz, &A[index].matkul[j].UTS, &A[index].matkul[j].UAS)
		A[index].matkul[j].total = A[index].matkul[j].quiz+A[index].matkul[j].UTS+A[index].matkul[j].UAS
		var rata int
		rata = A[index].matkul[j].total / 3
		if rata > 80{
			A[index].matkul[j].grade = "A"
		}else if rata <= 80 && rata > 70{
			A[index].matkul[j].grade = "AB"
		}else if rata <= 70 && rata > 65{
			A[index].matkul[j].grade = "B"
		}else if rata <= 65 && rata > 60{
			A[index].matkul[j].grade = "BC"
		}else if rata <= 60 && rata > 50{
			A[index].matkul[j].grade = "C"
		}else if rata <= 50 && rata > 40{
			A[index].matkul[j].grade = "D"
		}else{
			A[index].matkul[j].grade = "E"
		}
		j++
		A[index].j = A[index].j + 1
		fmt.Println("Masukkan nama mata kuliah atau ketik 0 jika sudah selesai")
		fmt.Scan(&matkul)
	}
	fmt.Println("Data berhasil ditambahkan")
	menambahData(*&A, *&n)
}


func menampilkanData(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. memanggil procedure sesuai inputan user
	var pilihan int
	var x int
	jumlahIp(*&A, *n)
	menuCetak()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		cetakDataMhs(*&A, *&n)
	case 2:
		cetakNilaiMhs(*&A, *&n)
	case 3:
		searchMhs(*&A, *&n)
	case 4:
		searchNamaMatkul(*&A, *&n)
	case 5:
		sortIp(*&A, *&n)
	case 6:
		sortSks(*&A, *&n)
	case 7:
		sortSks(*&A, *&n)
		fmt.Println("Masukkan jumlah sks yang ingin dicari")
		fmt.Scan(&x)
		fmt.Println(binarySearch(*A, *n, x))
		menampilkanData(*&A, *&n)
	case 8:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		menampilkanData(*&A, *&n)
	}
}

func menuCetak() {
	fmt.Println("========== Cetak Data Mahasiswa ===========")
	fmt.Println("1. Daftar mahasiswa keseluruhan            ")
	fmt.Println("2. Transkrip Nilai Mahasiswa               ")
	fmt.Println("3. Search mahasiswa berdasarkan matakuliah ")
	fmt.Println("4. Search matakuliah mahasiswa             ")
	fmt.Println("5. Sort nilai mahasiswa                    ")
	fmt.Println("6. Sort jumlah sks mahasiswa               ")
	fmt.Println("7. Searching SKS                           ")
	fmt.Println("8. Kembali                                 ")
	fmt.Println("===========================================")
}

func cetakDataMhs(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Menampilkan seluruh data mahasiswa kecuali nilai
	fmt.Println("Data mahasiswa: ")
	fmt.Println("NAMA  NIM  KELAS  PRODI")
	for i := 0; i < *n; i++{
		fmt.Println(A[i].nama, A[i].nim, A[i].kelas, A[i].prodi)
	}
	menampilkanData(*&A, *&n)
}

func cetakNilaiMhs(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Menampilkan nilai mahasiswa dari satu orang mahasiswa
	var x string
	var idx int
	var j int
	fmt.Println("Masukkan id mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	j = A[idx].j
	fmt.Println("Transkrip nilai mahasiswa", A[idx].nama)
	for i:= 0; i<j; i++{
		fmt.Println(A[idx].matkul[i].matakuliah, A[idx].matkul[i].quiz, A[idx].matkul[i].UTS, A[idx].matkul[i].UAS, A[idx].matkul[i].total, A[idx].matkul[i].grade)
	}
	fmt.Println("IP", A[idx].ip)
	menampilkanData(*&A, *&n)
}

func searchNamaMatkul(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Menampilkan nama matakuliah dari seorang mahasiswa yang dicari
	var idx int
	var x string

	fmt.Println("Masukkan ID Mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	fmt.Println("Mata kuliah yang diambil oleh", A[idx].nama, "adalah")
	for i:= 0; i<*n; i++{
		fmt.Println(A[idx].matkul[i].matakuliah)
	}
	menampilkanData(*&A, *&n)
}


func searchMhs(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Menampilkan nama mahasiswa dari nama matakuliah yang dicari
	var matkul string
	var i, b int
	
	fmt.Println("Masukkan nama mata kuliahnya")
	fmt.Scan(&matkul)
	fmt.Println("Nama mahasiswa yang mengambil mata kuliah", matkul, "adalah")
	for i=0; i<*n; i++{
		for b = 0; b<A[i].j; b++{
			if matkul == A[i].matkul[b].matakuliah{
				fmt.Println(A[i].nama)
			}
		}
	}
	menampilkanData(*&A, *&n)
}

func jumlahIp(A *tabMhs, n int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Menghitung ip mahasiswa
	var i, j, p int
	var jumlah float64
	
	jumlah = 0
	for i = 0; i<n; i++{
		p = A[i].j
		for j = 0; j<p; j++{
			if A[i].matkul[j].grade == "A"{
				jumlah = jumlah + 4.0
			}else if A[i].matkul[j].grade == "AB" {
				jumlah += 3.5
			}else if A[i].matkul[j].grade == "B"{
				jumlah += 3.0
			}else if A[i].matkul[j].grade == "BC"{
				jumlah += 2.5
			}else if A[i].matkul[j].grade == "C"{
				jumlah += 2.0
			}else if A[i].matkul[j].grade == "D"{
				jumlah += 1.0
			}else{
				jumlah += 0.0
			}
		}
		A[i].ip = jumlah / float64(A[i].j)
	}
	
}

func sortIp(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Mengurutkan ip pada array A menggunakan selection sort (descending) dan menampilkannya ke layar
	var pass, max, i, a int
	var temp mahasiswa

	pass = *n-1
	for i = 0; i<pass; i++{
		max = i
		for a = i + 1; a<*n; a++{
			if A[max].ip < A[a].ip{
				max = a
			}
		}

		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}
	fmt.Println("Sort Berhasil")
	cetakIp(*A, *n)
	menampilkanData(*&A, *&n)
}

func cetakIp(A tabMhs, n int) {
	var i int
	fmt.Println("ID Mahasiswa, Nama Mahasiswa dan Ip Mahasiswa ")
	for i = 0; i<n; i++{
		fmt.Println(A[i].nim, A[i].nama, A[i].ip)
	}
}


func sortSks(A *tabMhs, n *int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Mengurutkan data mahasiswa berdasarkan sks dari yang terkecil (ascending) menggunakan insertion sort dan menampilkannya ke layar
	var i, j int
	var temp mahasiswa
	
	for i = 1; i<*n; i++{
		temp = A[i]
		j = i - 1
		for temp.j < A[j].j && j >= 0{
			A[j+1] = A[j]
			j=j-1
		}
		A[j+1] = temp
	}

	fmt.Println("Sort Berhasil")
	cetakSks(*A, *n)
	menampilkanData(*&A, *&n)
}

func cetakSks(A tabMhs, n int) {
	// IS. A terdefinisi sebagai data mahasiswa dan n terdefinisi sebagai banyaknya mahasiswa
	// FS. Menampilkan sks dari array A yang sudah diurutkan
	var i int
	fmt.Println("ID Mahasiswa, Nama Mahasiswa dan Jumlah SKS")
	for i = 0; i<n; i++{
		fmt.Println(A[i].nim, A[i].nama, A[i].j*3)
	}
}

func binarySearch(A tabMhs, n int, x int) bool {
	// mengembalikan nilai true apabila x terdapat pada A[].j, dan mengembalikan false apabila sebaliknya
	var ketemu bool
	var mid, right, left int

	x = x/3
	ketemu = false
	right = n-1
	left = 0
	for ketemu == false && left <= right{
		mid = (right + left) / 2
		if x == A[mid].j {
			ketemu = true
		} else if x > A[mid].j {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	return ketemu
}
