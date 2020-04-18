class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix: return False
        row = len(matrix)-1
        col = 0
        while row >=0 and col<len(matrix[0]):
            if matrix[row][col] > target:
                row -= 1
            elif matrix[row][col] < target:
                col += 1
            else:
                return True
        return False

# 思路: 选取左下角或者右上角作为起始点,
# 时间复杂度:O(m+n)
