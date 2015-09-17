package main
import "fmt"
import "testing"
func TestFibonacci(t *testing.T) { 
var i int64
var v int64
var j int64
var k int64
k = Fib(int64(-1)) 
if k != -1 { 
t.Error("TEST CASE 0 :Expected -1, got ", k) 
}else{
 fmt.Println("TEST CASE 0 ((Fibonacci(-1) $$No fibonacci for negative numbers) Passed")
}

v = Fib(int64(0)) 
if v != 0 { 
t.Error("TEST CASE 1 :Expected 0, got ", v) 
}else{
 fmt.Println("TEST CASE 1 ((Fibonacci(0) is correct) Passed")
}
i = Fib(int64(10)) 
if i != 55 { 
t.Error("TEST CASE 2 :Expected 55, got ", i) 
}else{
fmt.Println("TEST CASE 2 (Fibonacci(10) is correct Passed!!")
}
j = Fib(int64(40))
if j != 102334155 { 
t.Error("TEST CASE 3 :Expected 102334155, got ", j) 
}else{
fmt.Println("TEST CASE 3 (Fibonacci(40) is correct Passed!!")
}

}

