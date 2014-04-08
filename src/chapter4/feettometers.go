package main 
import "fmt"
func main() {
	fmt.Println("Enter vaule in feets: ")
	var feets float64
	fmt.Scanf("%f", &feets)
	meters := (feets * 0.3048)
	fmt.Println("Value in meters : ", meters)
}