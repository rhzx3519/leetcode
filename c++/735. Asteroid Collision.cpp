class Solution {
public:
    vector<int> asteroidCollision(vector<int>& asteroids) {
        int n = asteroids.size();
        if (n <= 1) return asteroids;
        int i = n-1;
        while (i > 0) {
            int a = asteroids[i], b = asteroids[i-1];
            if (a*b > 0 || (a>0 && b<0)) {
                i--;
            } else {
                if (abs(a) == abs(b)) {
                    asteroids.erase(asteroids.begin() + i);
                    asteroids.erase(asteroids.begin() + i - 1);
                    i -= 2;
                } else {
                    asteroids[i-1] = abs(a) > abs(b) ? a : b;
                    asteroids.erase(asteroids.begin() + i);
                    i--;
                }
            }
        }
        if (asteroids.size() == n) return asteroids;

        return asteroidCollision(asteroids);
    }
};