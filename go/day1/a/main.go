package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/freeformz/adventOfCode2022/go/internal/input"
)

func main() {
	in, err := input.Get(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()

	var ce, me int
	is := bufio.NewScanner(in)
	for is.Scan() {
		row := is.Text()
		if len(row) == 0 {
			if ce > me {
				me = ce
			}
			ce = 0
			continue
		}
		c, err := strconv.Atoi(row)
		if err != nil {
			fmt.Println(err)
			return
		}
		ce += c
	}
	if ce > me {
		me = ce
	}

	fmt.Println(me)
}
