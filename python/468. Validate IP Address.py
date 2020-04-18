class Solution(object):
    def validIPAddress(self, IP):
        """
        :type IP: str
        :rtype: str
        """
        if self.isIPv4(IP): return 'IPv4'
        if self.isIPv6(IP): return 'IPv6'
        return 'Neither'

    def isIPv4(self, IP):
        nums = IP.split('.')
        if len(nums) != 4: return False
        for num in nums:
            if not num.isdigit(): return False
            if len(num)>1 and num[0]=='0': return False
            if int(num)<0 or int(num)>255: return False
        return True
        
    def isIPv6(self, IP):    
        nums = IP.split(':')
        if len(nums) != 8: return False
        for num in nums:
            if not num or len(num)>4: return False
            for ch in num:
                if not (ch.isdigit() or ('a'<=ch<='f') or ('A'<=ch<='F')):
                    return False
        return True