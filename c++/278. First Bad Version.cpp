// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int l = 1, r = n;
        if (isBadVersion(l)) return l;
        if (!isBadVersion(r)) return r+1;
        int m;
        while (r - l > 1) {
            m = l+ (r-l)/2;
            if (isBadVersion(m)) 
                r = m;
            else
                l = m;
        }
        
        return r;
    }
};