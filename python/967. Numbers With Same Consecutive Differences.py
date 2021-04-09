class Solution(object):
    def numsSameConsecDiff(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[int]
        """
        if n == 1:
            return range(10)
        
        ans = []
        for num in self.numsSameConsecDiff(n-1, k):
            if num==0:
                continue
            last = num%10
            if last >= k:
                ans.append(num*10 + last - k)
            if k != 0 and last + k <= 9:
                ans.append(num*10 + last + k)
        return ans
            
