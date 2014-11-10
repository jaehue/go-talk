package main

func main() {
	done := make(chan bool)
	doSort := func(s []int) {
		sort(s)
		done <- true
	}
	i := pivot(s)
	go doSort(s[:i])
	go doSort(s[i:])
	<-done
	<-done
}
