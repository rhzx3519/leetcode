class Solution(object):
    def getMaximumConsecutive(self, coins):
        """
        :type coins: List[int]
        :rtype: int
        """
        n = len(coins)
        dp = [0]*(1<<n)
        for i in range(1<<n):
            for j in range(n):
                if i&(1<<j):
                    dp[i] += coins[j]

        dp.sort()

        print dp
        count = 0
        last = -1
        for num in dp:
            if num == last+1:
                count += 1
                last = num
            elif num == last:
                pass
            else:
                break
        return count

if __name__ == '__main__':
    su = Solution()
    su.getMaximumConsecutive([1,4,10,3,1])