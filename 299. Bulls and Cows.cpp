class Solution {
public:
    string getHint(string secret, string guess) {
        int n = secret.size();
        int visit[n] = {0};
        int bull = 0, cow = 0;
        vector<vector<int>> vv(10, vector<int>());
        for (int i = 0; i < n; i++) {
            if (secret[i] == guess[i]) {
                bull++;
                visit[i] = 1;
            } else 
                vv[secret[i] - '0'].push_back(i);
        }
        
        for (int i = 0; i < n; i++) {
            if (visit[i] == 0) {
                char ch = guess[i];
                if (!vv[ch - '0'].empty()) {
                    cow++;
                    vv[ch - '0'].erase(vv[ch - '0'].begin());
                }
            }
        }
        
        char buf[10];
        sprintf(buf, "%dA%dB", bull, cow);
        return string(buf);
    }
};