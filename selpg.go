package main

import (
	"bufio"
	"fmt"
	"io"
)

func selpg(src *bufio.Reader, dst  io.Writer, begin, end int, page bool) {
	fmt.Println(begin, end ,page)
	if page {
		count := 0
		for count < begin {
			_, err := src.ReadString('\f')
			if err != nil {
				panic(err)
			}
			count ++
		}
		for count < end {
			input, err := src.ReadString('\f')
			if err != nil {
				panic(err)
			}
			dst.Write([]byte(input))
			count ++
		}
	} else {
		for i := 0; i < begin; i++ {
			if _, err := src.ReadString('\n'); err != nil {
				panic(err)
			}
		}
		for i := 0; i < end - begin; i++ {
			input, err := src.ReadString('\n')
			if err != nil {
				panic(err)
			}
			dst.Write([]byte(input))
		}
	}
}