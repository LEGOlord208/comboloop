package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/legolord208/stdutil"
)

const DictAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DictNumbers = "0123456789"

var dict string
var result []int
var finished bool
var maxlen int

func main() {
	var start string
	var delay int

	flag.IntVar(&maxlen, "len", 0, "Specifies the max length")
	flag.IntVar(&delay, "delay", 0, "Specifies the delay between turns")
	flag.StringVar(&start, "start", "", "Specify at what string to start at")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printHelp()
		return
	}

	switch args[0] {
	case "alphabet":
		dict = DictAlphabet
	case "numbers":
		dict = DictNumbers
	case "custom":
		if len(args) < 2 {
			stdutil.PrintErr("No custom dictionary provided", nil)
			return
		}
		dict = args[1]

		if len(dict) <= 0 {
			stdutil.PrintErr("Dictionary is empty", nil)
			return
		}
	default:
		printHelp()
		return
	}

	if maxlen <= 0 {
		maxlen = len(dict)
	}
	if len(start) > 0 {
		result = make([]int, len(start))
		for i, c := range start {
			index := -1
			for i2, c2 := range dict {
				if c == c2 {
					index = i2
				}
			}

			if index < 0 {
				stdutil.PrintErr("Start includes items outside dictionary", nil)
				return
			}

			result[i] = index
		}
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		<-c
		finished = true
	}()

	for !finished {
		if len(result) > 0 {
			fmt.Println(getresult())
		}
		increment(len(result) - 1)

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}

func printHelp() {
	fmt.Println("Usage: comboloop [flags...] <built-in dictionary> OR comboloop [flags...] custom <dictionary>")
	fmt.Println()
	fmt.Println("Built-in dictionaries:")
	fmt.Println("- alphabet")
	fmt.Println("- numbers")
}
