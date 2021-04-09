class RandomizedCollection(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.cache = []
        self.mem = collections.defaultdict(set)


    def insert(self, val):
        """
        Inserts a value to the collection. Returns true if the collection did not already contain the specified element.
        :type val: int
        :rtype: bool
        """
        res = True
        if self.mem[val]:
            res = False
        self.cache.append(val)
        self.mem[val].add(len(self.cache)-1)
        return res


    def remove(self, val):
        """
        Removes a value from the collection. Returns true if the collection contained the specified element.
        :type val: int
        :rtype: bool
        """
        if not self.mem[val]:
            return False
        i = self.mem[val].pop()
        if i==len(self.cache)-1:
            self.cache.pop()
            return True
        lastVal = self.cache[-1]
        self.cache[i], self.cache[-1] = self.cache[-1], self.cache[i]
        # print lastVal, self.mem[lastVal]
        self.mem[lastVal].remove(len(self.cache)-1)
        self.mem[lastVal].add(i)
        self.cache.pop()
        return True


    def getRandom(self):
        """
        Get a random element from the collection.
        :rtype: int
        """
        # print self.cache
        return random.choice(self.cache)



# Your RandomizedCollection object will be instantiated and called as such:
# obj = RandomizedCollection()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()