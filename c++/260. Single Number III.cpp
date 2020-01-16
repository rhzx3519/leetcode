class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        int n = nums.size();
        int sum = 0;
        for (auto &i : nums) {
            sum ^= i;
        }
        
        sum = sum&(-sum);
        int n1 = 0, n2 = 0;
        for (auto &i : nums) {
            if (i&sum)
                n1 ^= i;
            else
                n2 ^= i;
        }
        
        vector<int> res = {n1, n2};
        return res;
    }
};