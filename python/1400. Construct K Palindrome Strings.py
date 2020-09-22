class Solution(object):
    def canConstruct(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: bool
        """
        return len(s)>=k and sum(map(lambda x: x%2, collections.Counter(s).values())) <= k