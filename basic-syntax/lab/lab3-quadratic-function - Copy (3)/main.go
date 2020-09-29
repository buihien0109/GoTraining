package main
 
import (
   "fmt"
   "math"
)
 
func quadratic(b float64, c float64) {
 
   fmt.Println("b : ", b)
   fmt.Println("c : ", c)

 
   discriminant := b*b - 4.0*c
   d := math.Sqrt(discriminant)
 
   fmt.Printf("%0.6f\n", ((-b + d) / 2.0))
   fmt.Printf("%0.6f\n", ((-b - d) / 2.0))  
}
 
func main(){
   quadratic(-1.0,-1.0)
}