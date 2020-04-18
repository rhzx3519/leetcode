class Solution {
public:
    int calculateMinimumHP(vector<vector<int>>& dungeon) {
        if (dungeon.empty()) return 0;
        vector<vector<int>> &d = dungeon;
        int m = d.size(), n = d[0].size();
        vector<vector<int>> dp(m, vector<int>(n, 0));
        dp[m-1][n-1] = max(1 - d[m-1][n-1], 1);
        
		for (int i = m-2; i >= 0; i--)
			dp[i][n-1] = max(1, dp[i+1][n-1] - d[i][n-1]);
			
		for (int i = n-2; i >= 0; i--)
			dp[m-1][i] = max(1, dp[m-1][i+1] - d[m-1][i]);
        
        for (int i = m-2; i >= 0; i--)
            for (int j = n-2; j >= 0; j--)
                dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1]) - d[i][j]);

        
        return dp[0][0];
    }
};