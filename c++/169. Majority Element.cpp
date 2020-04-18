class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int res, cnt = 0;
        for (auto &i : nums) {
            if (cnt == 0) {
                res = i;
                cnt++;
            } else if (res == i)
                cnt++;
            else
                cnt--;
        }
        
        return res;
    }
};