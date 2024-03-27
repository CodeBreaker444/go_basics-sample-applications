package main
import (
	"fmt"
	"sync"
	"time"
)

var global_data map[int]int32= map[int]int32{}
var swg = sync.WaitGroup{}
var m= sync.RWMutex{}
func coroutines(){
	fmt.Println("\n\n---Co Routines---")
	

	t0 := time.Now()
	data:=[]int32{1,2,3,4,5,6}
	var j int32
	for i,j = range data{
		swg.Add(1)
		go co_call_write(i,j) // sometime the program crashes due to concurrent writes to the map
	}
	swg.Wait()
	fmt.Println("\nGlobal Data:",global_data)

	fmt.Printf("Total Execution Time: %v",time.Since(t0))
}
func co_call_write(i int,j int32){
	time.Sleep(time.Duration(2000)*time.Millisecond)
	m.Lock()
	global_data[i]=j // concurrency crash here , overcome with mutex lock
	m.Unlock()
	co_call_read()// concurrency crash here, overcome with read lock

	swg.Done()
}
func co_call_read(){
	m.RLock()

	fmt.Println("Global Data Read:",global_data)
	m.RUnlock()
}

func main_channel(){
	fmt.Println("\n\n---Channels---")

	var uc = make(chan int) // stores only value at a time
	var bc = make(chan int,10)
	swg.Add(1)
	go add_queue(bc)
	go process_queue(bc)
	go channel_process(uc)

	swg.Wait()
	for i:=range uc{
		fmt.Println(i)
	}

}
func channel_process(uc chan int){
	defer close(uc)
	for i:=0; i<5; i++{
		uc <- i
	}


}
func add_queue(bc chan int){

	for i:=0;i<10;i++{

		bc<-i

	}
	close(bc)


}
func process_queue(bc chan int){
	var j int=0

	for i:=range bc{
		j=j+i
	}
	fmt.Printf("Processed Buffered channel data:%v\n",j)
	swg.Done()

}

func generics[T int | float32](i T) {
	fmt.Println("\n\n---Generics---")

	fmt.Printf("Multi type input squared:%v\n", i*i)
	generics_any([]int32{1,2,3})
	generics_any([]string{"1","2","3","hello"})

}
func generics_any[T any](i []T){
	fmt.Printf("Any type generic:%v\n",len(i))
}