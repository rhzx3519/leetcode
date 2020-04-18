class Solution {
public:
    int firstUniqChar(string s) {
        int dp[26] = {0};
        for (int i = 0; i < s.size(); i++)
            dp[s[i] - 'a']++;
        for (int i = 0; i < s.size(); i++)
            if (dp[s[i] - 'a'] == 1)
                return i;
        return -1;
    }
};