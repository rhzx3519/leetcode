class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        vector<int> res;
        int add = 1, t;
        for (int i = digits.size() - 1; i >= 0; i--) {
            t = (digits[i] + add) % 10;
            add = (digits[i] + add) / 10;
            res.push_back(t);
        }
        if (add != 0)
            res.push_back(add);
        
        reverse(res.begin(), res.end());
        return res;
    }
};