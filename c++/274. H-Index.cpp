class Solution {
public:
    int hIndex(vector<int>& citations) {
        vector<int>& a = citations;
        sort(a.begin(), a.end());
        int n = a.size();
        for (int i = 0; i < n; i++) {
            if (a[i] >= n-i)
                return n-i;
        }
        return 0;
    }
};