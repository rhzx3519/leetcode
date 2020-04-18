class Solution(object):
    def monotoneIncreasingDigits(self, N):
        """
        :type N: int
        :rtype: int
        """
        s = str(N)
        num = []
        for c in s:
            num.append(int(c))
        
        last = len(s)
        for i in range(len(s)-1, 0, -1):
            if num[i-1] > num[i]:
                num[i-1] -= 1
                last = i
        num = num[:last] + [9]*(len(s) - last)
        num = int(''.join(str(i) for i in num))
        return num


'''
先从尾到头遍历一遍，如果出现逆序，大的值减1，并记录最后一个逆序的位置i，然后把从i开始到末尾的数全部变9
'''        