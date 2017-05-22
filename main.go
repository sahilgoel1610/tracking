package main

import "fmt"
import "./lib"
// import "./lib/aerospike"
import "./lib/mem_store_factory"
// import "./lib/memory_store/aerospike"


type Point struct {
	lat float64
	long float64
}

type Snapshot struct {
	point Point
	createdAt string
	userId int
}


func main() {
	fmt.Println("Hello")
	p := Point {123.123, 12.213}
	fmt.Println(p.lat)
	ss := Snapshot {p, "12312313", 1}
	fmt.Println(ss.point.long)

	u1 := users.New()
	fmt.Println(u1.Id)
	u2 := users.New()
	fmt.Println(u2.Id)
	(&u1).AddFriend(&u2)
	fmt.Println(u1.Friends)
	memStore := memStoreFactory.Get("redis")
	// memStore := aerospike.GetAeroClient()
	fmt.Println(memStore.AddPair("hey", "bye"))

}