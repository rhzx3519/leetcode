class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        vector<int> res;
        if (matrix.empty())
            return res;
        int li = 0, lj = 0, hi = matrix.size() - 1, hj = matrix[0].size()-1;
        while (li <= hi && lj <= hj) {
            int i = li, j = lj;
            while (j <= hj) {
                res.push_back(matrix[i][j]);
                j++;
            }
            i++;j--;
            while (i <= hi) {
                res.push_back(matrix[i][j]);
                i++;
            }
            i--;j--;
            while (j > lj && li != hi) {
                res.push_back(matrix[i][j]);
                j--;
            }
            
            while (i > li && lj != hj) {
                res.push_back(matrix[i][j]);
                i--;
            }
            
            li++;lj++;hi--;hj--;
        }
        
        return res;
    }
};