#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def maxChunksToSorted(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        l = r = 0
        res = 0
        for i in range(len(arr)):
            r = max(r, arr[i])
            if r <= i:
                res += 1
                r = i+1
        return res

if __name__ == '__main__':
    arr = [1,0,2,3,4]     
    su = Solution()
    print su.maxChunksToSorted(arr)


# 数组arr是[0, 1, ..., arr.length - 1]的一种排列，我们将这个数组分割成几个“块”，并将这些块分别进行排序。之后再连接起来，使得连接的结果和按升序排序后的原数组相同。

# 我们最多能将数组分成多少块？

# 示例 1:

# 输入: arr = [4,3,2,1,0]
# 输出: 1
# 解释:
# 将数组分成2块或者更多块，都无法得到所需的结果。
# 例如，分成 [4, 3], [2, 1, 0] 的结果是 [3, 4, 0, 1, 2]，这不是有序的数组。

# 示例 2:

# 输入: arr = [1,0,2,3,4]
# 输出: 4
# 解释:
# 我们可以把它分成两块，例如 [1, 0], [2, 3, 4]。
# 然而，分成 [1, 0], [2], [3], [4] 可以得到最多的块数。
# 注意:

# arr 的长度在 [1, 10] 之间。
# arr[i]是 [0, 1, ..., arr.length - 1]的一种排列。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/max-chunks-to-make-sorted
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。