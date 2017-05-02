class Solution {
public:
    bool exist(vector<vector<char>>& board, string word) {
        if (board.empty() || word.empty()) return false;
        int m = board.size(), n = board[0].size();;
        vector<vector<int>> visit(m, vector<int>(n, 0));
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (word[0] == board[i][j]) {
                    if (dfs(0, board, word, visit, i, j))
                        return true;
                }
            }
        }
        
        return false;
    }
    
    bool dfs(int cur, vector<vector<char>>& board, string &word, vector<vector<int>>& visit, int i, int j) {
        if (cur == word.size()) {
            return true;
        }
        
        if (i >= 0 && j >= 0 && i <board.size() && j < board[i].size() && visit[i][j] == 0 && board[i][j] == word[cur]) {
            visit[i][j] = 1;
            int x, y;
            for (int k = 0; k < 4; k++) {
                x = i + dx[k];
                y = j + dy[k];
                if (dfs(cur+1, board, word, visit, x, y))
                    return true;
            }
            visit[i][j] = 0;
        }
        return false;
    }

private:
    vector<int> dx = {0,0,1,-1};
    vector<int> dy = {1,-1,0,0};
};