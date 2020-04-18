class Solution {
public:
    bool isHappy(int n) {
        unordered_map<int, int> um;
        int t;
        while (n != 1 && um.find(n) == um.end()) {
            um[n]++;
            t = 0;
            while (n != 0) {
                t += (n%10)*(n%10);
                n /= 10;
            }
            
            n = t;
        }
        
        return(n==1);
    }
};