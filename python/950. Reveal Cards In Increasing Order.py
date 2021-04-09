class Solution(object):
    def deckRevealedIncreasing(self, deck):
        """
        :type deck: List[int]
        :rtype: List[int]
        """
        ans = [0]*len(deck)
        que = range(len(deck))
        for card in sorted(deck):
            ans[que.pop(0)] = card
            if que:
                que.append(que.pop(0))
        return ans

# 对deck排序，一次放入下标为[0, 2, 4...]的卡牌