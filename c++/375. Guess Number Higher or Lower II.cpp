class Solution {
public:
    int getMoneyAmount(int n) {
        vector<vector<int>> dp (n+1, vector<int>(n+1, 0));
        return dfs(dp, 1, n);
    }
    
    int dfs(vector<vector<int>> &dp, int l, int r) {
        if (l >= r) return 0;
        if (dp[l][r] != 0)
            return dp[l][r];
        int res = INT_MAX;
        for (int i = l; i <= r; i++) {
            int t = i + max(dfs(dp, l, i-1), dfs(dp, i+1, r));
            res = min(res, t);
        }
        dp[l][r] = res;
        return res;
    }
};