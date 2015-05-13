package main

import (
	"fmt"
)

func main(){
	
	//set up termbox
	fmt.Println("START")
	r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    measure(r)
    measure(c)
}