class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        if (matrix.empty()) return 0;
        int m = matrix.size(), n = matrix[0].size();
        int res = 0;
        
        vector<vector<int>> a(m, vector<int>(n, 0));
        vector<int> ll(n), rr(n);
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int k = i, cnt = 0;
                while (k >= 0 && matrix[k][j] == '1') {
                    k--; cnt++;
                }
                a[i][j] = cnt;
            }
        }
        
        for (int i = 0; i < m; i++) {
            int j, k;
            
            ll[0] = 0;
            for (j = 1; j < n; j++) {
                ll[j] = j;
                k = j;
                while (ll[k] - 1 >= 0 && a[i][ll[k] - 1] >= a[i][j])
                    k = ll[k] - 1;
                ll[j] = ll[k];
            }
            
            rr[n-1] = n-1;
            for (j = n-2; j >= 0; j--) {
                rr[j] = j;
                k = j;
                while (rr[k] + 1 < n && a[i][rr[k] + 1] >= a[i][j])
                    k = rr[k] + 1;
                rr[j] = rr[k];
            }
            
            int max_s = 0;
            for (j = 0; j < n; j++) 
                max_s = max(max_s, (rr[j] - ll[j] + 1) * a[i][j]);
                
            res = max(max_s, res);
        }
        
        return res;
    }
};