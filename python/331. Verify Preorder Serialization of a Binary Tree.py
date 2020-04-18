class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        nodes = preorder.split(',')
        diff = 1
        for x in nodes:
            diff -= 1
            if diff < 0:
                return False
            if x != '#':
                diff += 2
        return diff == 0

# 也可以理解为出度入度相等问题：我命名有问题，根结点的入度为0出度为2，
# 其他非叶子结点的入度为1出度为2，叶子节点入度为1出度为0。因为根节点多出来一个出度，
# 所以初始化度为1，一个非叶子节点时度+1，加入一个空节点（叶子节点）时度-1，如果度为0，
# 即达到出度入度相等，已经形成一颗二叉树。