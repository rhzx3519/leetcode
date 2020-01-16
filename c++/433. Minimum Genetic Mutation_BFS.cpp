class Solution {
public:
    int minMutation(string start, string end, vector<string>& bank) {
        unordered_set<string> bk(bank.begin(), bank.end());
        unordered_set<string> vis;

        queue<string> que;
        int res = 0;
        que.push(start);
        vis.insert(start);

        while (!que.empty()) {

            int sz = que.size();
            while (sz-->0) {
                string s = que.front();
                que.pop();
                if (s==end)
                    return res;
                    
                for (int i = 0; i < s.size(); i++) {
                    for (int k = 0; k < letter.size(); k++) {
                        string t = s;
                        t[i] = letter[k];
                        if (vis.count(t)==0 && bk.count(t)!=0) {
                            vis.insert(t);
                            que.push(t);
                        }
                            
                    }
                }

            }
            res++;
        }

        return -1;
    }

private:
    string letter = "ACGT";
};