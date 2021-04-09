class Solution(object):
    def flipgame(self, fronts, backs):
        """
        :type fronts: List[int]
        :type backs: List[int]
        :rtype: int
        """
        cards = zip(fronts, backs)
        st = set()
        for f, b in cards:
            if f==b:
                st.add(f)
        ans = float('inf')
        for f, b in cards:
            if f not in st:
                ans = min(ans, f)
            if b not in st:
                ans = min(ans, b)
        return ans if ans!=float('inf')  else 0

# 思路：去掉那些正反面相等的卡片和数字, 在剩下的数字中选最小值