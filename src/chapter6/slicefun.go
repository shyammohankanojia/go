// Go includes two built-in functions to assist with slices:
// append and copy. 
package main 
import "fmt"
func main() {
	slices1 := []float64{1, 2, 3, 4}
	// slices2 := []float64{7, 5, 8, 9}
	x := append(slices1, 5, 6, 7, 8)
	fmt.Println(x)
	
}
