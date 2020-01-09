class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        vector<vector<int>>& t = triangle;
        if (t.empty()) return 0;
        
        for (int i = t.size() - 2; i >= 0; i--)
            for (int j = 0; j <= i; j++)
                t[i][j] += min(t[i+1][j], t[i+1][j+1]);

        return t[0][0];
    }
};
