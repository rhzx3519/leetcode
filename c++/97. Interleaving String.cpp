class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
        int m = s1.size(), n = s2.size();
        if (m + n != s3.size()) return false;
        
        vector<vector<bool> > dp(m + 1);
        for (int i = 0; i < dp.size(); i++) 
            dp[i].resize(n + 1);
        
        dp[0][0] = true;
        
        // 先考虑分别单独用S1或者S2能否构成S3
        for(int i = 1; i < m+1 && s1[i-1] == s3[i-1]; i++)
            dp[i][0] = true;
        for(int i = 1; i < n+1 && s2[i-1] == s3[i-1]; i++)
            dp[0][i] = true;
        
        for (int i = 1; i < m+1; i++) {
            for (int j = 1; j < n+1; j++) {
                if((s3[i+j-1] == s1[i-1]&&dp[i-1][j]) || 
                    (s3[i+j-1] == s2[j-1]&&dp[i][j-1]))
                    dp[i][j] = true;
            }
        }
        
        return dp[m][n];
    }
};