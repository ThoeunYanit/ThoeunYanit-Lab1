package main
import "fmt"

func main(){
	var x int
	var y int

	fmt.Printf("Enter value of x and y: ")
	fmt.Scanln(&x, &y)
	
	var result int = x + y
	fmt.Printf("x + y = %d\n", result)

	x += y
	fmt.Printf("x += y : %d\n", x)

	x -= y
	fmt.Printf("x -= y : %d\n", x)

	x *= y
	fmt.Printf("x *= y : %d\n", x)

	x /= y
	fmt.Printf("x /= y : %d\n", x)

	x %= y
	fmt.Printf("x %%= y : %d\n", x)
	
}