class Solution(object):
    def shipWithinDays(self, weights, D):
        """
        :type weights: List[int]
        :type D: int
        :rtype: int
        """
        def canShip(weights, D, K):
            cur = K
            for w in weights:
                if w > K:
                    return false
                if w > cur:
                    cur = K
                    D -= 1
                cur -= w
            return D > 0

        r = sum(weights)
        l = max(weights)
        while l <= r:
            mid = l + (r - l)/2
            if canShip(weights, D, mid):
                r = mid - 1
            else:
                l = mid + 1
        return l


