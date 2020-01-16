class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        int ls = s.size();
        vector<bool> dp(ls + 1, false);
        dp[0] = true;
        for (int i = 1; i <= ls; i++) {
            for (int j = i - 1; j >= 0; j--) {
                if (dp[j] && find(wordDict.begin(), wordDict.end(), s.substr(j, i - j)) != wordDict.end())
                    dp[i] = true;
            }
        }
        return dp[ls];
    }
};