class Solution(object):
    def xorQueries(self, arr, queries):
        """
        :type arr: List[int]
        :type queries: List[List[int]]
        :rtype: List[int]
        """
        pre = [0] * (len(arr) + 1)
        for i, e in enumerate(arr):
            pre[i + 1] = pre[i] ^ e
        return [pre[a] ^ pre[b+1] for a, b in queries]

# 由
# a ^ b = b ^ a
# 0 ^ x = x
# x ^ x = 0
# 易知：
# d ^ e ^ f =
#  0   ^    (d ^ e ^ f) = 
# ((a ^ a)    ^    (b ^ b)    ^   (c ^ c))      ^      (d ^ e ^ f) = 
# (a ^ b ^ c)   ^   (a ^ b ^ c ^ d ^ e ^ f)