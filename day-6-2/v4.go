//go:build v4

package main

import "math/bits"
import "fmt"
import "log"

func Answer(inputText string) (int, error) {
	fmt.Println("Answer")
	var school School

	// Assume the input has exactly the form x,x,x,x,....,x
	// with x in [0..8]
	for i, n := 0, len(inputText); i < n; i += 2 {
		c := inputText[i]
		age := int(c - '0')
		school[age]++
	}

	const N = 256

	mat := ClockMatrix{
		{0, 1},
		{0, 0, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	/*
		for t := 0; t < N; t++ {
			if t < 18 {
				log.Println(school)
			}
			school = apply(mat, school)
		}
	*/
	//matN := power(mat, N)
	powerInPlace(&mat, N)
	school = apply(mat, school)

	var total int
	for _, count := range school {
		total += count
	}
	return total, nil
}

type School [9]int

type ClockMatrix [9][9]int

func mult(a, b ClockMatrix) (c ClockMatrix) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func multInto(a, b, c *ClockMatrix) {
	fmt.Println("multInto")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
}

func power(a ClockMatrix, n int) (b ClockMatrix) {
	if bits.OnesCount(uint(n)) != 1 {
		log.Fatal("power not implemented for n = ", n)
	}
	for i := 1; i < n; i *= 2 {
		a = mult(a, a)
	}
	return a
}

func powerInPlace(a *ClockMatrix, n int) {
	if bits.OnesCount(uint(n)) != 1 {
		log.Fatal("power not implemented for n = ", n)
	}
	for i := 1; i < n; i *= 2 {
		var b ClockMatrix
		multInto(a, a, &b)
		*a = b
	}
}

func apply(m ClockMatrix, s School) (t School) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			t[i] += m[i][j] * s[j]
		}
	}
	return t
}
