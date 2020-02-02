class Solution(object):
    def maximumSwap(self, num):
        """
        :type num: int
        :rtype: int
        """
        a = map(int, list(str(num)))
        for i in range(len(a)):
            idx = i
            for j in range(i+1, len(a)):
                if a[j] >= a[idx]:
                    idx = j
            
            if idx != i and a[idx]!=a[i]:
                a[i], a[idx] = a[idx], a[i]
                break
        return ''.join(map(str, a))
        

        