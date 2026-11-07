package works

import (
	"encoding/binary"
	"math"
	"testing"
)

// -------------------------------------------------------------------------

func TestSumAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want int
	}{
		// Базовые случаи
		{name: "три числа", nums: []int{1, 2, 3}, want: 6},
		{name: "с отрицательными", nums: []int{10, -2, 4, 7}, want: 19},
		{name: "одно число", nums: []int{42}, want: 42},
		{name: "одно отрицательное", nums: []int{-7}, want: -7},

		// Граничные случаи
		{name: "пустой вызов", nums: nil, want: 0},
		{name: "пустой слайс", nums: []int{}, want: 0},
		{name: "все нули", nums: []int{0, 0, 0}, want: 0},
		{name: "ноль и числа", nums: []int{5, 0, -3}, want: 2},

		// Знаки
		{name: "все отрицательные", nums: []int{-1, -2, -3}, want: -6},
		{name: "сумма в ноль", nums: []int{10, -10}, want: 0},
		{name: "сумма в ноль (длинная)", nums: []int{1, -1, 2, -2, 3, -3}, want: 0},

		// Границы int
		{name: "max int", nums: []int{math.MaxInt}, want: math.MaxInt},
		{name: "min int", nums: []int{math.MinInt}, want: math.MinInt},
		{name: "max + min", nums: []int{math.MaxInt, math.MinInt}, want: -1},
		{name: "max + 0", nums: []int{math.MaxInt, 0}, want: math.MaxInt},
	}

	for _, tc := range tests {
		tc := tc // для Go < 1.22
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := sumAll(tc.nums...)
			if got != tc.want {
				t.Errorf("SumAll(%v) = %d, ожидалось %d", tc.nums, got, tc.want)
			}
		})
	}
}

// -------------------------------------------------------------------------

func TestSumAll_SliceExpand(t *testing.T) {
	t.Parallel()

	values := []int{5, 10, 15, 20}
	if got := sumAll(values...); got != 50 {
		t.Errorf("SumAll(values...) = %d, ожидалось 50", got)
	}
}

// -------------------------------------------------------------------------

func TestSumAll_ManyArgs(t *testing.T) {
	t.Parallel()

	const n = 1000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1 // 1..1000
	}
	// Сумма 1..n = n*(n+1)/2
	want := n * (n + 1) / 2
	if got := sumAll(nums...); got != want {
		t.Errorf("SumAll(1..%d) = %d, ожидалось %d", n, got, want)
	}
}

// -------------------------------------------------------------------------
func TestSumAll_Overflow(t *testing.T) {
	t.Parallel()

	got := sumAll(math.MaxInt, 1)
	if got != math.MinInt {
		t.Errorf("SumAll(MaxInt, 1) = %d, ожидалось MinInt (%d) — переполнение", got, math.MinInt)
	}
}

// -------------------------------------------------------------------------
func FuzzSumAll(f *testing.F) {

	f.Add([]byte{})
	f.Add(encodeInts(1, 2, 3))
	f.Add(encodeInts(0))
	f.Add(encodeInts(-1, -2, -3))
	f.Add(encodeInts(math.MaxInt, math.MinInt))

	f.Fuzz(func(t *testing.T, data []byte) {
		nums := decodeInts(data)
		got := sumAll(nums...)

		// Свойство: коммутативность
		rev := make([]int, len(nums))
		for i, v := range nums {
			rev[len(nums)-1-i] = v
		}
		if sumAll(rev...) != got {
			t.Errorf("не коммутативно: %v vs %v", nums, rev)
		}

		want := 0
		for _, n := range nums {
			want += n
		}
		if got != want {
			t.Errorf("SumAll(%v) = %d, ожидалось %d", nums, got, want)
		}
	})
}

// encodeInts упаковывает int в []byte (big-endian, 8 байт на число).
func encodeInts(nums ...int) []byte {
	buf := make([]byte, len(nums)*8)
	for i, n := range nums {
		binary.BigEndian.PutUint64(buf[i*8:], uint64(n))
	}
	return buf
}

// decodeInts разбирает []byte в []int; неполный «хвост» отбрасывается.
func decodeInts(data []byte) []int {
	n := len(data) / 8
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(int64(binary.BigEndian.Uint64(data[i*8:])))
	}
	return nums
}
