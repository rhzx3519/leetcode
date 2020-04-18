class MyStack(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.a = [] # 输入队列
        self.b = [] # 输出队列


    def push(self, x):
        """
        Push element x onto stack.
        :type x: int
        :rtype: None
        """
        self.a.append(x)
        while self.b:
            self.a.append(self.b.pop(0))
        self.a, self.b = self.b, self.a


    def pop(self):
        """
        Removes the element on top of the stack and returns that element.
        :rtype: int
        """
        return self.b.pop(0)


    def top(self):
        """
        Get the top element.
        :rtype: int
        """
        return self.b[0]


    def empty(self):
        """
        Returns whether the stack is empty.
        :rtype: bool
        """
        return not bool(self.b)



# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()