class Solution(object):
    def avoidFlood(self, rains):
        """
        :type rains: List[int]
        :rtype: List[int]
        """
        cnt0 = []
        n = len(rains)
        cnt = collections.defaultdict(list)
        ans = [-1] * n
        for i, x in enumerate(rains):
            if x == 0:
                cnt0.append(i)
                continue

            cnt[x].append(i)
            if len(cnt[x]) >= 2:
                
                idx = self.can(cnt0, cnt[x][0])
                # print cnt0, cnt[x], idx
                if idx != -1:
                    cnt[x].pop(0)
                    ans[cnt0.pop(idx)] = x
                    continue
                return []
        while cnt0:
            ans[cnt0.pop()] = 1
        return ans
    
    def can(self, cnt0, t):
        if not cnt0:
            return -1
        
        for i, x in enumerate(cnt0):
            if x > t:
                return i
        return -1

                
        