package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"unicode/utf8"
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
	
	for i<3{
		i+=1
		fmt.Printf("while loop alter:%v",i)
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

func interesting_things(s1, s2 string)(length1, length2 int8){
	// len() returns the number of bytes but not the characters
	fmt.Println("---Intersting Things---")
	result,error_check:=error_catch(1,0)	
	result2,error_check2:=error_catch(1,2)

	fmt.Printf("Error Check: %v,%v & Result: %v %v",error_check,error_check2,result,result2)
	fmt.Println("\n\n---Intersting Things with strings---")
	var myrune []rune=[]rune("helloγ")
	norune:="helloγ"
	var myrune2 = 'γ' //rune character
	fmt.Println("Rune Single character",myrune2)
	fmt.Println("Data:",myrune,"Rune:",len(myrune),"NoRune:",len(norune))

	return int8(len(s1)),int8(utf8.RuneCountInString(s2))
}

func error_catch(i1, i2 int)(int, error){
	// built in error handling and returns nil when there are no errors, 
	// useful with external packages
	var err error
	if i2==0{
		err=errors.New("divide By zero error")
		return 0, err
	}
	return i1/i2, err

}

func arrays_slices(){
	fmt.Println("\n\n---Arrays/Slices---")
	// arrays vs slices - slices have dynamic sizes and are more common
	array1 := [3]int{1,2,3}
	array2 := [...]int8{1,2,3}
 	slice1 := []int{1,2,3}
	var slice2 []int= []int{123,45,2} 
	var slice3 []int
	slice4:=make([]int32,3,8) //length 3 and capacity 8
	fmt.Println("Length,Capacity of Slice1:",len(slice1),cap(slice1))
	slice1=append(slice1,4)
	slice3=append(slice3,slice2...)
	fmt.Println(array1,array2,slice1,slice3)
	fmt.Println("Length,Capacity of Slice1:",len(slice1),cap(slice1))
	fmt.Println("Data,Length,Capacity of Slice3:",slice3,len(slice4),cap(slice4))
	// Iterate Maps
	for index, value := range slice2{
		fmt.Println("Index:",index,",Value:",value)
	}
	maps()
}

func maps(){
	fmt.Println("\n---Maps---")

	var myMap map[string]int = make(map[string]int) // make can be replace with a literal at end {}
	myMap2:=map[string]uint{"Alice":15,"Bob":20}
	fmt.Println("All Maps:",myMap2,myMap,myMap2["data"])
	data,ok:=myMap2["data3"]
	fmt.Println("No data and boolean:",data,ok)
	myMap2["Adam"]=14
	// delete by ref
	delete(myMap2,"data2")
	fmt.Println("Append and Delete:",myMap2)
	// Iterate maps
	for name,age:=range myMap2{
		fmt.Println("Name:",name,",Age:",age)
	}
}