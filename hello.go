package main

import "fmt"
import "rsc.io/quote"
import (
	"math/rand"
	"math"
	"time"
)

var global_var = "global" // cannot use := outside a function

func main() {
	fmt.Println("Hello, World!")
	variables()
	low_level_dec()
}
func variables(){
	fmt.Println("\n---Variables---\n")
	var a string = "initial"
	fmt.Println(a)

	var b int = 1
	fmt.Println(b)

	var c,d int = 1,2
	fmt.Println(c,d)

	var e = "string"
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
	f="short2"
	fmt.Println(f)

	fmt.Println("---ALL Variables---\n",a,b,c,d,e,f,global_var)
	// formatted print with %v, value of variable
	fmt.Printf("a=%v, b=%v, c=%v, d=%v, e=%v, f=%v, global_var=%v\n",a,b,c,d,e,f,global_var)
	// formatted print with %T, type of variable
	fmt.Printf("a=%T, b=%T, c=%T, d=%T, e=%T, f=%T, global_var=%T\n",a,b,c,d,e,f,global_var)
	fmt.Printf("---END---\n\n")

}
func low_level_dec(){
	fmt.Println("---Low Level Declaration---\n")
	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 255
	var numFour int16 = 32767
	var numFive uint16 = 65535
	var numSix int32 = 2147483647
	var numSeven uint32 = 4294967295
	var numEight int64 = 9223372036854775807
	var numNine uint64 = 18446744073709551615
	var numTen float32 = 3.4028235E+38
	var numEleven float64 = 1.7976931348623157E+308
	var numTwelve complex64 = 3.4028235E+38
	var numThirteen complex128 = 1.7976931348623157E+308
	var numFourteen byte = 255
	var numFifteen rune = 2147483647
	var numSixteen bool = true
	var numSeventeen string = "string"
	var numEighteen error = nil
	fmt.Printf("numOne=%v, numTwo=%v, numThree=%v, numFour=%v, numFive=%v, numSix=%v, numSeven=%v, numEight=%v, numNine=%v, numTen=%v, numEleven=%v, numTwelve=%v, numThirteen=%v, numFourteen=%v, numFifteen=%v, numSixteen=%v, numSeventeen=%v, numEighteen=%v\n",numOne,numTwo,numThree,numFour,numFive,numSix,numSeven,numEight,numNine,numTen,numEleven,numTwelve,numThirteen,numFourteen,numFifteen,numSixteen,numSeventeen,numEighteen)
	fmt.Println("Total of",18,"variables\n")
	fmt.Printf("---END---\n\n")
	packages()

}

func packages(){
	fmt.Println("---Packages---\n")
	fmt.Println(quote.Go())
	//In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
	fmt.Println("My favorite number is", rand.Intn(10),math.Pi)
	fmt.Println("The time is", time.Now())
	fmt.Printf("\n---END---\n\n")
	level1add :=add(1,3)
	fmt.Println("Level-1 add:",level1add)
	level1swap1,level1swap2 :=swap("hello","world")
	fmt.Println("Level-1 swap:",level1swap1,level1swap2)
	level1random1,level1random2 :=random_pair(1)
	fmt.Println("Level-1 random:",level1random1,level1random2)
	level1same :=same_type(1,2)
	fmt.Println("Level-1 same:",level1same)
	default_values()
	type_conversion()
	bin:=binary(35)
	fmt.Println("Binary:",bin)
	loops(3)

}