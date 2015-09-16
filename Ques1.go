package main

import "fmt"

//import "Users/Vibhuti/Go/src/golang-book/LAB1/math"

type MemoryFunction func(int, ...int) interface{}

var v MemoryFunction

func init() {
	v = CacheStore(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return v(x-1).(int) + v(x-2).(int)
	})
}
func CacheStore(function MemoryFunction) MemoryFunction {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
}
func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println("The", input, "th term for the Fibonnaci series is:", v(input).(int))

}
