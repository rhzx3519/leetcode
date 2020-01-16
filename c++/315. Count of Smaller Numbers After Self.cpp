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

    vector<int> countSmaller(vector<int>& nums) {
        vector<pair<int, int>> A;
        for (int i = 0; i < nums.size(); i++)
            A.push_back(make_pair(nums[i], i));
        res.resize(nums.size(), 0);
        
        merge_sort(A, 0, A.size() - 1);
        return res;
    }
    
    void merge_sort(vector<pair<int, int>>& A, int l, int r)
    {
        if (l < r)
        {
            int mid = l + (r - l)/2;
            merge_sort(A, l, mid);
            merge_sort(A, mid+1, r);
            merge(A, l, mid, mid+1, r);
        }
    }
    
    void merge(vector<pair<int, int>>& A, int l1, int r1, int l2, int r2)
    {
        vector<pair<int, int>> tmp;
        int i = l1, j = l2;
        while (i <= r1 && j <= r2)
        {
            if (A[i].first <= A[j].first)
            {
                res[A[i].second] += j - l2;
                tmp.push_back(A[i++]);
            }
            else
                tmp.push_back(A[j++]);
        }
        
        while (i <= r1)
        {
            res[A[i].second] += j - l2;
            tmp.push_back(A[i++]);
        }
        while (j <= r2)
            tmp.push_back(A[j++]);
        
        for (i = l1; i <= r2; i++)
            A[i] = tmp[i - l1];
    }
    
private:
    vector<int> res;    
};