class Solution(object):
    def lengthLongestPath(self, input):
        """
        :type input: str
        :rtype: int
        """
        st = [(-1, 0)] # st保存了目录深度和对应的字符串长度
        max_len = 0
        for p in input.split('\n'):
            depth = p.count('\t')
            p = p.replace('\t', '')
            while st and st[-1][0] >= depth:
                st.pop()
            
            if '.' in p:
                max_len = max(max_len, len(p) + st[-1][1])
            else:
                st.append((depth, len(p) + st[-1][1] + 1)) # +1 表示多一个/分隔符
        
        return max_len