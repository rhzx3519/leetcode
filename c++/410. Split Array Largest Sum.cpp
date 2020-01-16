/*
	记忆化搜索+dp
	定义函数find(i, m)，返回从i位置开始，分成m组的解。
	那么这个子问题就等于，min{max{A[i], find(i + 1, m - 1)},  
	max{A[i] + A[i + 1], find(i + 2, m - 1)} , ...... }。
*/
class Solution {
public:
    int splitArray(vector<int>& nums, int m) {
        int n = nums.size();
        vector<vector<int>> dp(n, vector<int>(m+1, 0));
        vector<int> prefix(n, 0);
        prefix[n-1] = nums[n-1];
        for (int i = n-2; i >= 0; i--)
            prefix[i] = nums[i] + prefix[i+1];
        return foo(0, nums, dp, prefix, m);
    }
    
    int foo(int idx, vector<int>& nums, vector<vector<int>>& dp, vector<int>& prefix, int m)
    {
        if (m == 1) return prefix[idx];
        if (dp[idx][m] > 0) return dp[idx][m];	// 已经访问过的直接返回
        
        int min_val = INT_MAX, sum = 0;
        for (int i = idx; i <= nums.size() - m; i++)
        {
            sum += nums[i];
            min_val = min(min_val, max(sum, foo(i+1, nums, dp, prefix, m-1)));  
        }
        dp[idx][m] = min_val;
        return dp[idx][m];
    }
};