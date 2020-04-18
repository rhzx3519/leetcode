class Solution {
public:
    int maximumGap(vector<int>& nums) {
        if(nums.empty() || nums.size() == 1)
            return 0;
        sort(nums.begin(), nums.end());
        int ret = 0;
        for(int i = 1; i < nums.size(); i ++)
            ret = max(ret, nums[i]-nums[i-1]);
        return ret;
    }
};