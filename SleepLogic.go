package main
import ( "time" )
import "fmt"

func Sleep(x int)int64 {
startTime := time.Now()
    <-time.After(time.Second * time.Duration(x))
endTime := time.Now()
 return int64(endTime.Sub(startTime).Seconds())
}

func Sleep_caller() {
fmt.Println("**Welcome to implementation of sleep function using Time.After function")
fmt.Println("Press enter to break the output in between")
c1 := make(chan string)
c2 := make(chan string)
go func() {
		for{
		
		c1 <- "from1"
		//calling sleep function
		Sleep(6)
		}
		}()
		
	    go func() {
		for{
		c2 <- "from2"
		//calling sleep function
		Sleep(10)
		}
		
		}()
		go func() {
		for {
		select{
		//select will print the message it receives first
		case msg1 := <-c1:
		fmt.Println(msg1)		
		case msg2 := <-c2:
		fmt.Println(msg2)
		case  <- time.After(time.Second ):
		fmt.Println("timeout");
		}
		}
		}()
				
		var input string
	fmt.Scanln(&input)
	}
