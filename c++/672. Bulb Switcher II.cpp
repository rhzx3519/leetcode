/*
所有的操作进行偶数次是会抵消的，那么所有操作只存在0,1两种，即无效果和有效果；
且各一次操作2,3等效为一次操作1；
画一个类似真值表，可以推出  n>=3且m>=3时,结果只会是8；
接下来考虑个别情况即可。
*/

class Solution {
public:
    int flipLights(int n, int m) {
        if (n==0)
            return 0;
        if (m==0)
            return 1;            
        if (n==1)
            return 2;
        if (n==2) {
            if (m==1) return 3;
            else return 4;
        }
        if (m==1)
            return 4;
        if (m==2) 
            return 7;

        return 8;
    }
};