package main
import (
	"fmt"
	"time"
	
)
func worker(id int, jobs <-chan SampleData) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
       
    }
}
const keyServerAddr = "serverAddr"

