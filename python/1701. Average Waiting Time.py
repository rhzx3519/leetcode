class Solution(object):
    def averageWaitingTime(self, customers):
        """
        :type customers: List[List[int]]
        :rtype: float
        """
        wait = 0
        que = customers
        n = len(customers)
        t = -1
        while que:
            arrival, time = que.pop(0)
            if t < arrival:
                t = arrival + time
                wait += time
            else:
                t += time
                wait += (t - arrival)
            # print t, wait
        
        return float(wait)/n