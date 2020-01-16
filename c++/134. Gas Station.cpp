class Solution {
public:
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int total_gas = 0, total_cost = 0, cur = 0, start = 0;
        for (int i = 0; i < gas.size(); i++) {
            cur += gas[i];
            cur -= cost[i];
            total_gas += gas[i];
            total_cost += cost[i];
            if (cur < 0) {
                start = i + 1;
                cur = 0;
            }
        }
        
        if (total_gas >= total_cost)
            return start;
        return -1;
    }
};