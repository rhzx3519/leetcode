class Solution {
public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string> s(wordList.begin(), wordList.end());

        queue<string> que;
        unordered_map<string, int> count;
        que.push(beginWord);
        count[beginWord] = 1;
        s.erase(beginWord);
        
        while (!que.empty()) {
            string cur = que.front();
            que.pop();
            int n = count[cur];
            
            for (int i = 0; i < cur.size(); i++) {
                string tmp = cur;
                for (char j = 'a'; j <= 'z'; j++) {
                    if (tmp[i] == j)
                        continue;
                    else
                        tmp[i] = j;
                    
                    if (s.find(tmp) != s.end()) {
                        count[tmp] = n + 1;
                        s.erase(tmp);
                        que.push(tmp);
                        
                        if (tmp == endWord)
                            return count[cur] + 1;
                    }

                }
            }
        }
        
        return 0;
    }
    
};