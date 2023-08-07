package main

func main() {
	s := "a{b{d,e{g,h{,i}}},c{f}}"
	idxs := []int{}
	stack := ""
	for i := range s {
		c := s[i]
		if c == '}' {
			idx := idxs[len(idxs)-1]
			idxs = idxs[:len(idxs)-1]
			root := stack[idx-1]
			stack[idx+1:]
		}
		if c == '{' {
			idxs = append(idxs, len(stack))
		}
		stack += string(c)
	}
}
