class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        mp = [
                ['', 'I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX'],
                ['', 'X', 'XX', 'XXX', 'XL', 'L', 'LX', 'LXX', 'LXXX', 'XC'],
                ['', 'C', 'CC', 'CCC', 'CD', 'D', 'DC', 'DCC', 'DCCC', 'CM'],
                ['', 'M', 'MM', 'MMM']
            ]
        a = []
        while num:
            a.append(num%10)
            num /= 10
        i = 0
        res = ''
        while i < len(a):
            res = mp[i][a[i]] + res
            i += 1
        
        return res
        