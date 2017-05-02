class BIT {
public:
    BIT(int _n = 0):  n(_n) {
        c.resize(n + 1, 0);
    }
    
    void add(int idx, int val) {
        if (idx < 1 || idx > n) {
            return;
        }
        while (idx <= n) {
            c[idx] += val;
            idx += lowbit(idx);
        }
    }
    
    int sum(int idx) {
        idx = max(idx, 0);
        idx = min(idx, n);
        int s = 0;
        while (idx > 0) {
            s += c[idx];
            idx -= lowbit(idx);
        }
        return s;
    }
    
    ~BIT() {
        c.clear();
    }

private:
    int n;
    vector<int> c;
    
    int lowbit(int x) {
        return x & -x;
    }    
};

class Solution {
public:
    vector<int> countSmaller(vector<int>& nums) {
        int n = nums.size();
        vector<int> res;
        if (n == 0) return res;
        
        set<int> st;
        for (auto &k : nums)
            st.insert(k);
        
        unordered_map<int, int> um;
        int i = 0;
        for (auto it = st.begin(); it != st.end(); it++)
            um[*it] = ++i;
        
        BIT tree(i);
        for (i = n - 1; i >= 0; i--) {
            res.push_back(tree.sum(um[nums[i]] - 1));
            tree.add(um[nums[i]], 1);
        }
        
        reverse(res.begin(), res.end());
        return res;
    }
};