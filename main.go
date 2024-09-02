package main

import (
	"log"
	"vblog/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("vblog server start failed: %v\n", err)
	}
}
