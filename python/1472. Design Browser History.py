class BrowserHistory(object):

    def __init__(self, homepage):
        """
        :type homepage: str
        """
        self.st = [homepage]
        self.cur = 0


    def visit(self, url):
        """
        :type url: str
        :rtype: None
        """
        self.st = self.st[:self.cur+1] + [url]
        self.cur = len(self.st) - 1


    def back(self, steps):
        """
        :type steps: int
        :rtype: str
        """
        self.cur = max(0, self.cur - steps)
        return self.st[self.cur]


    def forward(self, steps):
        """
        :type steps: int
        :rtype: str
        """
        self.cur = min(len(self.st) - 1, self.cur + steps)
        return self.st[self.cur]



# Your BrowserHistory object will be instantiated and called as such:
# obj = BrowserHistory(homepage)
# obj.visit(url)
# param_2 = obj.back(steps)
# param_3 = obj.forward(steps)