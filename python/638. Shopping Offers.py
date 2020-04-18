class Solution(object):
    def shoppingOffers(self, price, special, needs):
        """
        :type price: List[int]
        :type special: List[List[int]]
        :type needs: List[int]
        :rtype: int
        """
        def shopping(special, needs):
            if not sum(needs):
                return 0
            # 过滤掉比当前needs多的礼包
            special = list(filter(lambda x: all(x[i] <= needs[i] for i in range(l)), special))
            # 如果没有可选礼包，直接返回以单品购买的价格
            if not special:
                return sum([needs[i]*price[i] for i in range(l)])
            # 回溯, 收集本次购买每种礼包的花费加上若购买该礼包后剩余needs递归的最低花费
            res = []
            for x in special:
                res.append(x[-1] + shopping(special, [needs[i] - x[i] for i in range(l)]))
            return min(res)

        l = len(needs)
        # 先过滤掉比单独购买价格高的礼包
        special = list(filter(lambda x: x[-1] < sum(x[i]*price[i] for i in range(l)), special))
        return shopping(special, needs)

    