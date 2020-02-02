# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):

    def PreOrder(self, root):
    	res = []
    	st = []
    	while st or root:
    		if root:
    			res.append(root.val)
    			st.append(root)
    			root = root.left
    		else:
    			root = st.pop()
    			root = root.right
    	return res

    def InOrder(self, root):
    	res = []
    	st = []
    	while st or root:
    		if root:
    			st.append(root)
    			root = root.left
    		else:
    			root = st.pop()
    			res.append(root.val)
    			root = root.right
    	return res

    def PostOrder(self, root):
    	res = []
    	st = []
    	last = None
    	while st or root:
    		if root:
    			st.append(root)
    			root = root.left
    		else:
    			root = st[-1]
    			if root.right and root.right!=last:
    				root = root.right
    			else:
    				last = root
    				res.append(root.val)
    				st.pop()
    				root = None

    	return res

if __name__ == '__main__':
	root = TreeNode(1)
	root.left = TreeNode(2)
	root.right = TreeNode(3)
	root.right.left = TreeNode(4)
	root.right.right = TreeNode(5)
	su = Solution()
	print su.PostOrder(root)    				






