static const int DICT_SIZE = 26;

struct TrieNode {
public:
    TrieNode(char ch = '\0', struct TrieNode *parent = NULL) {
        this->is_word = false;
        this->ch = ch;
        this->parent = parent;
        this->children.resize(DICT_SIZE, NULL);
    }
    
    ~TrieNode() {
        int len = this->children.size();
        int i;
        for (i = 0; i < len; i++) {
            if (this->children[i] != NULL) {
                delete this->children[i];
                this->children[i] = NULL;
            }
        }
        this->children.clear();
    }
    
public:    
    char ch;
    bool is_word;
    vector<TrieNode*> children;
    TrieNode* parent;
};

class Trie {
public:
    /** Initialize your data structure here. */
    Trie() {
        root = new TrieNode();
    }
    
    /** Inserts a word into the trie. */
    void insert(string word) {
        int lw = word.size();
        int i, idx;
        TrieNode *p = root;
        for (i = 0; i < lw; i++) {
            idx = word[i] - 'a';
            if (p->children[idx] == NULL)
                p->children[idx] = new TrieNode(word[i], p);
            p = p->children[idx];
        }
        p->is_word = true;
    }
    
    /** Returns if the word is in the trie. */
    bool search(string word) {
        int lw = word.size();
        int i, idx;
        TrieNode *p = root;
        for (i = 0; i < lw; i++) {
            idx = word[i] - 'a';
            p = p->children[idx];
            if (p == NULL)
                return false;
        }
        return(p->is_word);
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        int lw = prefix.size();
        int i, idx;
        TrieNode *p = root;
        for (i = 0; i < lw; i++) {
            idx = prefix[i] - 'a';
            p = p->children[idx];
            if (p == NULL)
                return false;
        }
        return(p != NULL);
    }

private:
    TrieNode *root;
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * bool param_2 = obj.search(word);
 * bool param_3 = obj.startsWith(prefix);
 */