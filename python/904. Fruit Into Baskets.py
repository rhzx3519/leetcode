class Solution(object):
    def totalFruit(self, tree):
        """
        :type tree: List[int]
        :rtype: int
        """
        n = len(tree)
        l = 0
        ans = 0
        st = collections.defaultdict(int)
        for r in range(n):
            st[tree[r]] += 1
            while len(st) > 2:
                st[tree[l]] -= 1
                if st[tree[l]] == 0:
                    del st[tree[l]]
                l += 1
            ans = max(ans, r - l + 1)
        return ans