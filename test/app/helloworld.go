package main
import ( "fmt"
         "time"
       )   

func main() {
	for true {
		fmt.Println("Hello auto-home!")
		duration := time.Second
		time.Sleep(duration)
	}
}