class Solution(object):
    def validateBinaryTreeNodes(self, n, leftChild, rightChild):
        """
        :type n: int
        :type leftChild: List[int]
        :type rightChild: List[int]
        :rtype: bool
        """
        ind = [0] * n
        for i in range(n):
            if leftChild[i] >= 0:
                ind[leftChild[i]] += 1
            if rightChild[i] >= 0:
                ind[rightChild[i]] += 1
        root = -1
        for i in range(n):
            if ind[i]==0:
                root = i
                break
        if root == -1:
            return False
        vis = [0] * n
        vis[root] = 1
        que = [root]
        while que:
            root = que.pop(0)
            l, r = leftChild[root], rightChild[root]
            if l >= 0:
                if vis[l]:
                    return False
                vis[l] = 1
                que.append(l)

            if r >= 0:
                if vis[r]:
                    return False
                vis[r] = 1
                que.append(r)
        return all(vis)
        
# 对于一个包含 n 个节点 m 条边的无向图，如果它是一棵树，那么必须满足以下三个条件中的两个：
# m = n - 1；
# 该无向图连通；
# 该无向图无环。