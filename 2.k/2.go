//2311102223 package main 
 
import ( 
    "fmt" 
    "strconv" 
) 
 
func determinePrize(ticketNumber int) string { 
 
    digits := strconv.Itoa(ticketNumber) 
 
    allSame := true     isEven := (digits[0]-'0')%2 == 0     for _, d := range digits {         if d != digits[0] {             allSame = false             break 
        } 

    } 
 
    if allSame && isEven {         return "Hadiah Utama" 
    } 
 
    allOdd := true     for _, d := range digits {         if (d-'0')%2 == 0 {             allOdd = false             break 
        } 
    } 
 
    if allOdd {         return "Hadiah Sembako" 
    } 
 
    return "Hadiah Konsol" 
} 
 
func main() {     var N int     fmt.Print("Masukkan jumlah peserta anda: ")     fmt.Scan(&N) 
 
    hadiahCount := map[string]int{ 
        "Hadiah Utama":   0, 
        "Hadiah Sembako": 0, 
        "Hadiah Konsol":  0, 
    } 
 
    for i := 0; i < N; i++ {         var ticketNumber int         fmt.Print("Masukkan nomor tiket: ")         fmt.Scan(&ticketNumber) 
 
        prize := determinePrize(ticketNumber)         fmt.Println(prize) 
 
        hadiahCount[prize]++ 
    } 
 
    fmt.Println("Jumlah Hadiah Utama:", hadiahCount["Hadiah Utama"])     fmt.Println("Jumlah Hadiah Sembako:", hadiahCount["Hadiah Sembako"])     fmt.Println("Jumlah Hadiah Konsol:", hadiahCount["Hadiah Konsol"]) 
} 
 
 
