class Solution {
public:
    int maxProfit(vector<int>& prices, int fee) {
        int n = prices.size();
        if (n < 2) return 0;
        vector<int> dp1(n, 0); // 第i天手上没有股票
        vector<int> dp2(n, 0); // 第i天手上有股票
        dp2[0] = -prices[0];
        for (int i = 1; i < n; i++) {
            dp1[i] = max(dp1[i-1], dp2[i-1] + prices[i] - fee);
            dp2[i] = max(dp2[i-1], dp1[i-1] - prices[i]);
        }

        return dp1[n-1];
    }
};