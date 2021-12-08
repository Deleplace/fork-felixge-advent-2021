//go:build v2

package main

func Answer(inputText string) (int, error) {
	school := make([]int, 9)

	// Assume the input has exactly the form x,x,x,x,....,x
	// with x in [0..8]
	for i, n := 0, len(inputText); i < n; i += 2 {
		c := inputText[i]
		age := int(c - '0')
		school[age]++
	}

	// Sliding school
	p := 0
	for i := 0; i < 256; i++ {
		if p <= 1 {
			school[p+7] += school[p]
		} else {
			school[p-2] += school[p]
		}
		p++
		if p == 9 {
			p = 0
		}
	}

	var total int
	for _, count := range school {
		total += count
	}
	return total, nil
}
