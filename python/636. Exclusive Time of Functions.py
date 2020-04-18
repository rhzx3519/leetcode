class Solution(object):
    def exclusiveTime(self, n, logs):
        """
        :type n: int
        :type logs: List[str]
        :rtype: List[int]
        """
        a = [0]*n
        stack = []
        for log in logs:
            s = log.split(':')
            fid = int(s[0])
            t = int(s[-1])
            if s[1] == 'start':
                stack.append([fid, t])
            else:
                func = stack.pop()
                time = t - func[1] + 1
                a[fid] += time
                if stack:
                    a[stack[-1][0]] -= time
        
        return a
            

### 利用栈