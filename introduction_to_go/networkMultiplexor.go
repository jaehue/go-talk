package main

func worker(messages chan string) {
	for {
		var msg string // ... generate a message
		messages <- msg
	}
}
func main() {
	messages := make(chan string)
	conn, _ := net.Dial("tcp", "example.com")

	for i := 0; i < 100; i++ {
		go worker(messages)
	}
	for {
		msg := <-messages
		conn.Write([]byte(msg))
	}
}
