package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	// Println writes to the std logger
	log.Println("message")

	// Fatalln is Println followed by a cal to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println followed by a call to panic()
	log.Panicln("panic message")
}

// 1 << 0 = 000000001 =1
// 1 << 1 = 000000010 = 2
// 1 << 2 = 000000100 = 4
// 1 << 3 = 000001000 = 8
// 1 << 4 = 000010000 = 16
