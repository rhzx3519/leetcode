class Solution {
public:
    vector<vector<int>> updateMatrix(vector<vector<int>>& matrix) {
        vector<vector<int>> res;
        if (matrix.empty()) return res;
        int m = matrix.size(), n = matrix[0].size();

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 1)
                    matrix[i][j] = m + n;

                if (i > 0)
                    matrix[i][j] = min(matrix[i][j], matrix[i-1][j] + 1);
                if (j > 0)
                    matrix[i][j] = min(matrix[i][j], matrix[i][j-1] + 1);
                
            }
        }

        for (int i = m-1; i >= 0; i--) {
            for (int j = n-1; j >= 0; j--) {
                if (i < m-1)
                    matrix[i][j] = min(matrix[i][j], matrix[i+1][j] + 1);
                if (j < n-1)
                    matrix[i][j] = min(matrix[i][j], matrix[i][j+1] + 1);
            }
        }        

        return matrix;
    }

};

// dp， 两次遍历，一次从左上到右下，一次从右下到左上