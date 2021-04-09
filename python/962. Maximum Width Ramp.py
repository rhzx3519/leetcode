class Solution(object):
    def maxWidthRamp(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        n = len(A)
        t = sorted(range(n), key=lambda i: A[i])
        min_val = n
        max_val = 0
        for i in t:
            max_val = max(max_val, i - min_val)
            min_val = min(min_val, i)
        return max_val

    def maxWidthRamp(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        n = len(A)
        st = []
        for i in range(n):
            if not st or A[st[-1]] > A[i]:
                st.append(i)
        print st

        ans = 0
        for i in range(n-1, -1, -1):
            while st and A[st[-1]] <= A[i]:
                ans = max(ans, i - st[-1])
                st.pop()
        return ans

# 思路1：构造单调递减栈，从后往前遍历，比较当前元素和栈顶指向的元素，
# 如果大于等于，则出栈，知道栈顶指向元素小于当前元素，在当中得到最大宽度

# 思路2：数组下标根据值排序，然后顺序求出下标数组的最大差值
            

        