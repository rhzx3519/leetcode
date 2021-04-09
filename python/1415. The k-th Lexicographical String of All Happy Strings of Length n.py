class Solution(object):
    def getHappyString(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """

        def dfs(st):
            if cnt[0]==k:
                return
            if len(st)==n:
                cnt[0] += 1
                ans[0] = ''.join(st)
                return
            for i in range(3):
                ch = chr(ord('a') + i)
                if st and st[-1]==ch: continue
                st.append(ch)
                dfs(st)
                st.pop()
        
        ans = ['']
        cnt = [0]
        dfs([])
        # print cnt,  ans
        return ans[0] if cnt[0]==k else ''

# 思路：回溯查找