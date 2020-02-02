class Solution(object):
    def findSubstringInWraproundString(self, p):
        """
        :type p: str
        :rtype: int
        """
        n = len(p)
        count = collections.defaultdict(int)
        add = 0
        for i in range(n):
            if (i>0) and (ord(p[i-1]) + 1==ord(p[i]) or ord(p[i-1])-ord(p[i])==25):
                add += 1
            else:
                add = 1
            count[p[i]] = max(count[p[i]], add)
        
        res = 0
        for _, num in count.iteritems():
            res += num
        return res

        # 统计以每个字符作为结尾的最长连续序列(可以覆盖掉重复的短序列的情况), 他们的和即为所求
        # 例如:abcdbcd, 对于以d结尾的有abcd, bcd, cd和d, 而bcd产生的序列都会被abcd所覆盖
        # 总和即以a、b、c和d结尾的所有连续最长序列1 + 2 + 3 + 4 = 10