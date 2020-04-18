class LFUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.cache = {}
        self.frequency = collections.defaultdict(list)
        self.capacity = capacity
        

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self.cache: return -1
        for freq in self.frequency:
            if key in self.frequency[freq]:
                self.frequency[freq].remove(key)
                if not self.frequency[freq]:
                    self.frequency.pop(freq)
                self.frequency[freq+1].append(key)
                break
        return self.cache[key]
        

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        if self.capacity <= 0: return
        if key not in self.cache:
            if len(self.cache) == self.capacity:
                pop_key = self.frequency[min(self.frequency)].pop(0)
                self.cache.pop(pop_key)
            self.cache[key] = value
            self.frequency[1].append(key)
        else:
            self.cache[key] = value
            self.get(key)
        


# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)