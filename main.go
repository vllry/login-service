package main

func main() {
	initializeTracing()

	s := newServer()
	s.start()
}
