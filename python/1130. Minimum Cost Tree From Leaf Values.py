class Solution(object):
    def mctFromLeafValues(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        ans = 0
        st = [float('inf')]
        for num in arr:
            while st and st[-1] < num:
                ans += st.pop()*min(num, st[-1])
            st.append(num)
        
        while len(st) > 2:
            ans += st.pop()*st[-1]
        return ans

# //1. 如果栈顶元素比当前元素小，弹出栈顶元素，
# //   栈顶元素与当前元素和栈顶下一个元素中得最小值组合
# //2. 如果栈顶元素比当前元素大，入栈