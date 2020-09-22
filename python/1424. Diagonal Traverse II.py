class Solution(object):
    def findDiagonalOrder(self, nums):
        """
        :type nums: List[List[int]]
        :rtype: List[int]
        """
        ans = []
        for i in range(len(nums)):
            for j in range(len(nums[i])):
                ans.append(((i+j), j, (nums[i][j])))
        ans.sort()
        return [t[2] for t in ans]

# 思路：位于同一对角线上元素的索引下标i+j的值相同，下一条对角线的j值与上一条相同或者比上一条大1
# 按照 i+j, j这两项排序


if __name__ == '__main__':
    # nums = [[1,2,3],[4,5,6],[7,8,9]]
    # nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
    # nums = [[6],[8],[6,1,6,16]]
    nums = [[11,6,9,20],[16,1,20],[14,19,14,17,15],[8,19,11,3],[3,13,17,4]]
    su = Solution()
    su.findDiagonalOrder(nums)        