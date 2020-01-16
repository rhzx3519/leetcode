class Solution {
public:
    int findMin(vector<int>& nums) {
        int ls = nums.size(), l, r, m;
        l = 0;
        r = ls - 1;
        if (nums[l] <= nums[r])
            return nums[l];
        while (r - l > 1) {
            m = l + (r-l>>1);
            if (nums[m] > nums[l])
                l = m;
            else
                r = m;
        }
        
        return nums[r];
    }
};