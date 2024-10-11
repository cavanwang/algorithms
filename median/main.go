package main

import "fmt"

func main() {
	got1 := GetMedia([]int{1}, []int{2})
	fmt.Println("....................................")
	got2 := GetMedia([]int{}, []int{})
	fmt.Println(got1, got2)
}

// All elements in a and b are all larger than 0. Return -1 if not found.
func GetMedia(a, b []int) float64 {
	if len(a) == 0 && len(b) == 0 {
		return -1
	}

	shiftCount := (len(a) + len(b) + 1) / 2
	anotherShiftCount := shiftCount
	if (len(a) + len(b)) % 2 == 0 {
		anotherShiftCount += 1
	}
	println("shiftCount, anothercount: ", shiftCount, anotherShiftCount)

	if len(a) == 0 {
		return float64(b[shiftCount-1]+b[anotherShiftCount-1]) / 2.0
	} else if len(b) == 0 {
		return float64(a[shiftCount-1]+a[anotherShiftCount-1]) / 2.0
	}

	n := 0
	shiftValue := 0
	leftA, leftB := a, b
	for {
		leftA, leftB, shiftValue = shiftNextMin(leftA, leftB)
		fmt.Println("leftA= leftB= shift=", leftA, leftB, shiftValue)
		n++

		// 如果只剩下一个数组了
		var left []int
		if len(leftA) == 0 {
			left = leftB
		} else if len(leftB) == 0 {
			left = leftA
		}

		// 如果已经移除了shiftCount，可以计算最终结果了
		if n == shiftCount {
			if shiftCount == anotherShiftCount { // 只需要一个中间数
				println("shiftvalue: ", shiftValue)
				return float64(shiftValue)
			}
			// 需要2个中间数的平均值
			_, _, secondValue := shiftNextMin(leftA, leftB)
			println("two value: ", shiftValue, secondValue)
			return float64(shiftValue+secondValue) / 2
		}

		// 如果尚未移除shiftCount个，但是已经有一个数组为空了，也可以直接计算最终结果了
		if left != nil {
			if shiftCount == anotherShiftCount {
				println("one value: ", left[shiftCount-n])
				return float64(left[shiftCount-n-1])
			}
			println("two values: ", left[shiftCount-n], left[shiftCount-n+1])
			return float64(left[shiftCount-n-1]+left[shiftCount-n]) / 2
		}

		// 2个数组不为空，且移除数量不够，则继续移除
	}
}

func shiftNextMin(a, b []int) (leftA, leftB []int, shiftValue int) {
	if len(a) == 0 {
		return a, b[1:], b[0]
	}
	if len(b) == 0 {
		return a[1:], b, a[0]
	}

	if a[0] <= b[0] {
		return a[1:], b, a[0]
	}
	return a, b[1:], b[0]
}
