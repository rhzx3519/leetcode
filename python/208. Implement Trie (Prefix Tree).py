#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class TrieNode(object):
    """docstring for TrieNode"""
    def __init__(self, val, is_word=False, parent=None):
        super(TrieNode, self).__init__()
        self.val = val
        self.is_word = is_word
        self.children = [0]*26
        self.parent = parent



class Trie(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TrieNode(0)

    def insert(self, word):
        """
        Inserts a word into the trie.
        :type word: str
        :rtype: None
        """
        parent = self.root
        for i in range(len(word)):
            idx = ord(word[i]) - ord('a')
            if parent.children[idx]:
                child = parent.children[idx]
            else:
                child = TrieNode(word[i], False, parent) 
                parent.children[idx] = child            
            parent = child
        parent.is_word = True
            


    def search(self, word):
        """
        Returns if the word is in the trie.
        :type word: str
        :rtype: bool
        """
        parent = self.root
        for i in range(len(word)):
            idx = ord(word[i]) - ord('a')
            child = parent.children[idx]
            if not child:
                return False
            if child.is_word and i == len(word)-1:
                return True
            parent = child
        return False


    def startsWith(self, prefix):
        """
        Returns if there is any word in the trie that starts with the given prefix.
        :type prefix: str
        :rtype: bool
        """
        parent = self.root
        for i in range(len(prefix)):
            idx = ord(prefix[i]) - ord('a')
            child = parent.children[idx]
            if not child:
                return False
            parent = child
        return True

if __name__ == '__main__':
    trie = Trie();

    trie.insert("hello");
    print trie.search("hell");   # 返回 true
    print trie.search("helloa");     # 返回 false
    print trie.search("hello");     # 返回 false
    print trie.startsWith("hell"); # 返回 true
    print trie.startsWith("helloa"); # 返回 true
    print trie.startsWith("hello"); # 返回 true

# ["Trie","insert","search","search","search","startsWith","startsWith","startsWith"]
# [[],["hello"],["hell"],["helloa"],["hello"],["hell"],["helloa"],["hello"]]

# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)