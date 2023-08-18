package main

import (
	"fmt"
	"golang/defer-dan-exit/lib"
	"os"
	"time"
)

func main() {
	now := time.Now()
	waktu := make(chan string)

	// penggunaan keyword defer
	defer lib.ReceiveData(waktu) // statement ini akan dijalankan paling akhir
	go lib.SendData(waktu)

	// contoh penggunaakn keyword defer di return
	var makanan = [2]string{"pizza", "burger"}
	lib.OrderSomeFood(makanan[0])
	lib.OrderSomeFood(makanan[1])

	end := time.Now()
	fmt.Printf("program berjalan selama : %v\n", end.Sub(now))

	// eksekusi defer adalah di akhir block function, buka block lainnya, seperti block seleksi kondisi

	// if true {
	// 	fmt.Println("Hallo 1")
	// 	defer fmt.Println("Hallo 3")
	// }
	// fmt.Println("Hallo 2")
	/*
		pada contoh di atas, `hallo 3` akan tetap di print setelah `hallo 2`,
		meskipun statement defer dipergunakan dalam seleksi block kondiis if.
		Hal ini karena defer eksekusinya terjadi pada akhir block function (dalam contoh diatas main), bukan pada akhir block if
	*/

	// agar block `Hallo 3` muncul di akhir block if
	// maka dibungkus kedalam closure
	fmt.Println()
	if true {
		fmt.Println("Hallo 1")
		func() {
			defer fmt.Println("Hallo 3")
		}()
	}
	fmt.Println("Hallo 2")

	// penerapan function `os.Exit()`
	defer fmt.Println("Hello World")
	os.Exit(1)
	fmt.Println("Selamat Datang")
}

/*
	pada statement defer lib.ReceiveData(waktu), berjalan di akhir (paling akhir) block function
	karena pada akhir block function terdapat pemanggilan function os.Exit(),
	maka statement defer lib.ReceiveData(waktu) tidak akan di eksekusi,
	karena di tengah tengah function, program dihentikan secara paksa.
*/
