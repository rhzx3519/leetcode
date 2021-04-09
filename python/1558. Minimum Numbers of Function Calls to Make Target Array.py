class Solution(object):
    def minOperations(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        def move(num):
            # print num
            step0 = step1 = 0
            while num:
                if num&1:
                    num -= 1
                    step0 += 1
                else:
                    num >>= 1
                    step1 += 1
            # print (step0, step1)
            return (step0, step1)

        total_op0 = 0
        max_op1 = 0
        for num in nums:
            op0, op1 = move(num)
            max_op1 = max(max_op1, op1)
            total_op0 += op0
        return total_op0 + max_op1
        
        