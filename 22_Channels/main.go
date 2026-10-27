//1. Basic Syntax & Operations
// 1. Creation
ch := make(chan int) // Creates a channel that transmits integers

// 2. Sending Data
ch <- 42 // Send the value 42 into the channel

// 3. Receiving Data
value := <-ch // Receive a value from the channel and assign it to a variable
<-ch          // Receive and discard the value

// 4. Closing a Channel
close(ch) // Signals that no more values will be sent


//3. Channel Directionality

func producer(ch chan<- string) {
    ch <- "data" // OK
    // val := <-ch // Compile Error!
}

func consumer(ch <-chan string) {
    val := <-ch // OK
    // ch <- "data" // Compile Error!
}
