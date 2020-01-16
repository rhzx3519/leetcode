typedef long long LL;

class Solution {
public:
    int countDigitOne(int n) {
        LL res = 0;
        if (n <= 0) return res;
        
        LL f = 1;
        LL h = 0, c = 0, l = 0;
        while (n/f != 0) {
            l = n - (n/f)*f;
            c = (n/f)%10;
            h = n/(f*10);
            
            switch(c) {
                case 0:
                    res += h * f;
                    break;
                case 1:
                    res += h * f + l + 1;
                    break;
                default:
                    res += (h+1)*f;
                    break;
            }
            
            f *= 10;
        }
        
        return res;
    }
};