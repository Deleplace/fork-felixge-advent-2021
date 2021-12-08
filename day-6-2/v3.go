//go:build v3

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

	const N = 256
	m, k := N/9, N%9
	// Chunks of 9 units of time
	for i := 0; i < m; i++ {
		school[7] += school[0]
		school[8] += school[1]
		school[0] += school[2]
		school[1] += school[3]
		school[2] += school[4]
		school[3] += school[5]
		school[4] += school[6]
		school[5] += school[7]
		school[6] += school[8]
	}
	// Just a few remaining units of time
	p := 0
	for i := 0; i < k; i++ {
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
