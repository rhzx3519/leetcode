class Solution {
public:
    string originalDigits(string s) {
        int a[10] = {0};
        int letters[26] = {0};
        for (char c : s) {
            letters[c-'a']++;
            if (c=='z')
                a[0]++;
            else if (c=='w')
                a[2]++;
            else if (c=='u')
                a[4]++;
            else if (c=='x')
                a[6]++;
            else if (c=='g')
                a[8]++;
        }

        for (int i = 0; i < 10; i++) {
            string num = nums[a[i]];
            if (a[i] != 0) {
                for (char c : num) {
                    letters[c-'a'] -= a[i]
                }
            }
        }

        a[1] = letters['o'-'a'];
        a[3] = letters['t' - 'a'];
        a[5] = letters['f' - 'a'];
        a[7] = letters['s' - 'a'];
        for (char c : nums[1])
            letters[c-'a'] -= a[1];
        for (char c : nums[3])
            letters[c-'a'] -= a[3];
        for (char c : nums[5])
            letters[c-'a'] -= a[5];
        for (char c : nums[7])
            letters[c-'a'] -= a[7];                                
        a[9] = letters['e'-'a'];

        string res;
        for (int i = 0; i <10 ;i++)
            res.append(string(a[i], '0'+i));
        return res;
    }

private:
    string nums[] = {"zero", "two", "three", "four", "five",
                    "six", "seven", "eight", "nine"};
};