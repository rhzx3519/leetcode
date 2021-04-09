class Solution:
    def countRangeSum(self, nums: List[int], lower: int, upper: int) -> int:
        res, pre, now = 0, [0], 0
        for n in nums:
            now += n
            res += bisect.bisect_right(pre, now - lower) - bisect.bisect_left(pre, now - upper)
            bisect.insort(pre, now)
        return res


# 解题思路
# 使用前缀数组prepre，然后每个前缀和pre[i]pre[i]二分查找前面i-1i−1个和的pre[i]-lowerpre[i]−lower和pre[i]-upperpre[i]−upper的位置得出区间和的数量，然后把pre[i]pre[i]二分插入到数组中保持数组有序

# 作者：fan-cai
# 链接：https://leetcode-cn.com/problems/count-of-range-sum/solution/python3-6xing-dai-ma-jian-ji-qian-zhui-he-er-fen-c/
# 来源：力扣（LeetCode）
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。