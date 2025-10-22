package main 

import "fmt"

func main(){
	var key int
	temps := make(map[int][]float64)
	example := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	
	for i:=0;i<len(example);i++{
	key = int(example[i])/10 * 10
	temps[key] = append(temps[key],example[i])
	}
	fmt.Println(temps)
}

