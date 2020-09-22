class Solution(object):
    def simplifiedFractions(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        def gcd(a, b):
            return gcd(b, a%b) if b else a
        
        ans = []
        for i in range(1, n+1):
            for j in range(1, i):
                if gcd(j, i)==1:
                    ans.append('{}/{}'.format(j,  i))
        return ans    
                