package main

func openLock(deadends []string, target string) int {
	var visited = make(map[string]struct{})

	for _, v := range deadends {
		visited[v] = struct{}{}
	}

	var step int
	var queue = make([]string, 0)
	queue = append(queue, "0000")

	for len(queue) > 0 {
		elemCount := len(queue)
		for i := 0; i < elemCount; i++ {
			elem := queue[i]
			_, ok := visited[elem]
			if ok {
				continue
			}

			visited[elem] = struct{}{}
			if elem == target {
				return step
			}

			for j := 0; j < 4; j++ {
				addedElem := addOne(elem, j)
				subbedElem := subOne(elem, j)
				queue = append(queue, addedElem, subbedElem)
			}
		}

		queue = queue[elemCount:]
		step++
	}

	return -1
}

func addOne(elem string, pos int) string {
	elemByteSlice := []byte(elem)
	if elemByteSlice[pos] == 57 {
		elemByteSlice[pos] = 48
	} else {
		elemByteSlice[pos]++
	}

	return string(elemByteSlice)
}

func subOne(elem string, pos int) string {
	elemByteSlice := []byte(elem)
	if elemByteSlice[pos] == 48 {
		elemByteSlice[pos] = 57
	} else {
		elemByteSlice[pos]--
	}

	return string(elemByteSlice)
}
