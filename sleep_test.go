package main
import "fmt"
import "testing"
func TestSleep(t *testing.T) { 
var v int64 
v= Sleep(4)
if v!= 4 {
t.Error("Expected 4,got", v)
}else{
fmt.Println("Test Case for Sleep function is passed")

}
}