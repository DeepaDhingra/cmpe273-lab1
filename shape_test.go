package main
import "fmt"
import "testing"
func TestShape(t *testing.T) { 
c:= Circle{0,0,10}
r := Rectangle{10,10,15,25}
var v float64
v = Measure_Area(r) 
if v != 75 { 
t.Error("TEST CASE 1 :Expected 75, got ", v) 
}else{
 fmt.Println("TEST CASE 1 ((Area of Rectangle is correct) Passed")
}
v = Measure_Area(c) 
if v != 314.1592653589793 { 
t.Error("TEST CASE 2 :Expected 314.1592653589793, got ", v) 
}else{
 fmt.Println("TEST CASE 2 ((Area of Circle is correct) Passed")
}
v = measure_Perimeter(r) 
if v != 40 { 
t.Error("TEST CASE 3 :Expected 40, got ", v) 
}else{
 fmt.Println("TEST CASE 3 ((Perimeter of Rectangle is correct) Passed")
}
v = measure_Perimeter(c) 
if v != 62.83185307179586 { 
t.Error("TEST CASE 4 :Expected 62.83185307179586, got ", v) 
}else{
 fmt.Println("TEST CASE 4 ((Perimeter of Circle is correct) Passed")
}


}
