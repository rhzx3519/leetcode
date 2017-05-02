class Solution {
public:
    int rob(vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        
        int res1 = nums[0];
        vector<int> dp(3, 0);
        // 打劫第一个房子
        for (int i = 0; i < n - 1; i++) {
            int t = 0;
            for (int j = 0; j < 2; j++) 
                t = max(t, dp[j] + nums[i]);
            res1 = max(t, res1);
            dp.push_back(t);
            dp.erase(dp.begin());
        }
        if (n == 1)
            return res1;
        int res2 = nums[1];
        // 不打劫第一个房子
        dp.clear();
        dp.resize(3, 0);
        for (int i = 1; i < n; i++) {
            int t = 0;
            for (int j = 0; j < 2; j++) 
                t = max(t, dp[j] + nums[i]);
            res2 = max(t, res2);
            dp.push_back(t);
            dp.erase(dp.begin());
        }
        return(max(res1, res2));
    }
};