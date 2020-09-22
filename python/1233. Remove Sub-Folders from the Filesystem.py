class TrieTree(object):
    def __init__(self, val, is_leaf=False):
        self.val = val
        self.is_leaf = is_leaf
        self.child = {}


class Solution(object):
    def removeSubfolders(self, folder):
        """
        :type folder: List[str]
        :rtype: List[str]
        """
        ans = []
        root = TrieTree(-1)
        folder.sort(key=len)
        for fold in folder:
            is_sub = False
            node = root
            for w in fold.split('/'):
                if not w:
                    continue
                if w not in node.child:
                    node.child[w] = TrieTree(w)
                node = node.child[w]
                if node.is_leaf:
                    is_sub = True
            node.is_leaf = True
            if not is_sub:
                ans.append(fold)
        return ans