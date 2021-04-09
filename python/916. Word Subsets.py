class Solution(object):
    def wordSubsets(self, A, B):
        """
        :type A: List[str]
        :type B: List[str]
        :rtype: List[str]
        """
        ans = []
        bcount = {}
        for b in B:
            cnt = collections.Counter(b)
            for ch, num in cnt.iteritems():
                bcount[ch] = max(bcount.get(ch, 0), num)
        
        for a in A:
            acount = collections.Counter(a)
            f = True
            for ch in bcount.keys():
                if bcount[ch] > acount.get(ch, 0):
                    f = False
                    break
            if f:
                ans.append(a)
        return ans