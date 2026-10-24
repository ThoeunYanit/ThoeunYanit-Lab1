package main

import "fmt"

func Add(a, b int) int {
	
	return a+b
}

func Substract(a, b int) int {
	
	return a-b
}

func Multiply(a, b int) int {
	
	return a*b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Integer divide by zero")
	}
	return a/b, nil
}

func Modulo(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Integer divide by zero")
	}
	return a%b, nil
}

func takeInput () (int, int, error){
	var a int
	var b int
	var err error

	fmt.Print("Enter a: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		
		var discard string
		fmt.Scanln(&discard)
		return 0,0, err
	} 

	fmt.Print("Enter b: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		
		var discard string
		fmt.Scanln(&discard)
		return 0,0, err
	} 

	return a,b,nil
}

func main(){
	var choice int

	for choice != 6 {

		fmt.Println("\n===== Mini Calculator =====\n")
		fmt.Println("1). Add")
		fmt.Println("2). Sub")
		fmt.Println("3). Mul")
		fmt.Println("4). Div")
		fmt.Println("5). Mod")
		fmt.Println("6). Exit")
		fmt.Println("")

		fmt.Print("Choose: ")

		_,err := fmt.Scanln(&choice)
		// Check input , if not integer, throw error
		if err != nil {
			fmt.Println("Expect Integer! Please try again!")
			var discard string // discard take input from the &choice, since choice expect integer so choice input expect someone to take it
			fmt.Scanln(&discard) //discard will be freed automatically
			
			continue //skip
		} 
		
		fmt.Println("")

		switch choice {
			case 1:
				a,b, err := takeInput()

				if err != nil { //skip if get error
					fmt.Println("Error:",err)
					continue
				}

				result := Add(a,b)

				fmt.Printf("\n%d + %d = %d\n", a,b,result)

			case 2:
				a,b, err := takeInput()

				if err != nil { //skip if get error
					fmt.Println("Error:",err)
					continue
				}

				result := Substract(a,b)

				fmt.Printf("\n%d - %d = %d\n", a,b,result)

			case 3:
				a,b, err := takeInput()

				if err != nil { //skip if get error
					fmt.Println("Error:",err)
					continue
				}

				result := Multiply(a,b)

				fmt.Printf("\n%d * %d = %d\n", a,b,result)

			case 4:
				a,b, err := takeInput()

				if err != nil { //skip if get error while input
					fmt.Println("Error:",err)
					continue
				}

				var result int
				result, err = Divide(a,b)

				if err != nil { //skip if get error after getting result
					fmt.Println("Error:",err)
					continue
				}

				fmt.Printf("\n%d / %d = %d\n", a,b,result)

			case 5:
				a,b, err := takeInput()

				if err != nil { //skip if get error while input
					fmt.Println("Error:",err)
					continue
				}

				var result int
				result, err = Modulo(a,b)

				if err != nil { //skip if get error after getting result
					fmt.Println("Error:",err)
					continue
				}

				fmt.Printf("%d %% %d = %d\n", a,b,result)

			case 6:
				break
				
			default:
				fmt.Println("Invalid Input! Try Again!")
		}
	}
}