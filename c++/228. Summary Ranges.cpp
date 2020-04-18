class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector<string> res;
        int n = nums.size();
        if (n == 0) return res;
        
        int start = 0, i;
        string str;
        for (i = 1; i < n; i++) {
            if (nums[i-1] + 1 != nums[i]) {
                if (start == i-1)
                    str = to_string(nums[start]);
                else
                    str = to_string(nums[start]) + "->" + to_string(nums[i-1]);
                start = i;
                res.push_back(str);
            }
        }
        
        if (start < n) {
            if (start == n-1)
                str = to_string(nums[start]);
            else
                str = to_string(nums[start]) + "->" + to_string(nums[n-1]);
            res.push_back(str);
        }
        
        return res;
    }
}; 