package main

func worker(start chan bool) {
	heartbeat := time.Tick(30 * time.Second)
	for {

		select {
		// ... do some stuff
		case <-heartbeat:
			// ... do heartbeat stuff
		}
	}
}
