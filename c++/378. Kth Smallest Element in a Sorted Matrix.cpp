class Solution {
public:
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        if (matrix.empty()) return 0;
        int m = matrix.size(), n = matrix[0].size();
        int l = matrix[0][0], r = matrix[m-1][n-1];
        
        while (l < r) {
            int mid = l + (r - l)/2;
            int count = 0, j = n - 1;
            for (int i = 0; i < m; i++) {
                while (j >= 0 && matrix[i][j] > mid) j--;
                count += (j+1);
            }
            if (count < k) l = mid + 1;
            else r = mid;
        }
        
        return l;
    }
};