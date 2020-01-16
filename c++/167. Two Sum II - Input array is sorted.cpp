class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        vector<int> res;
        int l =  0, r = numbers.size() - 1, m;
        while (l <= r) {
            m = numbers[l] + numbers[r];
            if (m == target) {
                res.push_back(l+1);
                res.push_back(r+1);
                return res;
            } else if (m > target)
                r--;
            else
                l++;
        }
        return res;
    }
};