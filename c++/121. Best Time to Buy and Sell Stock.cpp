class Solution {
public:
    int maxProfit(vector<int>& prices) {
        auto &a = prices;
        int i, res = 0, min_val = INT_MAX;
        for (i = 0; i < a.size(); i++) {
            min_val = min(min_val, a[i]);
            res = max(res, a[i] - min_val);
        }
        return res;
    }
};