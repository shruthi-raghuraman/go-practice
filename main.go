package main

//Go is made up of packages and programs start running main
//Program uses packages with import paths in the factored import statement
import (
	"fmt"
	"math"
	"math/cmplx"
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
//No initializer
var c, python, java bool

//var with initializer
var d, e int = 1, 2

//variable statements can be factored into blocks like imports
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

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

	//var with initializer and no type declaration
	var c, python, java = true, false, "gus"
	fmt.Println(d, e, c, python, java)

	//short variable assignment statement without declaration/type
	k := 3
	f, g, h := true, true, false
	fmt.Println(k, f, g, h)

	//Go basic types are bool, string, int, byte, rune, float32, float64
	//complex64, complex128, %T is type and %v is value
	fmt.Println("Basic variable types")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
