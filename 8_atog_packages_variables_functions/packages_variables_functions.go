// 1. Go programs made up of packages

// 2. Programs start running in "main" package
package main

// 3. Importing fmt and math/rand packages. Package name is the same as the last element of import path
// 4. fmt = package fmt
// 5. math/rand = package rand
// 7. The following is called "factored" import statement. Remember you can also import 
// each individually but factored import is generally better style
import(
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

// 9. function takes zero or more arguments
// type comes after variable name and function definition
func add(x int, y int) int {
	return x + y
}

// 11. when you have 2 consecutative paramts of the same type, you can
func subtract(x, y int) int {
	return x - y
}

// 13. return multiple things
func swap(s1, s2 string) (string, string) {
	return s2, s1
}

// 15. Return values can be named. If you return blank, it will return the named values in the function 
// definition. But don't do the latter. The former is good for documenting
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
// 17 package level variables 
var c, python, java bool

// 19. You can initialize the above, one per variable on declaration
var j, k int = 1, 2

// 23. declaring variables using "factored" style (used for importing)
// complex is cool
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// 6. rand.Seed(15). Printf vs Println. Think c 
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// 8. In packages, you can only refer to exported names (which begin with capital letter)
	fmt.Println(math.Pi)
	// 10. Calling add function
	fmt.Println(add(5, 7))
	// 12. Calling subtract function
	fmt.Println(subtract(12, 5))
	// 14. Calling swap method
	fmt.Println(swap("wq", "bq"))
	// 16. Calling split method
	fmt.Println(split(10))
	// 18. function level variable
	var i int
	fmt.Println(i, c, python, java)
	// 20. Variable will take up type of initializer if not defined
	var cc, ppython, jjava = true, false, "no!"
	fmt.Println(j, k, cc, ppython, jjava)
	// 21. Inside a function you can use shorthand assignment for implicit declaration type. 
	// Cannot do this outside of function
	l := 3
	fmt.Println(i, j, k, l, c, python, java)
	// 22. Go Basic Types
	// bool, string, int -> int64, uint -> uint64, byte, rune, float32, float64, complex64, complex128
	// 24. printing the vars out
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// 25 Variables declared without explicit initial value are given zero value
	var a int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v, %v, %v, %q\n", a, f, b, s)
	// 26. Casting is a option
	var xx, yy int = 3, 4
	var ff float64 = math.Sqrt(float64(xx*xx + yy*yy))
	var zz uint = uint(ff)
	fmt.Println(xx, yy, zz)
	// 27.  Type inference
	var i_inf int
	j_inf := i // j is an int
	fmt.Println(i_inf, j_inf)
	// 28. Based on value int given, it will either be int, float or complex
	// 29. Constants
	const World = "world"
	fmt.Println("hello", World)
	// Constants cannot be created with :=
	// 30. Numeric constants are high precision values (not sure what this mean, maybe it can handle
	// all numeric types big or small? Not sure...)
	// Similarly, untyped constants will take from its context
}