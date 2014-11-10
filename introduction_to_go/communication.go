package main

type Task struct {
	// some state
}
type Pool struct {
	Mu    sync.Mutex
	Tasks []Task
}

func worker(pool *Pool) { // many of these run concurrently
	for {
		pool.Mu.Lock()
		task := pool.Tasks[0]
		pool.Tasks = pool.Tasks[1:]
		pool.mu.Unlock()
		process(task)
	}
}
