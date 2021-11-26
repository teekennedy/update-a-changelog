package main

import (
	"flag"
	"fmt"
	"time"
)

var who = flag.String("who", "World", "Say hello to who")

func main() {
	flag.Parse()
	fmt.Println("Hello,", *who)
	fmt.Printf("::set-output name=time::%v\n", time.Now().Format(time.RFC822))
}
