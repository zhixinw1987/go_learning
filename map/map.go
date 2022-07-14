package main

import (
	"fmt"
)

func addColor(m map[string]string, key string, val string){
	m[key] = val
}

func delColor(m map[string]string, key string){
	delete(m, key)
}

func showMapDemo (){
	colors := initMap()
	delColor(colors, "black")
	addColor(colors, "blue", "#0000FF")
	iterateAll(colors)
}

func initMap () map[string]string{
	// var colors map[string]string
	// colors := make(map[string]string)
	//above are 2 equal way to declare empty map
	//below is to declare a map with init value
	return map[string]string {
		"red": "#FF0000",
		"black": "#000000",
		"silver": "#C0C0C0",
		"green": "#008000",
	}
}

func iterateAll (m  map[string]string){
	for key, val := range m {
		fmt.Println("the hex value of color", key, "is", val )
	}
}