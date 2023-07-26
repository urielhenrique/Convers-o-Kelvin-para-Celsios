package main

import "fmt"

const ebulicaoK = 373.0

func main() {
	 
	 tempK := ebulicaoK        
	 tempC := (tempK - 273) 
	
	fmt.Printf("A temperatura de ebulição da água em °F = %g  e a temperatura de ebulição da água em °C = %g  ", tempK, tempC)
	
}
