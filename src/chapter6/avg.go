package main 
import "fmt"
func main() {
	var score [5]float64
	score[0] = 74
	score[1] = 84
	score[2] = 64
	score[3] = 54
	score[4] = 84

	var total float64 = 0
	for i := 0; i < len(score); i++ {
		total += score[i]
	}
	fmt.Println("Avg score : ", (total / float64(len(score)) ))
}