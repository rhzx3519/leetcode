class Solution(object):
    def corpFlightBookings(self, bookings, n):
        """
        :type bookings: List[List[int]]
        :type n: int
        :rtype: List[int]
        """
        diff = [0]*(n+1)
        for j, k, l in bookings:
            diff[j-1] += l
            diff[k] -= l

        nums = [0]*n
        nums[0] = diff[0]
        for i in range(1, n):
            nums[i] = nums[i-1] + diff[i]
        # print diff
        # print nums
        return nums