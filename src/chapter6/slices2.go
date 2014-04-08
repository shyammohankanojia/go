package main 
import "fmt"
func main() {
	// The make function also allows a 3rd parameter:
	x := make([]float64, 5, 10)
	// 10 represents the capacity of the underlying array
 	//  which the slice points to:
	fmt.Println(x)

	// Another way to create slices is to use the [low : high]
	// expression:
	arr := []float64{1,2,3,4,5}
	y := arr[0:5]
	// low is the index of where to start the slice and high is
	// the index where to end it

	fmt.Println(y)

}