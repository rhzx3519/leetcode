static const int DICT_SIZE = 26;

struct TrieNode {
public:
    TrieNode(char ch = '\0', TrieNode* parent = NULL): ch (ch), parent(parent){
        child.resize(DICT_SIZE, NULL);
        is_word = false;
    }
    
    ~TrieNode() {
        for (int i = 0; i < child.size(); i++) {
            if (child[i] != NULL) {
                delete child[i];
                child[i] = NULL;
            }
        }
        
        child.clear();
    }

    char ch;
    vector<TrieNode*> child;
    bool is_word;
    TrieNode* parent;
    
};

class Solution {
public:
    Solution() {
        root = new TrieNode();
    }
    
    ~Solution() {
        delete root;
        root = NULL;
    }
    
    void dfs(vector<vector<char>>& board, int i, int j, string &s, vector<string> &res, TrieNode *p) {
        if (p == NULL) return;
        if (p->is_word) {
            res.push_back(s);
            // Next time you won't stop by.
            p->is_word = false;
        }
        
        char ch;
        int m = board.size(), n = board[0].size();
        int x, y, k;
        for (k = 0; k < 4; k++) {
            x = i + offset[k][0];
            y = j + offset[k][1];
            if (x < 0 || x > m-1 || y < 0 || y > n-1)
                continue;
                
            if (visit[x][y])
                continue;
            ch = board[x][y];
            if (p->child[ch - 'a'] == NULL)
                continue;
                
            visit[x][y] = true;
            s.push_back(ch);
            dfs(board, x, y, s, res, p->child[ch - 'a']);
            s.pop_back();
            visit[x][y] = false;
        }
        
    }
    
    vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
        vector<string> res;
        for (auto &s : words)
            insert(s);
            
        int m = board.size(), n = board[0].size();
        visit.resize(m, vector<bool>(n, false));
        string s;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                visit[i][j] = true;
                s.push_back(board[i][j]);
                
                dfs(board, i, j, s, res, root->child[board[i][j] - 'a']);
                
                s.pop_back();
                visit[i][j] = false;
            }
        }
        
        return res;
    }
    
    void insert(const string &word) {
        int lw = word.size();
        int idx;
        TrieNode *p = root;
        for (int i = 0; i < lw; i++) {
            idx = word[i] - 'a';
            if (p->child[idx] == NULL) 
                p->child[idx] = new TrieNode(word[i], p);
            p = p->child[idx];
        }
        p->is_word = true;
    }
    
private:
    TrieNode* root;
    int offset[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    vector<vector<bool>> visit;
    
};

// 思路： 字典树 + 回溯查找，先将words中的单词构造成一颗字典树，深度遍历的时候可以
// 根据字典树来剪枝，每次dfs的深度不会超过字典树的高度
// k 为words的单词数目，D为单词平均长度
// time: O(M*N*logK), space(K*D)




