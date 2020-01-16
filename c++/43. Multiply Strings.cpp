class Solution {
public:
    string multiply(string num1, string num2) {
        int n = num1.size() + num2.size();
        vector<int> c(n, 0);
        for (int i = num1.size() - 1; i >= 0 ; i--)
        {
            for (int j = num2.size() - 1; j >= 0; j--)
            {
                c[i+j+1] += (num1[i] - '0')*(num2[j] - '0');
            }
        }
        
        for (int i = n-1; i > 0; i--)
        {
            c[i-1] += c[i]/10;
            c[i] %= 10;
        }
        
        int i = 0;
        while (c[i] == 0)
            i++;
        
        string res;
        while (i < n)
        {
            res.push_back(c[i++] + '0');
        }
        if (res.empty()) return "0";
        return res;
    }
};