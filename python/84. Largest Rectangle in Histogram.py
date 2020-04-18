class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        res = 0
        heights = [0] + heights + [0]
        stack = []
        for i, h in enumerate(heights):
            while stack and heights[stack[-1]] > h:
                top = stack.pop()
                res = max(res, heights[top]*(i-stack[-1]-1))
            stack.append(i)
        

        return res

# 思路：维持一个单调递增栈   
# 这道题考虑用一个栈来维持一个递增序列，栈中保存的是当前元素的索引。如果当前值比栈顶元素小，那么栈中元素就需要弹出，直到栈为空或者栈顶元素小于当前元素。这样的话可以保证栈中第i个元素对应的高度总是小于等于栈中第i-1个元素至当前欲插入元素对应的矩阵高度，这样当元素弹出时，以它为最低高度的最大矩阵面积就很好求出来了。高直接可以得出，而长度则为第i-1个元素对应的索引至当前欲插入元素的索引的差值。
# ————————————————
# 版权声明：本文为CSDN博主「沧海漂游_」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
# 原文链接：https://blog.csdn.net/lv1224/article/details/79974175

        
if __name__ == '__main__':
    heights = [0, 9]
    su = Solution()
    print su.largestRectangleArea(heights)