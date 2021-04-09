class Solution(object):
    def pathInZigZagTree(self, label):
        """
        :type label: int
        :rtype: List[int]
        """
        a = []
        while label != 1:
            a.append(label)
            label >>= 1
            label = label ^ (1 << (label.bit_length() - 1)) - 1
        return [1] + a[::-1]

# 求满二叉树的父节点下标：father(x) = x>>1
# 求之字形满二叉树的父节点下标: father(x) = x右移一位，除首位外，其余位取反