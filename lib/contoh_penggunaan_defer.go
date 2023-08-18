package lib

import (
	"fmt"
	"time"
)

func SendData(ch chan<- string) {
	for i := 0; i < 5; i++ {
		now := time.Now()
		x, x1, x2 := now.Clock()
		ch <- fmt.Sprintf("%d:%d:%d:%d", x, x1, x2, now.Nanosecond()/1000000)
	}
	close(ch)
}

func ReceiveData(ch <-chan string) {
	for data := range ch {
		fmt.Println(data)
		time.Sleep(time.Duration(time.Second * 2))
	}
}

func OrderSomeFood(menu string) {
	// statement yang di defer akan tetap muncul meskipun block kode diberhentikan dengan menggunakan keyword return
	defer fmt.Println("selamat menikmati hidangan kami")

	// menggunakan switch case dan if else disini maksudnya sama saja

	// ketika banyak statement yang di-defer, maka seluruhnya akan di eksekusi di akhir secara berurutan
	switch menu {
	case "pizza":
		fmt.Printf("Pilihan tepat!, ")
		fmt.Println("Pizza di tempat kami paling enak!")
		return
	case "burger":
		fmt.Print("Pilihan tepat!, ")
		fmt.Println("Burger di tempat kami paling enak!")
	}

	// if menu == "pizza" {
	// 	fmt.Printf("Pilihan tepat!, ")
	// 	fmt.Println("Pizza di tempat kami paling enak!")
	// 	return
	// } else if menu == "burger" {
	// 	fmt.Print("Pilihan tepat!, ")
	// 	fmt.Println("Burger di tempat kami paling enak!")
	// }

	fmt.Println("pesanan anda :", menu)
}
