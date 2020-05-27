class Solution(object):
    def findTheLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        VOWEL = 'aeiou'
        N = len(VOWEL)
        mem = collections.defaultdict(int)
        mem[0] = -1
        state = 0
        res = 0

        for i, ch in enumerate(s):
            for j, v in enumerate(VOWEL):
                if (ch==v):
                    state ^= (1 << (N-1-j))
                    break

            if state in mem:
                res = max(res, i - mem.get(state))
                continue # 不更新 mem了，因为这时候mem[state]肯定比i小
            
            mem[state] = i
        return res

# 状态压缩＋哈希表。类似的题出现好几次了。 如1124。 状态压缩后，哈希表的用处是求最长的连续子串，满足子数组的和为k。 此题k为0， 1124题k为1.

# 遇到奇偶个数校验，想到XOR
# 遇到有限的参数（小于20个）表状态， 想到状态压缩 （bitmask）
# 遇到求最长的连续子串使得和为k（maximum continuous subarray(substring) with sum equal to k）想到 前缀和 加哈希表记录第一次出现某一状态的位置。
        