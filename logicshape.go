package main
import "math"
import "fmt"
type Shape interface {
	Area() float64
	perimeter() float64
}

func Measure_Area(s Shape)float64 {
return s.Area()
}
func measure_Perimeter(s Shape) float64{
return s.perimeter()
}

func Distance(x1,y1,x2,y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
	}
type Circle struct {
	x float64
	y float64
	r float64
	}

type Rectangle struct{
	x1,y1,x2,y2 float64
	}

func (r Rectangle) Area() float64 {
	l:=Distance(r.x1,r.y1,r.x1,r.y2)
	w:=Distance(r.x1,r.y1,r.x2,r.y1)
	return l*w
}
func (r Rectangle) perimeter() float64 {
	l:=Distance(r.x1,r.y1,r.x1,r.y2)
	w:=Distance(r.x1,r.y1,r.x2,r.y1)
	return 2*(l+w);

}
func (c Circle) Area() float64 {
    return math.Pi * c.r * c.r
}
func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}
	 
func Shape_caller(){
var x_axis,y_axis,radius float64;
var xcor1,ycor1,xcor2,ycor2 float64;
fmt.Println("Enter x_axis y_axis radius for circle and x1,y1,x2,y2 for rectangle")
fmt.Scanf("%f %f %f %f %f %f %f",&x_axis,&y_axis,&radius,&xcor1,&ycor1,&xcor2,&ycor2)
c:= Circle{x_axis,y_axis,radius}
fmt.Println("%%","Area of circle",Measure_Area(c))
fmt.Println("%%","Perimeter of circle",measure_Perimeter(c));
    r := Rectangle{xcor1,ycor1,xcor2,ycor2}
fmt.Println("%%","Area of rectangle",Measure_Area(r))
fmt.Println("%%","Perimeter of rectangle",measure_Perimeter(r));
    }