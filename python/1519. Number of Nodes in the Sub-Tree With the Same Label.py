class Solution(object):
    def countSubTrees(self, n, edges, labels):
        """
        :type n: int
        :type edges: List[List[int]]
        :type labels: str
        :rtype: List[int]
        """
        tree = [[] for _ in range(n)]
        for a, b in edges:
            tree[a].append(b)
            tree[b].append(a)
        ans = [0]*n

        # print tree
        vis = [0]*n


        def post(idx):
            vis[idx] = 1

            cnt = {}
            for chd in tree[idx]:
                if vis[chd]: continue
                t = post(chd)
                for ch, num in t.iteritems():
                    cnt[ch] = cnt.get(ch, 0) + num

            ch =labels[idx]
            cnt[ch] = cnt.get(ch, 0) + 1
            ans[idx] = cnt[ch]
            print idx, cnt
            return cnt

        post(0)
        print ans
        return ans


if __name__ == '__main__':
    n = 7
    edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]
    labels = "abaedcd"

    n = 4
    edges = [[0,1],[1,2],[0,3]]
    labels = "bbbb"

    n = 5
    edges = [[0,1],[0,2],[1,3],[0,4]]
    labels = "aabab"

    su = Solution()
    su.countSubTrees(n, edges, labels)







                
