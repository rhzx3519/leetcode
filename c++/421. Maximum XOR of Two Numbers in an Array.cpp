class Solution {
public:
    int findMaximumXOR(vector<int>& nums) {
        int max_val = 0;
        int mask = 0;
        for (int i = 31; i >= 0; i--) {
            // 为获取前n位的临时变量
            mask |= (1<<i);
            set<int> s;
            // 将所有数字的前n位放入set中
            for (int n : nums)
                s.insert(mask&n);
            
            // 假定第n位为1,前n-1位max为之前迭代求得
            int t = max_val | (1<<i);
            for (int pre : s) {
                // 查看`b`是否在
                if (s.count(t^pre))
                {
                    // b存在，第n位为1
                    max_val = t;
                    break;
                }
            }
        }

        return max_val;
    }
};

//https://kingsfish.github.io/2017/12/15/Leetcode-421-Maximum-XOR-of-Two-Numbers-in-an-Array/