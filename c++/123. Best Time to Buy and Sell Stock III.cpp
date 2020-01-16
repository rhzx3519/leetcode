class Solution {
public:
    int maxProfit(vector<int>& prices) {
        auto &a = prices;
        int n = a.size();
        if (n < 2) return 0;
        
        vector<int> v1(n, 0), v2(n, 0);
        
        int res = 0;
        int min_val = a[0];
        for (int i = 0; i < n; i++) {
            res = max(res, a[i] - min_val);
            v1[i] = res;
            min_val = min(min_val, a[i]);
        }
        
        res = 0;
        int max_val = a[n-1];
        for (int i = n-1; i >= 0; i--) {
            res = max(res, max_val - a[i]);
            v2[i] = res;
            max_val = max(max_val, a[i]);
        }
        
        res = 0;
        for (int i = 0; i < n; i++)
            res = max(res, v1[i]+v2[i]);
        
        return res;
    }
};