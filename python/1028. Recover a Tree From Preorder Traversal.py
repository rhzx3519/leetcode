#!usr/bin/env python  
#-*- coding:utf-8 _*-  


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def recoverFromPreorder(self, S):
        """
        :type S: str
        :rtype: TreeNode
        """
        st = []
        self.root = None

        def dfs(S):
            if not S:
                return
            n = len(S)
            d = 0
            while d<n and S[d]=='-':
                d += 1

            j = d
            while j<n and S[j]!='-':
                j += 1
            val = int(S[d:j])

            cur_node = TreeNode(val)

            if d == len(st) - 1:
                st.pop()
                st[-1].right = cur_node
            elif d == len(st):
                if not self.root:
                    self.root = cur_node
                else:
                   st[-1].left = cur_node
            else:
                while d < len(st):
                    st.pop()
                st[-1].right = cur_node

            st.append(cur_node)
            dfs(S[j:])

        dfs(S)
        return self.root


    def preorder(self, root):
        if not root: return
        print root.val
        self.preorder(root.left)
        self.preorder(root.right)

    def level_order(self, root):
        que = [root]
        while que:
            sz = len(que)
            while sz:
                sz -= 1
                root = que.pop(0)
                print root.val
                if root.left:
                    que.append(root.left)
                if root.right:
                    que.append(root.right)




if __name__ == '__main__':
    # S = "1-2--3--4-5--6--7"
    # S = "1-2--3---4-5--6---7"
    S = "1-401--349---90--88"
    su = Solution()
    root = su.recoverFromPreorder(S)
    su.level_order(root)

        
