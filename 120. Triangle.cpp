class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        vector<vector<int>>& a = triangle;
        if (a.empty()) return 0;
        
        int n = a.size();
        for (int i = n-2; i >= 0; i--)
        {
            for (int j = 0; j <= i; j++)
                a[i][j] += min(a[i+1][j], a[i+1][j+1]);
        }
        
        return a[0][0];
    }
};