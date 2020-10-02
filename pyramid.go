package main
import "fmt"
 
func main() {
    var layers int
    var k int = 0
    fmt.Print("Enter number of rows :")
    fmt.Scan(&layers)     
    for i := 1; i <= layers; i++ {     
        k=0
        for space := 1; space <= layers-i; space++ {
            fmt.Print("  ")         
        }
        for {
            fmt.Print("* ")
            k++
            if(k == 2*i-1){             
                break
            }
        }       
        fmt.Println("")
    }
}
