package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func run(args []string, stdout io.Writer) error {

	ints, err := strsToInts(args[1:])
	if err != nil {
		return err
	}

	drawStacks(ints)

	return nil
}

func drawStacks(ints []int) {
	m := max(ints)
	// fmt.Printf("max %d\n", m)

	x := 0

	for r := m; r > 0; r-- {
		x += drawRow(ints, r)
	}

	fmt.Printf("%s\n", strings.Repeat("-", len(ints)*2-1))
	fmt.Printf("%d\n", x)
}

func drawRow(ints []int, n int) int {

	catching := false
	tot := 0
	lev := 0

	for p := 0; p < len(ints); p++ {
		i := ints[p]

		switch {
		case i >= n && catching:
			tot += lev
			lev = 0
			catching = false
		case i < n && catching:
			lev++
		case i >= n && !catching:
			// nada
		case i < n && !catching:
			if p > 0 && ints[p-1] >= n {
				catching = true
				lev++
			}
		}

	}

	for _, i := range ints {
		if i >= n {
			fmt.Printf("# ")
		} else {
			fmt.Printf("  ")
		}
	}

	fmt.Printf("  %d\n", tot)

	return tot
}

func max(ints []int) int {
	m := 0
	for _, i := range ints {
		if i > m {
			m = i
		}
	}
	return m
}

func strsToInts(args []string) ([]int, error) {

	ints := []int{}

	for _, s := range args {
		v, err := strconv.Atoi(s)
		if err != nil {
			return ints, fmt.Errorf("strsToInts failed to convert %#v: %w", s, err)
		}
		ints = append(ints, v)
	}

	return ints, nil
}
