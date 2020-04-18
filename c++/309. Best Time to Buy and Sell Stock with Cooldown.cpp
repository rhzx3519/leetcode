// 状态的跳转依旧是时间之间的跳转，即第i天的收益情况依赖于第i-1天的收益情况。
// 不过现在需要三个状态，即buy，sell，cooldown，我们记录第i-1天的这三个状态的收益情况是last_buy,last_sell,last_cooldown。
// 那么第i天的这三个收益情况的依赖关系是： 
// buy=max(last_buy, last_cooldown - price[i]) 
// sell = max(last_sell, last_buy + price[i]) 
// cooldown = max(cooldown, last_sell); 
// 然后就是每天计算完之后将该天的这三个状态值赋值last。 

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int n = prices.size();
        if (n==0) return 0;
        int last_buy = -prices[0];
        int last_sell = INT_MIN;
        int last_rest = 0;

        int buy, sell, rest;
        for (int i = 1; i < prices.size(); i++) {
            int rest = max(last_rest, last_sell);
            int buy = max(last_rest-prices[i], last_buy);
            int sell = last_buy + prices[i];

            last_rest = rest;
            last_buy = buy;
            last_sell = sell;
        }
        
        return max(last_sell, last_rest);
    }
};


