class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        nodes = preorder.split(',')
        diff = 1
        for x in nodes:
            diff -= 1
            if diff < 0:
                return False
            if x != '#':
                diff += 2
        return diff == 0