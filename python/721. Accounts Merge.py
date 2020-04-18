class Solution(object):
    def accountsMerge(self, accounts):
        """
        :type accounts: List[List[str]]
        :rtype: List[List[str]]
        """
        def find(x):
            if x != look_up[x]:
                look_up[x] = find(look_up[x])
            return look_up[x]

        def join(x, y):
            px = find(x)
            py = find(y)
            if px != py:
                look_up[py] = px

        n = len(accounts)
        look_up = {i:i for i in range(n)}
        for idx, account in enumerate(accounts):
            name = account[0]

            for email in account[1:]:
                if email in look_up:
                    join(idx, look_up[email])
                else:
                    look_up[email] = idx

        from collections import defaultdict
        d = defaultdict(set)

        for idx, account in enumerate(accounts):
            tmp = find(idx)
            for email in account[1:]:
                d[tmp].add(email)

        res = []
        for idx, st in d.iteritems():
            res.append([accounts[idx][0]] + sorted(list(st)))

        return res

# 并查集, 父节点是account下标