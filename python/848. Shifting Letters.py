class Solution(object):
    def shiftingLetters(self, S, shifts):
        """
        :type S: str
        :type shifts: List[int]
        :rtype: str
        """
        n = len(S)
        for i in range(n-2, -1, -1):
            shifts[i] += shifts[i+1]
        print shifts
        ans = []
        for i in range(n):
            # print ((ord(S[i]) - ord('a') + d[i]) % 26)
            ans.append(chr((ord(S[i]) - ord('a') + shifts[i]) % 26 + ord('a')))
        return ''.join(ans)

if __name__ == '__main__':
    S = "abc"
    shifts = [3,5,9]
    su = Solution()
    print su.shiftingLetters(S, shifts)