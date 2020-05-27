class MinStack(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.st = []
        self.min_st = []


    def push(self, x):
        """
        :type x: int
        :rtype: None
        """
        self.st.append(x)
        if not self.min_st or (self.min_st and self.min_st[-1] >= x):
            self.min_st.append(x)


    def pop(self):
        """
        :rtype: None
        """
        if not self.st:
            return
            
        x = self.st.pop()
        if x==self.min_st[-1]:
            self.min_st.pop()


    def top(self):
        """
        :rtype: int
        """
        return self.st[-1]


    def getMin(self):
        """
        :rtype: int
        """
        return self.min_st[-1]



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()

if __name__ == '__main__':
    st = MinStack()
    st.push(-2)
    st.push(0)
    st.push(-3)
    st.getMin()
    print st.min_st
    st.pop()
    st.top()
    print st.min_st
    st.getMin()






