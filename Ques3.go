package main
import ("fmt";"time")
func main() {

    c1 := make(chan string, 1)

    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "Case1: Runs: Sleep Time is 1 second and timeout set to 4 seconds"
    }()

      select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 4):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "Case1: Runs: Sleep Time is 2 second and timeout set to 3 seconds"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }

    c3 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 5)
        c3 <- "Case3: Runs: Sleep Time is 5 second and timeout set to 3 seconds"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("TimeOut Occurs for Case3: Sleep Time(5 seconds) more than Timeout(3 seconds)")
    }
}
