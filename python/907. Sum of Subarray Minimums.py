class Solution(object):
    def sumSubarrayMins(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        MOD = 10**9 + 7
        ans = cur_sum = 0
        n = len(A)
        st = []
        for i in range(n):
            while st and A[st[-1]] >= A[i]:
                top = st.pop()
                new_top = -1 if not st else st[-1]
                # new_top... top ... i, A[top]是这段区间的最小值
                cur_sum += (A[i] - A[top])*(top - new_top)
            cur_sum += A[i] # 子数组[A[i]]的最小值一定是A[i]
            st.append(i)
            ans += cur_sum
        return ans%MOD

# 思路：... A[i] ... ，i之前有m个数比A[i]大，i之后有n个数比A[i]大，则以A[i]为
#       最小值的连续数组有(m-1)*(n-1)个 
#