class Solution {
public:
    bool PredictTheWinner(vector<int>& nums) {
        int n = nums.size();
        // dp[i][j]代表在i, j范围内，先手比后手拿到最好的多多少。
        vector<vector<int>> dp(n, vector<int>(n, 0));
        // 当i==j时，取得的在i-j范围内，最大值即为nums[i]；
        for(int i = 0; i < n; ++i){
            dp[i][i] = nums[i];
        }
        for(int end = 1; end < n; ++end) {
            for(int i = 0, j = end; j < n; ++i, ++j) {
                // 当先手拿num[i](左侧), 后手拿的最好值为 dp[i+1][j]
                // 当先手拿num[j](右侧), 后手拿的最好值为 dp[i][j-1]
                dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1]);
            }
        }
        // 返回dp[0][n-1],因为开始是0，最后一次是n-1,代表在0~n-1范围内先手是否能比后手多
        return dp[0][n-1] >= 0;
    }
};


// dp[i][j] 表示玩家在数组 i 到 j 区间内游戏能赢对方的最大值（i <= j）。当 i==j 时，dp[i][j] 显然等于 nums[i]。
// 当 i != j 时，若先手取左端 nums[i]，后手则为 dp[i+1][j] ，若先手取右端 nums[j]，后手则为dp[i][j-1]，故状态转移方程为 dp[i][j] = max(nums[i] - dp[i-1][j], nums[j] - dp[i][j-1])。