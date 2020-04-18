class MapSum {
public:
    /** Initialize your data structure here. */
    MapSum() {
        
    }
    
    void insert(string key, int val) {
        int n = key.size();

        if (mp.count(key) != 0) {
            int v = mp[key];
            for (int i = 1; i <= n; i++) {
                string s = key.substr(0, i);
                sum_mp[s] -= v;
            }
        }
        mp[key] = val;
        for (int i = 1; i <= n; i++) {
            string s = key.substr(0, i);
            sum_mp[s] += val;
        }
    }
    
    int sum(string prefix) {
        if (sum_mp.count(prefix) != 0) return sum_mp[prefix];
        return 0;
    }

    unordered_map<string, int> mp;
    unordered_map<string, int> sum_mp;
};

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum* obj = new MapSum();
 * obj->insert(key,val);
 * int param_2 = obj->sum(prefix);
 */

// 注意覆盖的时候，需要更新前缀和数据结构