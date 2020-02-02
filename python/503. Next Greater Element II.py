class Solution(object):
    def nextGreaterElements(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        st = []
        n = len(nums)
        res = [-1] * n
        for i in range(2*n):
            idx = i%n
            while st and nums[st[-1]] < nums[idx]:
                # print st, nums[st[-1]]
                res[st.pop()] = nums[idx]
            if i < n:
                st.append(i)
        return res

if __name__ == '__main__':
    nums = [2,4,1,3]
    su = Solution()
    print su.nextGreaterElements(nums)        