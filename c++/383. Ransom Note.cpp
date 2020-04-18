class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        int letters1[26] = {0}, letters2[26] = {0};
        for (auto &ch : ransomNote)
            letters1[ch-'a']++;
        for (auto &ch : magazine)
            letters2[ch-'a']++;
        
        for (int i = 0; i < 26; i++) {
            if (letters2[i] < letters1[i])
                return false;
        }
        return true;
    }
};