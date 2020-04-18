class Solution(object):
    def maxSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        res = []
        que = []
        for i in range(len(nums)):
            while que and nums[que[-1]] < nums[i]:
                que.pop()
            que.append(i)
            if i-que[0]==k:
                que.pop(0)
            if i >=k-1:
                res.append(nums[que[0]])
        return res

# 队列que的队头存储最大值        