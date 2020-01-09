class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int i, j;
        int n = nums.size();
        i = j = 0;
        while (i < n) {
            if (nums[i] != 0){
                nums[j++] = nums[i];
            }
            i++;
        }
        while (j < n) {
            nums[j++] = 0;
        }
    }
};