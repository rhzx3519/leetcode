class Solution(object):
    def findRightInterval(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[int]
        """
        first = [i[0] for i in intervals]
        last = [i[1] for i in intervals]
        first = sorted(enumerate(first), key=lambda x: x[1])
        last = sorted(enumerate(last), key=lambda x: x[1])
        n = len(intervals)
        out = [-1] * n
        i = j = 0
        while i < n:
            while j < n:
                if last[i][1] <= first[j][1]:
                    out[last[i][0]] = first[j][0]
                    break
                else:
                    j += 1
            if j==n:
                break
            else:
                i += 1
        return out

# 把区间的始末分别记录为两个列表，分别进行升序排列并记录下在原区间列表中的位置，之后只需要对比两个列表中的值就好啦。
# 升序排列的话可以避免重复搜索的时间。

