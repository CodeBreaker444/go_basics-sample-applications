package main

import (
	"fmt"
	"math"
	"math/rand"
)

var output, output2 bool
var i, j int = 1, 2 //i:=2 is not allowed outside a function
const pi=3.14 // cannot be declared using :=
const (
	b=1<<80
	s=b>>79
)
func add(x int, y int) int {
	fmt.Println("Hello, from Level-1")
	return x+y
}
func swap(x, y string) (string, string) {
	return y, x
}

func random_pair(seed int64)(x,y int64){
	rand.Seed(seed) //deprecated
	//seed
	
	x = rand.Int63()
	y = rand.Int63()
	return
}

func same_type(x,y int) bool{
	if x==y{
		output = true
	}else{
		output2 = false
	}
	if output == output2{
		return true
	}else{
		return false
	}
}

func default_values(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("i=%v, f=%v, b=%v, s=%q\n",i,f,b,s)
}
func type_conversion(){
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x+y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)
}

func binary(num int)int{
	bin := num<<2
	return bin
}

func loops(num int){
	for i:=1;i<10;i++{
		fmt.Printf("%v * %v =%v\n",num,i,num*i)
	}
	
	for i<20{
		i+=1
		fmt.Println(i)
	}

for{
	fmt.Println("loops forever")
	break
}
}

func conditions_compare(i,j int)(same bool){
	if i==j{
	 same=true
	 
	}else{
		same=false
	}
	return same
	

}