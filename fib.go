package main
import ("fmt")

func Fib(x int64) int64 {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0

	} else if x == 1 {
		return 1
	} else {
		return Fib(x-1) + Fib(x-2)
	}

}

func Fib_caller() {
	var number int64
	var index int64
	fmt.Println("Fibonacci program using recursion Lab1_Question1")
	fmt.Println("Enter the range of series to print")
	fmt.Scanf("%d", &number)
	if number < 0 {
		fmt.Println("OOps negative number")
	} else {
		for index = 0; index <= number; index++ {
			fmt.Println("$$", index, "$$", "Fibonacci is", Fib(index))
		}
	}
}
