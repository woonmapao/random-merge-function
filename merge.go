package merge

// Merge returns a new ascending slice containing all values from the three inputs.
// collection1 must be sorted descending. collection2 and collection3 must be sorted ascending.
func Merge(collection1, collection2, collection3 []int) []int {
	totalLen := len(collection1) + len(collection2) + len(collection3)
	result := make([]int, 0, totalLen)

	i := len(collection1) - 1
	j := 0
	k := 0

	for len(result) < totalLen {
		source := 0
		next, ok := nextValue(collection1, collection2, collection3, i, j, k)
		if !ok {
			break
		}

		if i >= 0 && collection1[i] == next {
			source = 1
		} else if j < len(collection2) && collection2[j] == next {
			source = 2
		} else {
			source = 3
		}

		result = append(result, next)

		switch source {
		case 1:
			i--
		case 2:
			j++
		case 3:
			k++
		}
	}

	return result
}

func nextValue(collection1, collection2, collection3 []int, i, j, k int) (int, bool) {
	var next int
	hasNext := false

	if i >= 0 {
		next = collection1[i]
		hasNext = true
	}
	if j < len(collection2) && (!hasNext || collection2[j] < next) {
		next = collection2[j]
		hasNext = true
	}
	if k < len(collection3) && (!hasNext || collection3[k] < next) {
		next = collection3[k]
		hasNext = true
	}

	return next, hasNext
}
