class Solution {
public:
    bool isAnagram(string s, string t) {
        int a[26] = {0};
        for (auto &ch : s)
            a[ch - 'a']++;
        for (auto &ch : t)
            a[ch - 'a']--;
            
        for (int i = 0; i < 26; i++) 
            if (a[i] != 0)
                return false;
        
        return true;
    }
};