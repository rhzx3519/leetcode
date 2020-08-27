class TrieTreeNode(object):
    def __init__(self, val, is_word=False):
        self.val = val
        self.is_word = is_word
        self.children = [None]*26
        self.word = None
        

class Solution(object):
    def suggestedProducts(self, products, searchWord):
        """
        :type products: List[str]
        :type searchWord: str
        :rtype: List[List[str]]
        """
        self.CAPACITY = 3
        self.root = TrieTreeNode(0, False)
        for word in products:
            self.build(word)

        # a = []
        # self.dfs(self.root, a)
        # print a

        res = []
        for i, ch in enumerate(searchWord):
            res.append(self.query(searchWord[:i+1]))
        return res
            
    def query(self, prefix):
        res = []
        root = self.root
        for ch in prefix:
            if root.children[ord(ch) - ord('a')] is None:
                break
            node = root.children[ord(ch) - ord('a')]
            if node.is_word:
                res.append(node.word)
            root = node
        if len(res) >= 3:
            return res[:3]
        # print root and root.val
        if root.val == 0:
            return res
        self.dfs(root, res)
        if len(res) >= 3:
            res = res[:3]
        # print res
        return res

    def dfs(self, root, res):
        if not root:
            return

        for child in root.children:
            if child != None:
                if child.is_word:
                    res.append(child.word)
                self.dfs(child, res)

            
            
    def build(self, word):
        root = self.root
        for ch in word:
            idx = ord(ch) - ord('a')
            if root.children[idx] is None:
                node = TrieTreeNode(ch, False)
                root.children[idx] = node
            root = root.children[idx]
        root.is_word = True
        root.word = word

if __name__ == '__main__':
    # products = ["mobile","mouse","moneypot","monitor","mousepad"]
    # searchWord = "mouse"
    # products = ["havana"]
    # searchWord = "havana"
    # products = ["bags","baggage","banner","box","cloths"]
    # searchWord = "bags"
    products = ["havana"]
    searchWord = "tatiana"
    su = Solution()
    print su.suggestedProducts(products, searchWord)













