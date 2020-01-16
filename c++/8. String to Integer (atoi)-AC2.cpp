class Solution {
public:
    int myAtoi(string str) {
        int i = 0, n = str.size();
        while (i < n && str[i] == ' ')
            i++;
        
        if (i >= n) return 0;
        bool sign = true;
        if (str[i] == '+')
        {
            sign =true;
            i++;
        }
        else if (str[i] == '-')
        {
            sign = false;
            i++;
        }
        

        if (i >= n)
            return 0;
        
        long long res = 0;
        int start = i;
        while (i < n && str[i] >= '0' && str[i] <= '9')
        {
            res = res*10 + str[i] - '0';
            if (sign && res > INT_MAX) return INT_MAX;
            if (!sign && -res < INT_MIN) return INT_MIN;
            i++;
        }
        
        return sign ? res : -res;
    }
};