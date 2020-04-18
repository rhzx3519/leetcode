class Solution {
public:
    int largestRectangleArea(vector<int>& heights) {
        vector<int> &a = heights;
        if (a.empty()) return 0;
        int n = a.size();
        vector<int> ll(n, 0), rr(n, 0); // ll, rr 分别记录当前位置i能够向左和向右延伸的最大距离
        
        int i , j;
        ll[0] = 0; 
        for (i = 1; i < n; i++) {
            ll[i] = i;
            j = i;
            while (ll[j] - 1 >= 0 && a[ll[j]-1] >= a[i]) // if左边的高度大于等于a[i]，向左延伸一步
                j = ll[j] - 1;
            ll[i] = ll[j];
        }
        
        rr[n-1] = n-1;
        for (i = n-2; i >= 0; i--) {
            rr[i] = i;
            j = i;
            while (rr[j] + 1 < n && a[rr[j] + 1] >= a[i])
                j = rr[j] + 1;
            rr[i] = rr[j];
        }
        
        int res = 0;
        for (int i = 0; i < n; i++) {
            res = max(res, (rr[i] - ll[i] + 1) * a[i]);
        }
        
        return res;
    }
};