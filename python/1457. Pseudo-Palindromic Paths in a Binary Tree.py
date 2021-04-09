# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def pseudoPalindromicPaths (self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.ans = 0

        def foo(a):
            n = len(a)
            c = collections.Counter(a)
            cnt = 0
            for ch, num in c.iteritems():
                if num%1 == 0: continue
                else:
                    if n%1 == 0: return False
                    if cnt > 1: return False
                    cnt += 1
            return True

        def traverse(root, path):
            if not root: return
            if not root.left and not root.right:
                # print path + [root.val]
                self.ans += foo(path + [root.val])
                return
            
            if root.left:
                traverse(root.left, path + [root.val])
            if root.right:
                traverse(root.right, path + [root.val])
        
        traverse(root, [])
        return self.ans

import collections

def foo(a):
    n = len(a)
    c = collections.Counter(a)
    cnt = 0
    for ch, num in c.iteritems():
        if num&1 == 0: continue
        else:
            if n&1 == 0: return False
            if cnt > 1: return False
            cnt += 1
    return True

if __name__ == '__main__':
    print foo([1,2])










