class Solution {
public:
    int findTargetSumWays(vector<int>& nums, int S) {
        int sum = accumulate(nums.begin(), nums.end(), 0);
        if (sum < S || (S + sum)&1) return 0;

        int w = (S + sum) / 2;
        vector<int> dp(w+1, 0);
        dp[0] = 1;
        for (int num : nums) {
            for (int i = w; i >= num; i--) {
                dp[i] += dp[i-num];
            }
        }

        return dp[w];
    }
};

    /**
     * 494
     * 输入: nums: [1, 1, 1, 1, 1], S: 3
     * 输出: 5
     * 解释:
     * -1+1+1+1+1 = 3
     * +1-1+1+1+1 = 3
     * +1+1-1+1+1 = 3
     * +1+1+1-1+1 = 3
     * +1+1+1+1-1 = 3
     * 
     * sum(P) 前面符号为+的集合；sum(N) 前面符号为减号的集合
     * 所以题目可以转化为
     * sum(P) - sum(N) = target 
     * => sum(nums) + sum(P) - sum(N) = target + sum(nums)
     * => 2 * sum(P) = target + sum(nums) 
     * => sum(P) = (target + sum(nums)) / 2
     * 因此题目转化为01背包，也就是能组合成容量为sum(P)的方式有多少种
     */