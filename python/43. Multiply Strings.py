class Solution(object):
    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        n = len(num1) + len(num2)
        c = [0]*n
        for i in range(len(num1)-1, -1, -1):
            for j in range(len(num2)-1, -1, -1):
                c[i+j+1] += (ord(num1[i]) - ord('0')) * (ord(num2[j]) - ord('0'))

        for i in range(n-1, 0, -1):
            c[i-1] += c[i]/10
            c[i] %= 10

        i = 0
        while i<n and c[i]==0:
            i += 1
        if i==n: return "0"
        return ''.join([str(i) for i in c[i:]])

if __name__ == '__main__':
    su = Solution()
    print su.multiply("0", "0")        