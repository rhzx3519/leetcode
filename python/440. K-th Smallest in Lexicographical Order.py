class Solution(object):
    def findKthNumber(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: int
        """
        cur = 1
        k -= 1 # 初始化为cur=1, k自减1
        while k:
            step, first, last = 0, cur, cur+1
            while first<=n:
                step += min(n+1, last) - first # 
                first *= 10
                last *= 10
            if step <= k: # 不在子树中
                cur += 1
                k -= step
            else:   # 在子树中
                cur *= 10
                k -= 1
        return cur

# [字节跳动最常考题之一]本题史上最完整具体的手摸手解答，时间效率超越100%用户
# https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/solution/ben-ti-shi-shang-zui-wan-zheng-ju-ti-de-shou-mo-sh/
