package main 
import "fmt"
func main() {
	// Here's an example of a slice:
	var x []float64
	// The only difference between this and an array is the
	// missing length between the brackets.
	fmt.Println(x)

	// built-in make function:
	y := make([]float64, 5)
	fmt.Println(y)
}