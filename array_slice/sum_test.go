package array_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("累加任意大小的数组", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given,%v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want :=
		//"bob"
		[]int{3, 9}
	//reflect.DeepEqual 不是类型安全的，用不同类型去做对比也不会报错，但是结果不会是我们想要的
	checkSums(t, got, want)
}

func checkSums(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
