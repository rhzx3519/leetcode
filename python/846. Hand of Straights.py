class Solution(object):
    def isNStraightHand(self, hand, W):
        """
        :type hand: List[int]
        :type W: int
        :rtype: bool
        """
        N = len(hand)
        if N%W != 0:
            return False
        hand.sort()
        count = collections.Counter(hand)

        for num in hand:
            t = num
            if count.get(t, 0) > 0:
                for i in range(t, t+W):
                    if not count.get(i, 0):
                        return False
                    count[i] -= 1
        return True
