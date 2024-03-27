package main

import "fmt"

type houseSize struct{
	roomName string
	roomSize uint16
	roomOwner
}
type hotelSize struct{
	hotelRoomName string
	hotelRoomSize uint16
}
func (e hotelSize) remainingSpace() uint32{
	return uint32(1800-e.hotelRoomSize)
}
type roomOwner struct{
	ownerName string
}
type house interface{
	remainingSpace() uint32
}
func (e houseSize) remainingSpace() uint32{
	return uint32(1500-e.roomSize)
}
func spaceAllocated(e house, spaceRequired int){
	if e.remainingSpace()<uint32(spaceRequired){
		fmt.Print("No space Left\n")
	}else{
		fmt.Printf("Space is there!\n")
	}
}
func houseDetails(){
	fmt.Printf("\n---Structs---\n")
	var myHouse houseSize =houseSize{"Living Room",1200,roomOwner{"Party-A"}}
	var myHotel hotelSize = hotelSize{"Reception",1500}
	fmt.Printf("Room Owner: %v, Room name: %v, Room Size:%v, Remaining Space:%v",myHouse.ownerName,myHouse.roomName,myHouse.roomSize,myHouse.remainingSpace())
	fmt.Println("")// shell printing % everytime I am using the printf, to solve this we need to add a new line here
	spaceAllocated(myHouse,300)
	fmt.Printf("HotelRoom Name: %v, Hotel Room Size:%v, Remaining Space:%v\n",myHotel.hotelRoomName,myHotel.hotelRoomSize,myHotel.remainingSpace())
	spaceAllocated(myHotel,300)

}

func pointer(){
	fmt.Printf("\n---Pointers---\n")

	var p *int32 = new(int32)
	var i int32
	i=5
	*p=10
	fmt.Printf("I: %v, P:%v, address Of I:%v,address Of P: %v\n", i,*p, &i,p)
	p=&i
	fmt.Printf("I: %v, P:%v, address Of I:%v,address Of P: %v\n", i,*p, &i,p)
	// overwrite the i value with p (de ref value)
	*p=3
	fmt.Printf("I: %v, P:%v, address Of I:%v,address Of P: %v\n", i,*p, &i,p)
	slice:=[]int32{1,2,3}
	var slice2 *[]int32=&slice //stores the reference of slice which saves memory but oberwrites the slice (chanes in slice2)
	fmt.Printf("The address of slice,slice2:%p, %p\n values of slice,slice2:%v, %v",&slice,slice2,slice,*slice2)


}