package main

func main() {
	c := make(chan *Request)
	c <- &Request{name: "Nikhil"}
	// millions of Request
	Serve(c)
}

var sem = make(chan int, 10)

func handle(r *Request) {
	sem <- 1
	process(r) // May take 1min time
	<-sem
}
func Serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req)
	}
}
