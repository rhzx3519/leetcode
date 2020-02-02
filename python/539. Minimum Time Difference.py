class Solution(object):
    def findMinDifference(self, timePoints):
        """
        :type timePoints: List[str]
        :rtype: int
        """
        a = []
        for t in timePoints:
            t = t.split(':')
            minutes = int(t[0])*60 + int(t[1])

            a.append(minutes)

        a = sorted(a)
        res = 1<<30
        for i in range(len(a)):
            if i+1>=len(a):
                interval = 24*60 + a[(i+1)%len(a)] - a[i]
                print interval, a[(i+1)%len(a)]
            else:
                interval = a[(i+1)%len(a)] - a[i]
            res = min(interval, res)

        return res

if __name__ == '__main__':
    timePoints = ["23:59","01:00"]
    su = Solution()
    print su.findMinDifference(timePoints)        