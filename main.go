package main

import (
	"os"
	"users_generator/shell"

	"github.com/opentracing/opentracing-go/log"
)

func main() {
	err := shell.Init()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
