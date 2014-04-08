package main 
import "fmt"
func main() {
	fmt.Println("Enter first Number : ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}