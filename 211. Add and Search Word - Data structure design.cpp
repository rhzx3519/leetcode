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
    bool search(TrieNode *root, string word, int idx) {
        int lw = word.size();
        int t;
        while (idx < lw && word[idx] != '.') {
            t = word[idx] - 'a';
            root = root->children[t];
            if (root == NULL)
                return false;
            idx++;
        }
        
        if (idx == lw)
            return root->is_word;
            
        for (int i = 0; i < root->children.size(); i++) {
            if (root->children[i] != NULL && search(root->children[i], word, idx + 1))
                return true;
        }
        
        return false;
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

public:
    TrieNode *root;
};


class WordDictionary {
public:
    /** Initialize your data structure here. */
    WordDictionary() {
        trie = new Trie();
    }
    
    ~WordDictionary() {
        delete trie;
        trie = NULL;
    }
    
    /** Adds a word into the data structure. */
    void addWord(string word) {
        trie->insert(word);
    }
    
    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    bool search(string word) {
        return trie->search(trie->root, word, 0);
    }
    
private:
    Trie *trie;
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary obj = new WordDictionary();
 * obj.addWord(word);
 * bool param_2 = obj.search(word);
 */