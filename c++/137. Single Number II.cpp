class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int count[32] = {0}, res = 0;
        for (int i = 0; i < 32; i++) {
            for (int j = 0; j < nums.size(); j++) {
                count[i] += (nums[j]>>i&1);
            }
            count[i] %= 3;
            res |= (count[i]<<i);
        }
         
        return res;   
    }
};