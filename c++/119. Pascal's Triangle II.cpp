class Solution {
public:
    vector<int> getRow(int rowIndex) {
        vector<int> res;
        if (rowIndex < 0) return res;
        vector<vector<int>> v(2);
        v[0].push_back(1);
        
        int f = 1;
        int nf = !f;
        for (int i = 2; i <= rowIndex + 1; i++) {
            v[f].resize(i);
            v[f][0] = v[nf][0];
            for (int j = 1; j < i - 1; j++) {
                v[f][j] = (v[nf][j-1] + v[nf][j]);
            }
            v[f][i-1] = v[nf][i-2];
            
            f = !f;
            nf = !nf;
        }
        
        return v[nf];
    }
};