class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) < 2:
        	return 0
        res = 0
        i = pos = new_pos = 0
        while True:
        	new_pos = pos
        	while i <= pos:
        		new_pos = max(new_pos, i + nums[i])
        		i += 1
        	res += 1
        	pos = new_pos
        	if pos >= len(nums) - 1:
        		break
        return res

s = Solution()
print s.jump([2,3,1,1,4])       
