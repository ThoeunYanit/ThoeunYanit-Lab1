package main
import "fmt"

func myAND (a int, b int) int {
	return a&b
}

func myOR (a int, b int) int {
	return a|b
}

func myXOR (a int, b int) int {
	return a^b
}

func myNOT (a int, b int) (int,int) {
	return ^a, ^b
}

func myLeftShift (a int, b int) int {
	return a<<b
}

func myRightShift (a int, b int) int {
	return a>>b
}

func main(){
	var a int
	var b int

	fmt.Print("Input integer a: ")
	fmt.Scanln( &a)
	fmt.Print("Input integer b: ")
	fmt.Scanln(&b)
	
	fmt.Println(myAND(a,b))
	fmt.Println(myOR(a,b))
	fmt.Println(myXOR(a,b))
	fmt.Println(myNOT(a,b))
	fmt.Println(myLeftShift(a,b))
	fmt.Println(myRightShift(a,b))
}