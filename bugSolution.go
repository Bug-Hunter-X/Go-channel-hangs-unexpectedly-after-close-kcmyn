func main() {
    fmt.Println("This is a sample Go program.")
    data := make([]int, 0, 100)
    for i := 0; i < 100; i++ {
        data = append(data, i)
    }
    fmt.Println(data)
    closeChan := make(chan bool)
    go func() {
        for i := 0; i < 100; i++ {
            closeChan <- true
        }
        close(closeChan)
    }()
    for range closeChan {
        fmt.Println("Reading from the channel...")
    }
    fmt.Println("Program finished")
} 
//The solution is to remove the for...range loop and replace it with a loop that checks if the channel is closed or not.
//This way, we make sure that the program exits after it has read all data from the channel.
func main() {
    fmt.Println("This is a sample Go program.")
    data := make([]int, 0, 100)
    for i := 0; i < 100; i++ {
        data = append(data, i)
    }
    fmt.Println(data)
    closeChan := make(chan bool)
    go func() {
        for i := 0; i < 100; i++ {
            closeChan <- true
        }
        close(closeChan)
    }()
    for {
        select {
        case <-closeChan:
            fmt.Println("Reading from the channel...")
        default:
            if len(closeChan) == 0 {
                break
            }
        }
        if len(closeChan) == 0 {
            break
        }
    }
    fmt.Println("Program finished")
}