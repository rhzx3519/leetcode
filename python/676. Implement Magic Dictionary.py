class MagicDictionary(object):
    
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.look_up = {}
        

    def buildDict(self, dict):
        """
        Build a dictionary through a list of words
        :type dict: List[str]
        :rtype: None
        """
        for s in dict:
            n = len(s)
            l = self.look_up.setdefault(n, [])
            l.append(s)

    def search(self, word):
        """
        Returns if there is any word in the trie that equals to the given word after modifying exactly one character
        :type word: str
        :rtype: bool
        """
        n = len(word)
        for w in self.look_up.get(n, []):
            cnt = 0
            for i in xrange(n):
                if w[i] != word[i]:
                    cnt += 1
            if cnt==1:
                return True
            
        return False;


        


# Your MagicDictionary object will be instantiated and called as such:
# obj = MagicDictionary()
# obj.buildDict(dict)
# param_2 = obj.search(word)