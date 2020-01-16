class Solution {
public:
    int findLength(vector<int>& A, vector<int>& B) {
        int n1 = A.size(), n2 = B.size();
        vector<vector<int>> dp(n1+1, vector<int>(n2+1, 0));

        int res = 0;
        for (int i = 1; i <= n1; i++) {
            for (int j = 1; j <= n2; j++) {
                if (A[i-1]==B[j-1]) {
                    dp[i][j] = dp[i-1][j-1] + 1;
                    res= max(res, dp[i][j]);
                } else {
                    dp[i][j] = 0;
                }
            }
        }

        return res;
    }
};