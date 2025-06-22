package main

import (
	"flag"
)

func maint() {
	add := flag.Int("add", 0, 0, "add two integers")
	sub := flag.Int("sub", 0, 0, "subtract two integers")
	mul := flag.Int("mul", 0, 0, "multiply two integers")
	div := flag.Int("div", 0, 1, "divide two integers")

	flag.Parse()


}

func getInput(a int, b int)
