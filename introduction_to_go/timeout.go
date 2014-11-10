package main

func worker(start chan bool) {
	for {
		timeout := time.After(30 * time.Second)
		select {
		// ... do some stuff
		case <-timeout:
			return
		}

	}
}

func worker(start chan bool) {
	timeout := time.After(30 * time.Second)
	for {

		select {
		// ... do some stuff
		case <-timeout:
			return
		}
	}
}
