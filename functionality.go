package main

func increment(pos int) {
	if len(result) <= 0 {
		result = []int{0}
		return
	}
	if result[pos]+1 >= len(dict) {
		result[pos] = 0

		if pos <= 0 {
			if len(result) >= maxlen {
				finished = true
				return
			}
			result = append([]int{0}, result...)
		} else {
			increment(pos - 1)
		}
		return
	}
	result[pos]++
}

func getresult() string {
	res := ""
	for _, i := range result {
		res += string(dict[i])
	}
	return res
}
