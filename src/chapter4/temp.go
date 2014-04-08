package main 
import "fmt"
func main() {
	fmt.Println("Enter Temp in Fahrenheit : ")
	var ft float64
	fmt.Scanf("%f", &ft)
	ct := ((ft - 32) * 5/9)
	fmt.Println("Temp in Celsius : ", ct)
}