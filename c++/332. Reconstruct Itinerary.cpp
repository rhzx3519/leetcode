class Solution {
public:
    vector<string> findItinerary(vector<pair<string, string>> tickets) {
        int ne = tickets.size();
        for (auto &t : tickets) {
            graph[t.first][t.second]++;
        }
        
        vector<string> res;
        string s = "JFK";
        path.push_back(s);
        bool found = dfs(s, ne, res);
        graph.clear();
        path.clear();
        return res;
    }
    
private:
    map<string, map<string, int>> graph;
    vector<string> path;
    
    bool dfs(string s, int ec, vector<string> &res) {
        if (ec == 0) {
            res = path;
            return true;
        }
        
        string s1;
        for (auto it = graph[s].begin(); it != graph[s].end(); it++) {
            if (it->second == 0)
                continue;
            s1 = it->first;
            --it->second;
            path.push_back(s1);
            if (dfs(s1, ec-1, res))
                return true;
            ++it->second;
            path.pop_back();
        }
        
        return false;
    }
};