# Defer & Exit

Keyword `defer` digunakan untuk mengakhirkan eksekusi sebuah statement di paling akhirtepat sebelum block function selesai.
Sedangkan `exit` digunakan untuk menghentikan program secara paksa, (menghentikan program tidak sama seperti `return` yang hanya menghentikan block code).

## penerapa keyword `defer`

`defer` bisa ditempatkan di mana saja, di awal maupun di akhir block, tetap tidak mempengaruhi kapan waktunya di eksekusi, akan selalu di eksekusi paling akhir.

contoh penggunaakn keyword `defer` :

```go
package main

import "fmt"

func main(){
  defer fmt.Println("World")
  fmt.Println("Hello")
}
```

program di atas saat dijalankan hasilnya akan seperti berikut :

![contoh sederhana penggunakan keyword defer][defer]

keyword `defer` di atas akan mengakhirkan statement `fmt.Println("World")`, efeknya kata _Hallo_ muncul pertama kali.

Meskipun statement yang di-`defer` akan mengakhirkan eksekusinya, tapi statement tersebut masih akan tetap dijalankan meskipun block kode diberhentikan ditengah menggunakan keyword `return`, statement yang di-`defer` akan tetap dijalankan.

```go
package main

import "fmt"

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
  fmt.Println("pesanana anda :", menu)
}

func main(){
  var makanan = [2]string{"pizza", "burger"}
	OrderSomeFood(makanan[0])
	OrderSomeFood(makanan[1])
}
```

![hasil penggunaan keyword defer di statement yang diberhentikan menggunakan return][return]

pada contoh hasil di atas ketika kondisi pertama bernilai `true` A.K.A menu bernilai _"pizza"_, statement `fmt.Println("pesanana anda :", menu)` tidak dijalankan,
karena kode diberhentikan di tengah jalan menggunakan keyword `return`.
Tapi satatement `defer fmt.Println("selamat menikmati hidangan kami")` masih tetap dijalankan, karena penempatannya berada sebelum keyword `return` didefinisikan


eksekusi statement `defer` adalah di akhir block function, bukan block lainnya, seperti block seleksi kondisi.

```go
package main

import "fmt"

func main(){
  if true{
    fmt.Println("Hello 1")
    defer fmt.Println("Hello 3")
  }
  fmt.Println("Hello 2")
}
```

pada contoh di atas, `hallo 3` akan tetap di print setelah `hallo 2`,
meskipun statement defer dipergunakan dalam seleksi block kondiis if.
Hal ini karena defer eksekusinya terjadi pada akhir block function (dalam contoh diatas main), bukan pada akhir block `if`

agar `Hallo 3` bisa muncul di akhir block `if` maka harus dibungkus kedalam function closure

```go
package main

import "fmt"

func main(){
  if true {
		fmt.Println("Hallo 1")
		func() {
			defer fmt.Println("Hallo 3")
		}()
	}
	fmt.Println("Hallo 2")
}
```

## penerapa function `os.Exit()`

Exit digunakan untuk menghentikan program secara paksa pada saat itu juga (pada saat pemanggilan function `os.Exit()`), semua statement setelah exit tidak akan di eksekusi, termasuk juga defer yang berada sebelum function `os.Exit()` dipanggil.

function `os.Exit()` berada dalam package `os`, function ini memiliki parameter bertipe numeric yang wajib diisi, angka yang dimasukkan akan muncul sebagai pesan **exit status** ketika program berhenti.

contoh penerapan :

```go
package main

import (
  "fmt"
  "os"
)

func main(){
  defer fmt.Println("Hello World")
  os.Exit(1)
  fmt.Println("Selamat Datang")
}
```

![hasil penggunaan sederhana function os.Exit()][exit]

[defer]:./img/defer.jpg
[return]:./img/return.jpg
[exit]:./img/exit.jpg