package 差分数组

import (
	"fmt"
)

/**
所谓差分数组即 diff[i] = nums[i] - nums[i-1]
当一段区间的所有元素都+val时，例如nums[i]~nums[j]之间的元素都加上了val,
diff[i]-diff[j]保持不变,
diff[l] += val, diff[r+1] -= val, 可以表示 [l, r]区间的元素都加上了val
重建时： cnt[0] =  diff[0],
for i in range(1, n):
    cnt[i] = cnt[i-1] + diff[i]
 */

func ExampleDifferenceArray() {
	var N = 11
	gaps := [][]int{{1,3},{2,5},{9,10}} // 表示[i, j]之间的区间+1

	diff := make([]int, N+1)
	for _, g := range gaps {
		l, r := g[0], g[1]
		diff[l]++
		diff[r+1]--
	}

	nums := make([]int, N)
	for i := 1; i < N; i++ {
		nums[i] = nums[i-1] + diff[i]
	}
	fmt.Println(nums)
	// Output:
	// [0 1 2 2 1 1 0 0 0 1 1]
}


