class Solution(object):
    def validateStackSequences(self, pushed, popped):
        """
        :type pushed: List[int]
        :type popped: List[int]
        :rtype: bool
        """
        self.st = []
        n = len(pushed)
        j = 0
        for i in range(n):
            self.st.append(pushed[i])
            while j < n and self.st and self.st[-1] == popped[j]:
                self.st.pop()
                j += 1
        return not self.st
# 模拟栈的出入过程
# time:O(N), space: O(N)