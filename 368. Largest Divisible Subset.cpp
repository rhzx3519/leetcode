class Solution {
public:
    vector<int> largestDivisibleSubset(vector<int>& nums) {
        vector<int> res;
        int n = nums.size();
        vector<int> dp(n, 0), parent(n, 0);
        int mx = 0, mx_idx = 0;
        sort(nums.begin(), nums.end());
        
        for (int i = n-1; i >= 0; i--) {
            for (int j = i; j <= n-1; j++) {
                if (nums[j]%nums[i] == 0 && dp[i] < dp[j] + 1) {
                    dp[i] = dp[j] + 1;
                    parent[i] = j;
                    if (dp[i] > mx) {
                        mx = dp[i];
                        mx_idx = i;
                    }
                }
            }
        }
        
        for (int i = 0; i < mx; i++) {
            res.push_back(nums[mx_idx]);
            mx_idx = parent[mx_idx];
        }
        
        return res;
    }
};