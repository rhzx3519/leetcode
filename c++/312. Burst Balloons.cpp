class Solution {
public:
    int maxCoins(vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        
        vector<vector<int>> dp(n, vector<int>(n, 0));
        dfs(nums, dp, 0, n - 1, 1, 1);
        int res = dp[0][n-1];
        dp.clear();
        return res;
    }
    
    void dfs(const vector<int>& nums, vector<vector<int>> &dp, int i, int j, int ll, int rr) {
        if (dp[i][j] != 0) return;
        if (i == j) {
            dp[i][j] = ll * nums[i] * rr;
            return;
        }
        int s1, s2;
        int res = 0;
        for (int k = i; k <= j; k++) {
            if (i < k && dp[i][k-1] == 0)
                dfs(nums, dp, i, k-1, ll, nums[k]);
            s1 = i <= k-1 ? dp[i][k-1] : 0;
            
            if (j > k && dp[k+1][j] == 0)
                dfs(nums, dp, k+1, j, nums[k], rr);
            s2 = k+1 <= j ? dp[k+1][j] : 0;
            
            res = max(res, s1 + s2 + ll*nums[k]*rr);
        }
        
        dp[i][j] =res;
    }
};