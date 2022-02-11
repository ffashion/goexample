package main

//chan is type modifier
//channel 在一个时刻只能有一个goroutine 访问通道并发送数据
func main() {
	string_chan := make(chan string)

	//go 关键字 开启一个新的go routine
	go func() {
		//use channel output data
		//this code will be blocked
		string_chan <- "aaaa"
	}()

	data := <-string_chan
	print(data, "\n")

}
