class Solution {
public:
    int numDistinct(string s, string t) {
        int lt = t.size();
        vector<int> dp(lt+1, 0);
        dp[0] = 1;
        for (int i = 1; i <= s.size(); i++)
            for (int j = lt; j > 0; j--)
                dp[j] += s[i-1]==t[j-1]?dp[j-1]:0;
        
        return dp[lt];
    }
};