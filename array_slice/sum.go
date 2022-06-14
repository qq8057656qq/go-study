package array_slice

func Sum(numbers []int) int {
	sum := 0
	//rang会迭代数组，每次迭代都会返回数组元素的索引和值，使用_空白标志符来忽略索引
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			//我们可以使用语法 slice[low:high] 获取部分切片。
			//如果在冒号的一侧没有数字就会一直取到最边缘的元素。
			//在我们的函数中，我们使用 numbers[1:] 取到从索引 1 到最后一个元素。
			//你可能需要花费一些时间才能熟悉切片的操作。
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
