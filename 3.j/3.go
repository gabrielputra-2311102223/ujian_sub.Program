package main 
 
import "fmt" 
 
func divide(n, m int) (int, int) {     if n < m {         return 0, n 
    } 
    q, r := divide(n-m, m)     return q + 1, r 
} 
 
func main() {     var n, m int 
 
    fmt.Print("Masukkan bilangan pembilang (n): ")     fmt.Scan(&n)     fmt.Print("Masukkan bilangan penyebut (m): ")     fmt.Scan(&m) 
 
    quotient, remainder := divide(n, m) 
 
    fmt.Printf("Hasil dari pembagian %d dengan %d adalah: %d\n", n, m, quotient)     fmt.Printf("Sisa dari pembagian adalah: %d\n", remainder) 
} 
