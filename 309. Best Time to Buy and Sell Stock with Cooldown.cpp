状态的跳转依旧是时间之间的跳转，即第i天的收益情况依赖于第i-1天的收益情况。不过现在需要三个状态，即buy，sell，cooldown，我们记录第i-1天的这三个状态的收益情况是last_ buy,last_sell,last_cooldown。那么第i天的这三个收益情况的依赖关系是： 
buy=max(last_buy, last_cooldown - price[i]) 
sell = max(last_sell, last_buy + price[i]) 
cooldown = max(cooldown, last_sell); 
然后就是每天计算完之后将该天的这三个状态值赋值last。 