class Solution {
public:
    int rob(vector<int>& nums) {
        int ls = nums.size(), res = 0, t;
        vector<int> dp(3, 0);
        for (int i = 0; i < ls; i++) {
            t = 0;
            for (int j = 0; j < 2; j++)
                t = max(t, dp[j] + nums[i]);
            res = max(res, t);
            dp.push_back(t);
            dp.erase(dp.begin());
        }
        
        return res;
    }
};