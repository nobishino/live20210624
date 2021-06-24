package a

func f(ch chan int) { // want "channel argument should be directed"
}

func g(ch <-chan int) {}

func h(ch chan<- int) {}

// 10 min

func i() chan int { // want "channel result should be directed"
	return nil
}

func j() <-chan int {
	return nil
}

func k() chan<- int {
	return nil
}

var x = func(ch chan string) {} // want "channel argument should be directed"
