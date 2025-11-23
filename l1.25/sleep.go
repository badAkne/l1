package main

import "time"

func sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	d := time.Second * 5
	sleep(d)
}
