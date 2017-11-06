package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	head  = 0
	ptr   = 0
	plen  int
	bPair = make(map[int]int)
	buf   = make([]int, 100)
)

func main() {
	if len(os.Args) < 2 {
		panic("input brainf*ck file please.")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic("file not exist")
	}
	plen = len(data)

	var (
		bHead = 0
		bBuf  = make([]int, 100)
	)

	for i, r := range data {
		switch r {
		case '[':
			bBuf[bHead] = i
			bHead++
		case ']':
			bHead--
			lptr := bBuf[bHead]
			bPair[lptr] = i
			bPair[i] = lptr
		}
	}

	for head < plen {
		switch data[head] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			buf[ptr]++
		case '-':
			buf[ptr]--
		case '.':
			fmt.Print(string(buf[ptr]))
		case ',':
			var str string
			fmt.Scan(&str)
			buf[ptr] = int(str[0])
		case '[':
			if buf[ptr] == 0 {
				head = bPair[head]
			}
		case ']':
			if buf[ptr] != 0 {
				head = bPair[head]
			}
		}
		head++
	}
}
