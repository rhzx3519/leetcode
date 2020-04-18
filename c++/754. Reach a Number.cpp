class Solution {
public:
    int reachNumber(int target) {
        target = abs(target);
        int i = 1;
        int sum = 0;
        while(true) {
            sum += i;
            if (sum >= target && (sum-target)%2==0)
                return i;
            i += 1;
        }
    }


};

// 1、由于对称性 target是正或负不影响，都取正。
// 2、从0-target尽量都向右（都取正步数）才能得最优解。
// 3、步数都取正，直到sum>=target
// 4、当sum-target为偶数时，在前面的（sum-target）/2取负步数即可。例sum-target=2，则+1变为-1，即向左了2步。
// 5、向左时无论当前步数是奇是偶，相当于乘2，都会变成偶数。所以只有当(sum - target)为偶数时，才能通过向左移动使得最终到达target。