class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
           
        vector<vector<int>> res(n, vector<int>(n, 0));
        int li = 0, lj = 0, ri = n-1, rj = n-1;
        int x = 1;
        while (li <= ri && lj <= rj) {
            int i = li, j = lj;
            while (j <= rj) {
                res[i][j] = x++;
                j++;
            }
            i++; j--;
            while (i <= ri) {
                res[i][j] = x++;
                i++;
            }
            j--; i--;
            while (j > lj && lj != rj) {
                res[i][j] = x++;
                j--;
            }
            while (i > li && li != ri) {
                res[i][j] = x++;
                i--;
            }
            
            li++; lj++; ri--; rj--;
        }
        
        return res;
    }
};