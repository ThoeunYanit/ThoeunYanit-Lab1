package main
import "fmt"

func main(){
	var x int
	var y int

	fmt.Print("Input two integers: ")
	// fmt.Scanln(&x, &y)
	fmt.Scanf("%d %d", &x, &y)

	if x>0 && y>0 {
		fmt.Println(true)
	} else if x>1 || y>1 {
		fmt.Println(true)
	} else if x !=y {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}