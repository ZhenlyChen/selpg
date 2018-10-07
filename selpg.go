package main

import (
	"bufio"
	"io"
	"os/exec"
)

func selpg(src *bufio.Reader, dst io.Writer, args SelpgArgs) {
	if args.find {
		count := 0
		for count < args.start {
			_, err := src.ReadBytes('\f')
			check(err)
			count++
		}
		for count < args.end {
			input, err := src.ReadBytes('\f')
			check(err)
			output(args.dest, dst, input)
			count++
		}
	} else {
		for i := 0; i < args.start; i++ {
			_, err := src.ReadBytes('\n')
			check(err)
		}
		for i := 0; i < args.end-args.start; i++ {
			input, err := src.ReadBytes('\n')
			check(err)
			output(args.dest, dst, input)
		}
	}
}

func output(printer string, dst io.Writer, data []byte) {
	if printer == "" {
		dst.Write(data)
	} else {
		// cmd := exec.Command("lp", "-d" + printer)
		cmd := exec.Command("cat")
		in, err := cmd.StdinPipe()
		check(err)

		go func() {
			defer in.Close()
			io.WriteString(in, string(data))
		}()
		_, err = cmd.CombinedOutput()
		check(err)
		// fmt.Printf("%s\n", out)
	}
}