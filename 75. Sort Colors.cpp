class Solution {
public:
    void sortColors(vector<int>& nums) {
        int i = 0, l = 0, r = nums.size() - 1;
        for (i; i < min((int)nums.size(), r+1);) {
            if (nums[i] == 0) {
                if (i == l)
                    i++;
                else
                    swap(nums[i], nums[l]);
                l++;
            } else if (nums[i] == 2) {
                swap(nums[i], nums[r]);
                r--;
            } else
                i++;
        }
    }
};