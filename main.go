package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

func main() {
	var start, end, leap int
	var f bool

	flag.IntVarP(&start, "start", "s", 0, "start page")
	flag.IntVarP(&end, "end", "e", 1, "end page")
	flag.IntVarP(&leap, "leap", "l", 72, "the count of row in each page")
	flag.BoolVarP(&f, "find", "f", false, "find page")
	flag.Parse()
	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
	if end < start || start <= 0 {
		fmt.Println("invalid start and end")
		return
	}
	if leap <= 0 {
		fmt.Println("invalid the count of row in each page")
		return
	}
	if leap != 72 && f {
		fmt.Println("you can't use two mode in the same time")
		return
	}
	fmt.Println(start, end, leap, f)
	start--
	if !f {
		start *= leap
		end *= leap
	}

	// get fileName
	fileName := ""
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i][0] == '-' {
			if len(os.Args[i]) == 2 && os.Args[i][1] != 'f' {
				i++
			}
		} else {
			fileName = os.Args[i]
			break
		}
	}
	if fileName != "" {
		fmt.Println("file: ", fileName)
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		inputReader := bufio.NewReader(file)
		selpg(inputReader, os.Stdout, start, end, f)
	} else {
		inputReader := bufio.NewReader(os.Stdin)
		selpg(inputReader, os.Stdout, start, end, f)
	}
}
