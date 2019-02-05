package main
import (
	"fmt"
	"math"

	// "math"
)

func main1(){
	var sum int = 50
	sum += 25
	fmt.Println(sum)
}

func main2(){
	var sum int
	sum = 50
	fmt.Println(sum)
}

func main(){
	var val float64
	val = 4
	var val_res = math.Pow(val,2)
	fmt.Println(val_res)
}