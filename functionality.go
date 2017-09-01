package main

func each(handler func(string), text string) bool {
	if finished {
		return false
	}
	if len(text) >= maxlen {
		handler(text)
	} else {
		handler(text)
		for i, c := range dict {
			if len(text) < len(startAt) && i < startAt[len(text)] {
				continue
			}
			if !each(handler, text+string(c)) {
				return false
			}
		}
	}
	return true
}
