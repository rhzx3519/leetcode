class Solution(object):
    def filterRestaurants(self, restaurants, veganFriendly, maxPrice, maxDistance):
        """
        :type restaurants: List[List[int]]
        :type veganFriendly: int
        :type maxPrice: int
        :type maxDistance: int
        :rtype: List[int]
        """
        ans = []
        for r in restaurants:
            idi, ratingi, veganFriendlyi, pricei, distancei = r
            if veganFriendly and veganFriendly^veganFriendlyi:
                continue
            if pricei > maxPrice:
                continue
            if distancei > maxDistance:
                continue
            ans.append((idi, ratingi))
        
        ans.sort(key=lambda x: (-x[1], -x[0]))
        return [i for i, _ in ans]