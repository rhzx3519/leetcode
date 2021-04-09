class Solution(object):
    def alertNames(self, keyName, keyTime):
        """
        :type keyName: List[str]
        :type keyTime: List[str]
        :rtype: List[str]
        """
        def foo(times):
            for i, t in enumerate(times):
                if i < 2:
                    continue
                if times[i] - times[i-2] <= 60:
                    return True
            return False

        mem = collections.defaultdict(list)
        n = len(keyTime)
        for i in range(n):
            name = keyName[i]
            time = keyTime[i]
            hour, minute = map(int, time.split(':'))
            mem[name].append(hour*60 + minute)

        ans = []
        for name, times in mem.iteritems():
            times.sort()
            if foo(times):
                ans.append(name)
        return sorted(ans)