class Solution(object):
    def largestTimeFromDigits(self, arr):
        """
        :type arr: List[int]
        :rtype: str
        """
        # 0~23, 0~59
        def valid(a):
            hour = int(''.join(map(str, a[:2])))
            minute = int(''.join(map(str, a[2:])))
            return hour, minute

        self.max_val = -1
        self.ans = ''

        def dfs(track):
            if len(track)==len(arr):
                # print map(lambda i: arr[i], track)
                hour, minute = valid(map(lambda i: arr[i], track))
                if 0<=hour<=23 and 0<=minute<=59:
                    # print hour, minute
                    t = hour*60 + minute
                    if t > self.max_val:
                        self.max_val = t
                        self.ans = '{:0>2d}:{:0>2d}'.format(hour, minute)
                return
            for i in range(len(arr)):
                if i in track:
                    continue
                track.append(i)
                dfs(track)
                track.pop()

        dfs([])
        return self.ans