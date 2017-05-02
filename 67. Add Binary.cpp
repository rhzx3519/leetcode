class Solution {
public:
    string addBinary(string a, string b) {
        int m = a.size(), n = b.size();
        reverse(a.begin(), a.end());
        reverse(b.begin(), b.end());
        int add = 0, i = 0, j = 0, k = 0, t;
        vector<int> c(m+n+1, 0);
        
        while (i < m && j < n) {
            t = ((a[i] - '0') + (b[j] - '0') + add) % 2;
            add = ((a[i] - '0') + (b[j] - '0') + add) / 2;
            c[k++] = t;
            i++;j++;
        }
        
        while (i < m) {
            t = ((a[i] - '0') + add) % 2;
            add = ((a[i] - '0') + add) / 2;
            c[k++] = t;
            i++;
        }
        
        while (j < n) {
            t = ((b[j] - '0') + add) % 2;
            add = ((b[j] - '0') + add) / 2;
            c[k++] = t;
            j++;
        }
        
        if (add != 0)
            c[k++] = add;
        string res;
        for (i = k-1; i >= 0; i--)
            res += '0' + c[i];
        
        return res;
    }
};