//2311102223 package main 
 
import ( 
    "fmt" 
    "strconv" 
) 
 
func main() {     var input int     fmt.Print("Masukkan bilangan bulat positif lebih besar dari 100: ")     fmt.Scanln(&input) 
 
    if input <= 100 { 
        fmt.Println("Bilangan yang anda masukkan tidak valid. Masukkan bilangan bulat positif lebih besar dari 100.") 
        return 
    } 
 
    numberString := strconv.Itoa(input)     length := len(numberString) 
 
    var part1, part2, part3 string     if length%2 == 0 { 
         
        part1 = numberString[:length/3]         part2 = numberString[length/3 : length*2/3] 
        part3 = numberString[length*2/3:] 
    } else { 
         
        part1 = numberString[:(length+1)/3]         part2 = numberString[(length+1)/3 : (length+1)*2/3]         part3 = numberString[(length+1)*2/3:] 
    } 
 
    intPart1, _ := strconv.Atoi(part1)     intPart2, _ := strconv.Atoi(part2)     intPart3, _ := strconv.Atoi(part3) 
 
    fmt.Println(intPart1, intPart2, intPart3)     fmt.Printf("%.2f\n", float64(intPart1+intPart2+intPart3)/3) 
} 
 
 
