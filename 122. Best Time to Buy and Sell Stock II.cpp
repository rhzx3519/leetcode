class Solution {
public:
    int maxProfit(vector<int>& prices) {
        auto &a = prices;
        int res = 0;
        for (int i = 1; i < a.size(); i++) {
            if (a[i] > a[i-1])
                res += a[i] - a[i-1];
        }
        return res;
    }
};