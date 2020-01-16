class Solution {
public:
    int minMutation(string start, string end, vector<string>& bank) {
        unordered_set<string> vis;
        unordered_set<string> ba(bank.begin(), bank.end());
        int min_val = INT_MAX;
        dfs(start, end, 0, min_val, ba, vis);
        return min_val==INT_MAX ? -1 : min_val;

    }

    void dfs(string start, string end, int depth, int &min_val, unordered_set<string> &ba, unordered_set<string> &vis) {
        if (start == end) {
            min_val = min(min_val, depth);
            return;
        }

        for (int i = 0; i < start.size(); i++) {
            for (int k = 0; k < letter.size(); k++) {
                string t = start;
                if (letter[k] == start[i])
                    continue;
                t[i] = letter[k];
                if (vis.count(t) || ba.count(t)==0)
                    continue;
                vis.insert(t);
                dfs(t, end, depth+1, min_val, ba, vis);
                vis.erase(t);
            }
        }
    }

private:
    string letter = "ACGT";
};