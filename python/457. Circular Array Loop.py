class Solution(object):
    def circularArrayLoop(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        n = len(nums)
        if n <2: return False

        def getNextPos(i):
            return (((i + nums[i])%n) + n)%n # i + nums[i])%n 可能为负数

        for i in range(n):
            slow = i
            fast = getNextPos(i)
            if nums[i]==0:
                continue
            while True:
                if nums[slow]*nums[fast] < 0 or nums[fast]*nums[getNextPos(fast)] < 0:
                    break
                if slow==fast:
                    if slow==getNextPos(slow):
                        break
                    else:
                        return True
                slow = getNextPos(slow)
                fast = getNextPos(getNextPos(fast))
        return False
                    
            
