package ascendingindex

import "fmt"

func ascendingIndex(inputList []int) []int {
	// initialize empty output list
	outputList := make([]int, len(inputList))
	// initialize stack
	stack := [][2]int{}
	for i, v := range inputList {
		if len(stack) == 0 {
			stack = append(stack, [2]int{i, v})
			continue
		}
		for len(stack) > 0 {
			prev := peek(stack)
			fmt.Println(prev, i, v)
			// check how prev value compare to current value
			prevIndex := prev[0]
			prevValue := prev[1]
			if prevValue >= v {
				// current value is less than previous
				break
			}
			stack = stack[:len(stack)-1]
			outputList[prevIndex] = i - prevIndex
		}
		stack = append(stack, [2]int{i, v})
	}
	return outputList
}

func peek(stack [][2]int) [2]int {
	if len(stack) == 0 {
		return [2]int{}
	}
	return stack[len(stack)-1]
}
