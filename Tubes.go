/*
Kelompok 2
M NAUFAL RAMADHAN	103012300266
Farah Fadhilah		103012300451

Pengerjaan :
Deklarasi Variable Global : Naufal dan Farah
addProduct,updateProduct,deleteProduct,printProduct,menuProduct: Naufal
addBrand,updateBrand,deleteBrand,printBrand,menuBrand, isBrandExist, updateJumlahProduct : Naufal
searchProduct,searchBrand,menuSearch: Naufal
SortMenu,sortProductMenu,sortBrandMenu,insertionSortProductByYear,insertionSortProductByName
insertionSortProductByBrand,insertionSortProductBySales,insertionSortBrandByProductCount : Farah
top3BySales : Farah
fillDummyData : Naufal dan Farah
*/


package main

import (
	"fmt"
)

const NMAX int = 100

type Brand struct {
	brandId, brandName string
	jumlahProduct      int
}

type Product struct {
	productId, productName, jenisElektronik, productYear, brandId string
	jumlahPenjualan                                               int
}

type tabBrand [NMAX]Brand
type tabProduct [NMAX]Product

func main() {
	ClearScreen()
	header()
}

func header() {
	var dataProduct tabProduct
	var dataBrand tabBrand
	var nProduct, nBrand int

	
	fillDummyData(&dataProduct, &nProduct, &dataBrand, &nBrand)

	var pilihan int
	
	for  {
		ClearScreen()
		fmt.Println(" === SELAMAT DATANG === ")
		fmt.Println("  APLIKASI ELEKTRONIK   ")
		fmt.Println(" ALGORITMA PEMROGRAMAN  ")
		fmt.Println("       - 2024 -         ")
		fmt.Println(" ====================== ")
		fmt.Println("       - MENU -         ")
		fmt.Println("1. Product              ")
		fmt.Println("2. Brand                ")
		fmt.Println("3. Search               ")
		fmt.Println("4. Sort                 ")
		fmt.Println("5. Top 3 Sales          ")
		fmt.Println("6. Exit                 ")
		fmt.Println(" ====================== ")

		fmt.Print("Masukkan pilihan 1/2/3/4/5?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			ClearScreen()
			menuProduct(&dataProduct, &nProduct, &dataBrand, nBrand)
		case 2:
			ClearScreen()
			menuBrand(&dataBrand, &nBrand)
		case 3:
			ClearScreen()
			menuSearch(dataProduct, nProduct, dataBrand, nBrand)
		case 4:
			ClearScreen()
			sortMenu(&dataProduct, nProduct, &dataBrand, nBrand)
		case 5:
			ClearScreen()
			top3BySales(dataProduct, nProduct, dataBrand, nBrand)
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuProduct(A *tabProduct, n *int, B *tabBrand, m int) {
	var pilihan int
	for  {
		ClearScreen()
		fmt.Println("=== Menu Product ===")
		fmt.Println("1. Add Product")
		fmt.Println("2. Edit Product")
		fmt.Println("3. Delete Product")
		fmt.Println("4. Print Product")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Masukkan pilihan 1/2/3/4/5?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			ClearScreen()
			addProduct(A, n, B, m)
		case 2:
			ClearScreen()
			updateProduct(A, *n)
		case 3:
			ClearScreen()
			deleteProduct(A, n, B, m)
		case 4:
			ClearScreen()
			printProduct(*A, *n)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuBrand(A *tabBrand, n *int) {
	var pilihan int
	for {
		ClearScreen()
		fmt.Println("=== Menu Brand ===")
		fmt.Println("1. Add Brand")
		fmt.Println("2. Edit Brand")
		fmt.Println("3. Delete Brand")
		fmt.Println("4. Print Brand")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Masukkan pilihan 1/2/3/4/5?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			ClearScreen()
			addBrand(A, n)
		case 2:
			ClearScreen()
			updateBrand(A, *n)
		case 3:
			ClearScreen()
			deleteBrand(A, n)
		case 4:
			ClearScreen()
			printBrand(*A, *n)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuSearch(products tabProduct, nProduct int, brands tabBrand, nBrand int) {
	var pilihan int
	for  {
		ClearScreen()
		fmt.Println("=== Menu Search ===")
		fmt.Println("1. Search Product")
		fmt.Println("2. Search Brand")
		fmt.Println("3. Back to Main Menu")
		fmt.Print("Masukkan pilihan 1/2/3?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			ClearScreen()
			searchProduct(products, nProduct)
		case 2:
			ClearScreen()
			searchBrand(brands, nBrand, products, nProduct)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sortMenu(dataProduct *tabProduct, nProduct int, dataBrand *tabBrand, nBrand int) {
	var pilihan int
	for pilihan != 3 {
		ClearScreen()
		fmt.Println("=== Menu Sort ===")
		fmt.Println("1. Sort Products")
		fmt.Println("2. Sort Brands")
		fmt.Println("3. Back to Main Menu")
		fmt.Print("Masukkan pilihan 1/2/3?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			ClearScreen()
			sortProductMenu(dataProduct, nProduct)
		case 2:
			ClearScreen()
			sortBrandMenu(dataBrand, nBrand)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sortProductMenu(dataProduct *tabProduct, n int) {
	var pilihan int
	for {
		ClearScreen()
		fmt.Println(" === MENU SORT PRODUK === ")
		fmt.Println("1. Sort by Year           ")
		fmt.Println("2. Sort by Name           ")
		fmt.Println("3. Sort by Brand          ")
		fmt.Println("4. Sort by Sales          ")
		fmt.Println("5. Kembali                ")
		fmt.Println("==========================")
		fmt.Print("Masukkan pilihan 1/2/3/4/5?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			insertionSortProductByYear(dataProduct, n)
		case 2:
			insertionSortProductByName(dataProduct, n)
		case 3:
			insertionSortProductByBrand(dataProduct, n)
		case 4:
			insertionSortProductBySales(dataProduct, n)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
		ClearScreen()
		printProduct(*dataProduct, n)
	}
}

func sortBrandMenu(dataBrand *tabBrand, n int) {
	var pilihan int
	for {
		ClearScreen()
		fmt.Println(" === MENU SORT MEREK === ")
		fmt.Println("1. Sort by Product Count ")
		fmt.Println("2. Kembali               ")
		fmt.Println("=========================")
		fmt.Print("Masukkan pilihan 1/2?: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			insertionSortBrandByProductCount(dataBrand, n)
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
		ClearScreen()
		printBrand(*dataBrand, n)
	}
}

func insertionSortProductByYear(A *tabProduct, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].productYear > key.productYear {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func insertionSortProductByName(A *tabProduct, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].productName > key.productName {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func insertionSortProductByBrand(A *tabProduct, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].brandId > key.brandId {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func insertionSortProductBySales(A *tabProduct, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].jumlahPenjualan < key.jumlahPenjualan {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func insertionSortBrandByProductCount(A *tabBrand, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].jumlahProduct < key.jumlahProduct {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func addProduct(A *tabProduct, n *int, B *tabBrand, m int) {
	fmt.Print("Apakah anda ingin menambahkan Data Produk? (Y/N): ")
	var pilihan string
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		fmt.Println("Masukkan data produk:")
		for i := *n; i < *n+1; i++ {
			fmt.Print("Masukkan ProductId: ")
			fmt.Scan(&A[i].productId)
			fmt.Print("Masukkan nama produk: ")
			fmt.Scan(&A[i].productName)
			fmt.Print("Masukkan jenis elektronik: ")
			fmt.Scan(&A[i].jenisElektronik)
			fmt.Print("Masukkan tahun Keluar: ")
			fmt.Scan(&A[i].productYear)
			fmt.Print("Masukkan BrandId: ")
			fmt.Scan(&A[i].brandId)
			if !isBrandExist(*B, m, A[i].brandId) {
				fmt.Println("BrandId tidak valid")
				i--
				continue
			}
			fmt.Print("Masukkan jumlah penjualan: ")
			fmt.Scan(&A[i].jumlahPenjualan)
			
		}
		*n++
		updateJumlahProduct(*A, *n, B, m)
	} else if pilihan == "N" || pilihan == "n" {
		pauseAndClear()
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func printProduct(A tabProduct, n int) {
	fmt.Println("Data Produk:")
	fmt.Println("======================================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("%2s %25s %10s %5s %2s %5d\n",A[i].productId, A[i].productName, A[i].jenisElektronik, A[i].productYear, A[i].brandId, A[i].jumlahPenjualan)
	}
	fmt.Println("======================================================================")
	pauseAndClear()
}


func updateProduct(A *tabProduct, n int) {
	var index int
	var pilihan string
	fmt.Print("Apakah anda ingin mengupdate data produk? (Y/N): ")
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		fmt.Print("Masukkan Indeks yang akan diupdate: ")
		fmt.Scan(&index)
		if index < 0 || index >= n {
			fmt.Println("Index tidak valid")
			return
		}
		fmt.Print("Masukkan ProductId: ")
		fmt.Scan(&A[index].productId)
		fmt.Print("Masukkan nama produk: ")
		fmt.Scan(&A[index].productName)
		fmt.Print("Masukkan jenis elektronik: ")
		fmt.Scan(&A[index].jenisElektronik)
		fmt.Print("Masukkan tahun Keluar: ")
		fmt.Scan(&A[index].productYear)
		fmt.Print("Masukkan jumlah penjualan: ")
		fmt.Scan(&A[index].jumlahPenjualan)
		fmt.Println("Data telah berhasil diupdate")
	} else if pilihan == "N" || pilihan == "n" {
		pauseAndClear()
	}
}

func deleteProduct(A *tabProduct, n *int, B *tabBrand, m int) {
	var index int
	var pilihan string
	fmt.Print("Apakah anda ingin menghapus data produk? (Y/N): ")
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		fmt.Print("Masukkan Indeks yang akan dihapus: ")
		fmt.Scan(&index)
		if index < 0 || index >= *n {
			fmt.Println("Index tidak valid")
			return
		}
		for i := index; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		updateJumlahProduct(*A, *n, B, m)
		fmt.Println("Data berhasil dihapus")
	} else if pilihan == "N" || pilihan == "n" {
		pauseAndClear()
	}
}

func addBrand(A *tabBrand, n *int) {
	fmt.Print("Apakah anda ingin menambahkan Data Merek? (Y/N): ")
	var pilihan string
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		fmt.Println("Masukkan data merek:")
		for i := *n; i < *n+1; i++ {
			fmt.Print("Masukkan BrandId: ")
			fmt.Scan(&A[i].brandId)
			fmt.Print("Masukkan nama merek: ")
			fmt.Scan(&A[i].brandName)
			A[i].jumlahProduct = 0
		}
		*n++
	} else if pilihan == "N" || pilihan == "n" {
		pauseAndClear()
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func printBrand(A tabBrand, n int) {
	fmt.Println("Data Merek:")
	fmt.Println("==============================")
	for i := 0; i < n; i++ {
		fmt.Printf("%2s %15s %5d \n",A[i].brandId, A[i].brandName, A[i].jumlahProduct)
	}
	fmt.Println("==============================")
	pauseAndClear()
}

func updateBrand(A *tabBrand, n int) {
	var index int
	var pilihan string
	fmt.Print("Apakah anda ingin mengupdate data merek? (Y/N): ")
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		fmt.Print("Masukkan Indeks yang akan diupdate: ")
		fmt.Scan(&index)
		if index < 0 || index >= n {
			fmt.Println("Index tidak valid")
			return
		}
		fmt.Print("Masukkan BrandId: ")
		fmt.Scan(&A[index].brandId)
		fmt.Print("Masukkan nama merek: ")
		fmt.Scan(&A[index].brandName)
		fmt.Println("Data telah berhasil diupdate")
	} else if pilihan == "N" || pilihan == "n" {
		pauseAndClear()
	}
}

func deleteBrand(A *tabBrand, n *int) {
	var index int
	var pilihan string
	fmt.Print("Apakah anda ingin menghapus data merek? (Y/N): ")
	fmt.Scan(&pilihan)
	if pilihan == "Y" || pilihan == "y" {
		fmt.Print("Masukkan Indeks yang akan dihapus: ")
		fmt.Scan(&index)
		if index < 0 || index >= *n {
			fmt.Println("Index tidak valid")
			return
		}
		for i := index; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus")
	} else if pilihan == "N" || pilihan == "n" {
		pauseAndClear()
	}
}

func searchProduct(A tabProduct, n int) {
	var keyword string
	fmt.Print("Masukkan nama produk yang dicari: ")
	fmt.Scanln(&keyword) 

	
	fmt.Println("Input yang dimasukkan:", keyword)

	found := false
	for i := 0; i < n; i++ {
		if A[i].productName == keyword { 
			fmt.Printf("%2s %25s %10s %5s %2s %5d\n",
			A[i].productId, A[i].productName, A[i].jenisElektronik, A[i].productYear, A[i].brandId, A[i].jumlahPenjualan)
			found = true
		}
	}
	if !found {
		fmt.Println("Produk tidak ditemukan")
	}
	pauseAndClear()
}

func searchBrand(A tabBrand, n int, products tabProduct, nProduct int) {
	var keyword string
	fmt.Print("Masukkan nama merek yang dicari: ")
	fmt.Scanln(&keyword)
	fmt.Println("Input yang dimasukkan:", keyword)
	found := false

	for i := 0; i < n; i++ {
		if A[i].brandName == keyword {
			fmt.Println("Merek ditemukan:")
			fmt.Println("Brand ID:", A[i].brandId)
			fmt.Println("Brand Name:", A[i].brandName)
			fmt.Println("Jumlah Product:", A[i].jumlahProduct)
			fmt.Println("Produk terkait:")
			for j := 0; j < nProduct; j++ {
				if products[j].brandId == A[i].brandId {
					fmt.Printf("%2s %25s %10s %5s %5d\n",
						products[j].productId, products[j].productName, products[j].jenisElektronik, products[j].productYear, products[j].jumlahPenjualan)
				}
			}
			found = true
			
			
		}
	}

	if !found {
		fmt.Println("Merek tidak ditemukan")
	}
	pauseAndClear()
}


func isBrandExist(A tabBrand, n int, brandId string) bool {
	for i := 0; i < n; i++ {
		if A[i].brandId == brandId {
			return true
		}
	}
	return false
}

func updateJumlahProduct(products tabProduct, n int, brands *tabBrand, m int) {
	for i := 0; i < m; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if products[j].brandId == brands[i].brandId {
				count++
			}
		}
		brands[i].jumlahProduct = count
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func pauseAndClear() {
	fmt.Println("Tekan Enter untuk melanjutkan...")
	var dummy string
	fmt.Scanln(&dummy)
	ClearScreen()
}


func top3BySales(products tabProduct, nProduct int, brands tabBrand, nBrand int) {
	
	insertionSortProductBySales(&products, nProduct)
	
	//insertionSortBrandBySales(&brands, nBrand)

	fmt.Println("Top 3 Products by Sales:")
	for i := 0; i < 3 && i < nProduct; i++ {
		fmt.Printf("%d. %s (Sales: %d)\n", i+1, products[i].productName, products[i].jumlahPenjualan)
	}
	

	fmt.Println("\nTop 3 Brands by Sales:")
	for i := 0; i < 3 && i < nBrand; i++ {
		fmt.Printf("%d. %s (Sales: %d)\n", i+1, brands[i].brandName, products[i].jumlahPenjualan)
	}
	pauseAndClear()
}


func fillDummyData(products *tabProduct, nProduct *int, brands *tabBrand, nBrand *int) {
	// Dummy Brands
	brands[0] = Brand{"B001", "Samsung", 0}
	brands[1] = Brand{"B002", "Apple", 0}
	brands[2] = Brand{"B003", "Acer", 0}
	*nBrand = 3

	// Dummy Products
	products[0] = Product{"P001", "Samsung_43CU7000", "TV", "2020", "B001", 100}
	products[1] = Product{"P002", "MacBook_Pro_M1", "Laptop", "2021", "B002", 200}
	products[2] = Product{"P003", "Acer_Nitro_5", "Smartphone", "2022", "B003", 300}
	products[3] = Product{"P004", "Samsung_Galaxy_Tab_S9_FE", "Tablet", "2023", "B001", 150}
	products[4] = Product{"P005", "Apple_Watch_SE_2022", "Smartwatch", "2024", "B002", 250}
	*nProduct = 5

	updateJumlahProduct(*products, *nProduct, brands, *nBrand)
}
