from Queue import PriorityQueue
class Comparable(object):
    def __init__(self, fre, val):
        self.fre = fre
        self.val = val
    
    def __cmp__(self, other):
        if self.fre == other.fre:
            return 0
        elif self.fre > other.fre:
            return 1
        else:
            return -1

class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        mp = {}
        for num in nums:
            fre = mp.setdefault(num, 0)
            mp[num] = fre+1
        
        res = []
        que = PriorityQueue()
        for num, fre in mp.iteritems():
            que.put(Comparable(fre, num))
            if que.qsize() > k:
                que.get()

        while que.qsize():
            res.append(que.get().val)

        res.reverse()
        return res
        

if __name__ == '__main__':
    nums = [1,1,1,2,2,3]
    k = 2
    su = Solution()
    print su.topKFrequent(nums, k)



