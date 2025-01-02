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
}