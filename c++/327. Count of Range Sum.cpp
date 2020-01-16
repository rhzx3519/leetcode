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
    int countRangeSum(vector<int>& nums, int lower, int upper) {
        int n = nums.size();
        int64_t s = 0;
        map<int64_t, int> mm;
        mm[lower] = 0; mm[upper+1] = 0;
        
        int i, res = 0;
        for (i = 0; i < n; i++) {
            s += nums[i];
            mm[s] = 0;
            mm[s + lower] = 0;
            mm[s + upper + 1] = 0;
        }
        i = 0;
        for (auto it = mm.begin(); it != mm.end(); it++)
            it->second = ++i;
        
        BIT bit(i);
        
        bit.add(mm[lower], 1);
        bit.add(mm[upper+1], -1);
        s = 0;
        for (i = 0; i < n; i++) {
            s += nums[i];
            res += bit.sum(mm[s]);
            bit.add(mm[s + lower], 1);
            bit.add(mm[s + upper + 1], -1);
        }
        
        mm.clear();
        return res;
    }
};