class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int add = 1;
        for (int i = digits.size() - 1; i >= 0; i--)
        {
            int t = digits[i] + add;
            digits[i] = t%10;
            add = t/10;
        }
        if (add != 0)
            digits.insert(digits.begin(), add);
        return digits;
    }
};