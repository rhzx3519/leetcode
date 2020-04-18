#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self.cache = {}
        self.que = []
        

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self.cache:
            return -1
        # 更新优先级
        idx = self.que.index(key)
        self.que.pop(idx)
        self.que.insert(0, key)
        return self.cache[key]
        

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        self.cache[key] = value
        if key in self.que:
            idx = self.que.index(key)
            self.que.pop(idx)
        self.que.insert(0, key)    
        if len(self.que) > self.capacity:
            lastKey = self.que[-1]
            del self.cache[lastKey]
            self.que.pop()


        


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)