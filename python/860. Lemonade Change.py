class Solution(object):
    def lemonadeChange(self, bills):
        """
        :type bills: List[int]
        :rtype: bool
        """
        count5 = count10 = 0
        for bill in bills:
            if count5 < 0: break
            if bill==5:
                count5 += 1
            elif bill==10:
                count10 += 1
                count5 -= 1
            elif bill==20:
                if count10:
                    count10 -= 1
                    count5 -= 1
                else:
                    count5 -= 3
        
        return count5 >= 0

        