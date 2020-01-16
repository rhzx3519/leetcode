class Solution {
public:
    string reverseVowels(string s) {
        int i = 0, j = s.size() - 1;
        while (i <= j) {
            while (i <= j && !is_vowel(s[i])) i++;
            while (i <= j && !is_vowel(s[j])) j--;
            if (i <= j) {
                swap(s[i], s[j]);
                i++;
                j--;
            }
        }
        return s;
    }

private:
    bool is_vowel(char ch) {
        switch(ch) {
            case 'A':
            case 'E':
            case 'I':
            case 'O':
            case 'U':
            case 'a':
            case 'e':
            case 'i':
            case 'o':
            case 'u': return true;
            default: break;
        }
        return false;
    }

};