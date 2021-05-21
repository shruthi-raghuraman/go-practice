package main

//Go is made up of packages and programs start running main
//Program uses packages with import paths in the factored import statement
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Function with 2 params with type after variable name
func add(x, y int) int {
	return x + y
}

//Function can return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

//Function with return statement without arguments
func split(sum int) (x, y int)  {
	x = sum * 4 / 9
	y = sum -x
	return
}

//var statement at package level, list of variables that are bool
var c, python, java bool

func main()  {
	//Print String
	fmt.Println("Hello")

	//Print current time with time package
	fmt.Println("Time is: ", time.Now())

	//Print random number with seed
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number: ", rand.Intn(10))

	//Exported names begin with a capital letter
	fmt.Println(math.Pi)

	//Function call
	fmt.Println(add(8,3))

	//Function that swaps words and returns multiple results
	a, b := swap("hello", "world")
	fmt.Println(a,b)

	//Function with a naked return
	fmt.Println(split(20))

	//var statement at the function level
	var i, j int
	fmt.Println(i, j, c, python, java)
}
