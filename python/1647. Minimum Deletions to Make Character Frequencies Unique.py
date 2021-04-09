import collections

class Solution(object):
    def minDeletions(self, s):
        """
        :type s: str
        :rtype: int
        """
        ans = 0
        freq = collections.defaultdict(int)
        for ch, num in collections.Counter(s).iteritems():
            freq[num] += 1
        nums = sorted(freq.keys(), reverse=True)
        # print nums
        # print freq

        i = 0
        while i < len(nums) and nums[i] > 0:
            num = nums[i]
            if num-1 not in freq:
                nums.insert(i+1, num-1)

            t = max(0, freq[num] - 1)
            ans += t
            freq[num] = 1
            freq[num-1] += t
            i += 1
        # print nums
        # print freq
        return ans


if __name__ == '__main__':
    # s = "jbddhjemmnhaflahionjoddojoliimdcailihfdleahgbafnknblkheeicoonffenhhmgfhgmnjk"
    # s = "aab"
    # s = "aaabbbcc"
    s = "ceabaacb"
    su = Solution()
    print su.minDeletions(s)