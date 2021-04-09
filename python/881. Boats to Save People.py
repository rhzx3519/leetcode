class Solution(object):
    def numRescueBoats(self, people, limit):
        """
        :type people: List[int]
        :type limit: int
        :rtype: int
        """
        ans = 0
        n = len(people)
        l, r = 0, n-1
        people.sort()
        while l <= r:
            two = people[l] + people[r]
            if two <= limit:
                l += 1
                r -= 1
            else:
                r -= 1
            ans += 1
        return ans
        
