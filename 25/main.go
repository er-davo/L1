package main

import "time"

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	sleep(time.Second * 5)
}
