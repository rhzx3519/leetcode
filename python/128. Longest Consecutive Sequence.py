class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        max_len = 0
        mem = {}
        for i, num in enumerate(nums):
            if num not in mem:
                l = mem.get(num-1, 0)
                r = mem.get(num+1, 0)
                cur_len = l+r+1
                max_len = max(max_len, cur_len)
                mem[num-l] = cur_len
                mem[num+r] = cur_len
        return max_len

# 思路：mem记录元素num所在的连续序列的长度        

if __name__ == '__main__':
    nums = [3,1,2,100]
    su = Solution()
    su.longestConsecutive(nums)                        