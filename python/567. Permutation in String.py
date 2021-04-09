import collections

class Solution(object):
    def checkInclusion(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: bool
        """
        k, n2 = len(s1), len(s2)
        if k > n2: return False

        pt = collections.defaultdict(int)
        for i in range(k):
            pt[s1[i]] += 1

        count = collections.defaultdict(int)
        for i in range(n2):
            count[s2[i]] += 1
            if i < k-1: continue
            # print i, count

            if pt==count: return True
            last = s2[i-k+1]
            count[last] -= 1
            if count[last] == 0:
                del count[last]
                
        return False

if __name__ == '__main__':
    s1 = "ab"
    s2 = "eidbaooo"
    su = Solution()
    print su.checkInclusion(s1, s2)