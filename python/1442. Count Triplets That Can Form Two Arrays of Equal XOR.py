class Solution(object):
    def countTriplets(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        n = len(arr)
        if n < 2: return 0

        ans = 0
        for i in range(n):
            s = arr[i]
            for k in range(i+1, n):
                s ^= arr[k]
                if s==0:
                    ans += k-i
        return ans

#思路：i,j,k, 可以推出arr[i]^arr[i+1]^..arr[j]^...arr[k] = 0