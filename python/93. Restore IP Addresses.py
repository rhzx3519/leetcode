class Solution(object):
    def restoreIpAddresses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        res = []

        def dfs(start, end, path):
            if end > len(s) or len(path)>4: return
            sub = s[start:end]
            if len(sub)>1 and sub[0]=='0': return
            if int(sub)>255: return 
            path.append(sub)
            if end==len(s) and len(path)==4:
                res.append('.'.join(path))
                return
            
            for i in range(1, 4):
                dfs(end, end+i, path[:])

        for i in range(1, 4):
            dfs(0, i, [])
        return res

# 思路：定义start和end两个索引来进行dfs        