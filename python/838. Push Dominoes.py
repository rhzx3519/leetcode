class Solution(object):
    def pushDominoes(self, dominoes):
        """
        :type dominoes: str
        :rtype: str
        """
        n = len(dominoes)

        ans = ''
        left = 0
        dominoes = 'L' + dominoes + 'R'
        for i in range(1, len(dominoes)):
            ch = dominoes[i]
            if ch not in ('R', 'L'):
                continue
            if ch == 'L' and dominoes[left] == 'L':
                ans += 'L'*(i-left-1)
            elif ch == 'L' and dominoes[left] == 'R':
                num = i - left - 1
                ans += 'R'*(num//2)
                if num&1:
                    ans += '.'
                ans += 'L'*(num//2)
            elif ch == 'R' and dominoes[left] == 'R':
                ans += 'R'*(i-left-1)
            elif ch == 'R' and dominoes[left] == 'L':
                ans += '.'*(i-left-1)
            ans += ch
            left = i
        return ans[:-1]
