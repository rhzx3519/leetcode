class Solution(object):
    def escapeGhosts(self, ghosts, target):
        """
        :type ghosts: List[List[int]]
        :type target: List[int]
        :rtype: bool
        """
        dis = sum(map(abs, target))
        return all([dis <= abs(target[0] - g[0]) + abs(target[1] - g[1])for g in ghosts])