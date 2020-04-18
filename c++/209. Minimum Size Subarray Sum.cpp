class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int n = nums.size(), sum = 0, i, j;
        for (i = 0; i < n; i++)
            sum += nums[i];
            
        if (sum < s)
            return 0;
        
        int res = INT_MAX;
        i = j = sum = 0;
        while (j < n) {
            while (j < n && sum < s) {
                sum += nums[j++];
            }
            
            while (i < j && sum >= s) {
                res = min(res, j - i);
                sum -= nums[i++];
            }
        }
        
        return res;
    }
};