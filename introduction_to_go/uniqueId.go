package main

func main() {
	id := make(chan string)
	go func() {
		var counter int64 = 0
		for {
			id <- fmt.Sprintf("%x", counter)
			counter += 1
		}
	}()

	x := <-id // x will be 1
	x = <-id  //x will be 2
}
