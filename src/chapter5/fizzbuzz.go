package main 
import "fmt"
func main() {
	for i := 3; i < 100; i++ {
			if (i % 3 ==  0) && (i % 5 == 0)  {
				fmt.Println(i, "is divisible by 3 and 5 => FIZZBUZZ")
			} else if i % 3 == 0{
				fmt.Println(i, "is divisible by 3 => FIZZ")
			} else if i % 5 == 0 {
				fmt.Println(i, "is divisible by 5 => BUZZ")
			}
		}	
}