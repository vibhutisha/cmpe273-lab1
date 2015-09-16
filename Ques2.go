package main

import ("fmt";"math")

type Shape interface {
perimeter() float64
}

func totalPerimeter(shapes ...Shape) float64 {
var perimeter float64
for _, p := range shapes {
perimeter += p.perimeter()
}
return perimeter
}


type Circle struct {
r float64
}

func (c *Circle) perimeter() float64 {
return  2 * math.Pi * c.r
}

type Rectangle struct {
l, w float64
}

func (r *Rectangle) perimeter() float64 {
return (2 * (r.l + r.w))

}

func main() {
    c := Circle{20}
    r := Rectangle{30,50}
    fmt.Println("Perimeter of Circle:",c.perimeter())
    fmt.Println("Perimeter of Rectangle:",r.perimeter())
    fmt.Println("Perimeter of Shape:",totalPerimeter(&c, &r))
}
