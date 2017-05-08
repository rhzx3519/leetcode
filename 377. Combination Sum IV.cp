class Solution {
public:
    int combinationSum4(vector<int>& nums, int target) {
        int n = nums.size();
        vector<int> dp(target + 1, -1);
        dp[0] = 1;
        
        return dfs(dp, nums, target);
    }
    
    int dfs(vector<int> &dp, vector<int>& nums, int target) {
        if (dp[target] != -1)
            return dp[target];
        
        int res = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] <= target)
                res += dfs(dp, nums, target - nums[i]);
        }
        
        dp[target] = res;
        return res;
    }
};