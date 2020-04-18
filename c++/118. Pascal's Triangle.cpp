class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> res;
        if (numRows == 0) return res;
        
        res.resize(numRows);
        res[0].push_back(1);
        for (int i = 1; i < numRows; i++) {
            res[i].push_back(res[i-1][0]);
            for (int j = 1; j < i; j++)
                res[i].push_back(res[i-1][j-1] + res[i-1][j]);
            res[i].push_back(res[i-1][i-1]);
        }
        
        return res;
    }
};