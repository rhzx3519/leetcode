class SortedStack(object):

    def __init__(self):
        self.st = []
        self.help_st = []

    def push(self, val):
        """
        :type val: int
        :rtype: None
        """
        if not self.st or self.st[-1]>val:
            self.st.append(val)
            return
        while self.st and self.st[-1]<val:
            t= self.st.pop()
            self.help_st.append(t)
        self.st.append(val)
        while self.help_st:
            t = self.help_st.pop()
            self.st.append(t)


    def pop(self):
        """
        :rtype: None
        """
        if self.st:
            self.st.pop()


    def peek(self):
        """
        :rtype: int
        """
        if not self.st:
            return -1
        return self.st[-1]


    def isEmpty(self):
        """
        :rtype: bool
        """
        return len(self.st)==0

# 思路：入栈时，进行排序


# Your SortedStack object will be instantiated and called as such:
# obj = SortedStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.peek()
# param_4 = obj.isEmpty()