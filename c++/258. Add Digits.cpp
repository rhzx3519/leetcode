class Solution {
public:
    int addDigits(int num) {
        return 1 + (num - 1) % 9;
    }

    int addDigits(int num) {
        while (to_string(num).size() != 1)
        {
            string s = to_string(num);
            int t = 0;
            for (int i = 0; i < s.size(); i++)
                t += s[i] - '0';
            num = t;
        }
        
        return num;
    }
};