class Solution {
public:
    void solve(vector<vector<char>>& board) {
        if (board.empty()) return;
        int m = board.size(), n = board[0].size();
        vector<vector<bool>> visit(m, vector<bool>(n, false));
        
        for (int i = 0; i < m; i++)
        {
            if (board[i][0] == 'O')
                dfs(board, i, 0, visit);
        }
        
        for (int i = 0; i < m; i++)
        {
            if (board[i][n-1] == 'O')
                dfs(board, i, n-1, visit);
        }
        
        for (int i = 0; i < n; i++)
        {
            if (board[0][i] == 'O')
                dfs(board, 0, i, visit);
        }
        
        for (int i = 0; i < n; i++)
        {
            if (board[m-1][i] == 'O')
                dfs(board, m-1, i, visit);
        }
        
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (board[i][j] == 'Y')
                    board[i][j] = 'O';
                else
                    board[i][j] = 'X';
            }
        }
    }
    
    void dfs(vector<vector<char>>& board, int i, int j, vector<vector<bool>>& visit)
    {   
        visit[i][j] = true;
        board[i][j] = 'Y';
        // if (board[i][j] == 'Y' || board[i][j] == 'X') return;
        
        int m = board.size(), n = board[0].size();
        
        for (int k = 0; k < 4; k++)
        {
            int ni = i + dx[k], nj = j + dy[k];
            if (ni < 0 || ni > m-1 || nj < 0 || nj > n-1 || visit[ni][nj]) continue;
            if (board[ni][nj] == 'O')
                dfs(board, ni, nj, visit);
        }
    }
    
private:
    vector<int> dx = {1, -1, 0, 0};
    vector<int> dy = {0, 0, 1, -1};
};