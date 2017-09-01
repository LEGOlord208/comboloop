package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jD91mZM2/stdutil"
)

const DictAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DictNumbers = "0123456789"

var dict string
var finished bool
var maxlen int
var startAt []int

func main() {
	var delay int
	var start string

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
		if len(args) > 1 {
			printHelp()
			return
		}
	case "numbers":
		dict = DictNumbers
		if len(args) > 1 {
			printHelp()
			return
		}
	case "custom":
		if len(args) < 2 {
			stdutil.PrintErr("No custom dictionary provided", nil)
			return
		} else if len(args) > 2 {
			printHelp()
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
	if len(start) > maxlen {
		stdutil.PrintErr("Start can't be longer than max length!", nil)
		return
	}
	if len(start) > 0 {
		startAt = make([]int, len(start))

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

			startAt[i] = index
		}
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		<-c
		finished = true
	}()

	each(func(result string) {
		if len(result) > 0 {
			fmt.Println(result)
		}

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}, "")
}

func printHelp() {
	fmt.Println("Usage: comboloop [flags...] <built-in dictionary> OR comboloop [flags...] custom <dictionary>")
	fmt.Println()
	fmt.Println("Built-in dictionaries:")
	fmt.Println("- alphabet")
	fmt.Println("- numbers")
}
