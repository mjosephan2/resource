package watercontainer

// https://leetcode.com/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-interview-150
func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left <= right {
		hRight := height[right]
		hLeft := height[left]
		if area := min(hRight, hLeft) * (right - left); area > maxArea {
			maxArea = area
		}
		if hRight < hLeft {
			right--
		} else {
			left++
		}
	}
	return maxArea
}

func maxAreaWithContainerBlocked(height []int) int {
	maxArea := 0
	left, right := 0, 0
	for right < len(height) && left <= right {
		hRight := height[right]
		hLeft := height[left]
		if area := min(hRight, hLeft) * (right - left); area > maxArea {
			maxArea = area
		}
		// this means that we should set left = right
		// if not the left block would be blocked
		if height[right] > height[left] {
			left = right
		}
		right++
	}

	return maxArea
}
