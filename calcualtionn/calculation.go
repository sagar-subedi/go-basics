package calculationn

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func IsOdd(num int) bool {
	return num%2 != 0
}

func PrintDay(num int) {
	switch num {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Print("Invalid num provided")
	}

}

func Factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	result := 1

	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func PrintArray(names []string) {
	for idx, name := range names {
		fmt.Println(idx, name)
	}
}
