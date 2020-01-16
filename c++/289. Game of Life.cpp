class Solution {
public:
    void gameOfLife(vector<vector<int>>& board) {
        if (board.empty()) return;
        int m = board.size(), n = board[0].size();
        int i, j, k;
        int cnt;
        int x, y;
        
        for (i = 0; i < m; i++) {
            for (j = 0; j < n; j++) {
                cnt = 0;
                for (k = 0; k < 8; k++) {
                    x = i + offset[k][0];
                    y = j + offset[k][1];
                    if (x < 0 || x > m-1) continue;
                    if (y < 0 || y > n-1) continue;
                    if (board[x][y]&1)
                        cnt++;
                }
                if (board[i][j]&1) {
                    if (cnt >= 2 && cnt <= 3)
                        board[i][j] |= 2;
                } else {
                    if (cnt == 3)
                        board[i][j] |= 2;
                }
                
            }
        }
        
        for (i = 0; i < m; i++)
            for (j = 0; j < n; j++)
                board[i][j] >>= 1;
    }
    
private:
    int offset[8][2] = {{1,0}, {-1,0}, {0,1}, {0,-1},{-1,-1}, {-1,1}, {1,-1}, {1,1}};
};