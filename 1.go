package main
import "fmt"

func printNumRecursive(a, b *int) {
	*b += 1
}

func main() {
	var a1, b1 int
	fmt.Scan(&a1)
	
	for b1 < a1 {
		printNumRecursive(&a1, &b1)
		fmt.Print(b1, " ")
	}
}
