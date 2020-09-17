package main

import (
	"flag"
	"log"

	"github.com/yuriizinets/googletransx/server"
)

func main() {
	// Parse arguments
	addr := flag.String("addr", ":25021", "Address to listen (default :25021)")
	flag.Parse()

	instance := server.BuildServer(*addr)
	log.Println("Starting server at " + *addr)
	log.Fatal(instance.ListenAndServe())
}
