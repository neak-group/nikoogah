package main

import "github.com/neak-group/nikoogah/internal/boot"

func main() {
	run, err := boot.Boot()
	if err != nil {
		panic(err)
	}

	run.Run()
}
