class Solution(object):
    def maxProfit(self, inventory, orders):
        """
        :type inventory: List[int]
        :type orders: int
        :rtype: int
        """
        Mod = 10**9 + 7
        inventory.sort(reverse=True)
        ans = j = 0
        n = len(inventory)
        while orders > 0:
            while j < n and inventory[j] >= inventory[0]:
                j += 1

            next = 0
            if j != n:
                next = inventory[j]
            buck = j
            delta = inventory[0] - next
            rem = buck*delta
            if rem > orders:
                dec = orders/buck
                a1 = inventory[0] - dec + 1
                an = inventory[0]
                ans += (((a1 + an) * dec) / 2) * bucks;
                ans += (nums[0] - dec) * (orders % bucks);
            else:





if __name__ == '__main__':
    inventory = [3,5]
    orders = 6
    inventory = [2,8,4,10,6]
    orders = 20
    inventory = [1000000000]
    orders = 1000000000
    inventory = [2,5]
    orders = 4
    su = Solution()
    print su.maxProfit(inventory, orders)
