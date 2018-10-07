package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

type SelpgArgs struct {
	start int
	end int
	leap int
	find bool
	dest string
	fileName string
}

func main() {
	var args SelpgArgs
	parseArgs(&args)
	if args.fileName != "" {
		file, err := os.Open(args.fileName)
		check(err)
		inputReader := bufio.NewReader(file)
		selpg(inputReader, os.Stdout, args)
	} else {
		inputReader := bufio.NewReader(os.Stdin)
		selpg(inputReader, os.Stdout, args)
	}
}

func parseArgs(args *SelpgArgs) {
	flag.IntVarP(&args.start, "start", "s", -1, "Start page")
	flag.IntVarP(&args.end, "end", "e", -1, "End page")
	flag.IntVarP(&args.leap, "leap", "l", 72, "The count of row in each page")
	flag.BoolVarP(&args.find, "find", "f", false, "Find page breaks")
	flag.StringVarP(&args.dest, "dest", "d", "", "Destination")
	flag.Parse()
	if args.end < args.start || args.start <= 0 {
		fmt.Println("invalid start and end")
		return
	}
	if args.leap <= 0 {
		fmt.Println("invalid the count of row in each page")
		return
	}
	if args.leap != 72 && args.find {
		fmt.Println("you can't use two mode in the same time")
		return
	}
	args.start--
	if !args.find {
		args.start *= args.leap
		args.end *= args.leap
	}
	// get fileName
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i][0] == '-' {
			if len(os.Args[i]) == 2 && os.Args[i][1] != 'f' {
				i++
			}
		} else {
			args.fileName = os.Args[i]
			break
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
