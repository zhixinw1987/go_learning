package main

import (
	"fmt"
)

type player struct {
	firstName string
	lastName string
	age int
	//contactInfo contactInfo
	//embeded struct can use same attribute name as the type name, just simple remove the attribute name only declare the type
	//the default attribute name will be applied same as the type name "contactInfo contactInfo" 
	contactInfo
}

type contactInfo struct {
	email string
	postcode string
}

func main (){
	//specify attribute name when assigning the initiate value
	p1:=player{
		age:35, 
		lastName:"wang", 
		firstName:"zhixin",
		contactInfo: contactInfo {
			//whenever using multi-line struct, every single line must end with "," even if it's the last line
			email:"zhixinw1987@gmail.com",
			postcode:"521601",
		},
	}
	fmt.Println(p1.age)
	p1.printPlayer()
	//else all attribute must be assigned by sequence in which the attributes are declared
	p2:=player{"yunqing", "wang", 65, contactInfo{"wangzx870320@gmail.com", "610081"}}
	fmt.Println(p2)
	//declare a player variable p3 with empty initiate value
	var p3 player
	fmt.Println(p3.age)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++=")
	pointerDemo(p1)
}

func pointerDemo(p player){
	p.updateAge()
	p.printPlayer()
}

//receiver func with struct
func (p player) printPlayer (){
	fmt.Printf("%+v", p)
}

func (p *player) updateAge () {
	(*p).age = 40
}