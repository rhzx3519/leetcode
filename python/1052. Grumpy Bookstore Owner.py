class Solution(object):
    def maxSatisfied(self, customers, grumpy, X):
        """
        :type customers: List[int]
        :type grumpy: List[int]
        :type X: int
        :rtype: int
        """
        ans = max_val = 0
        n = len(customers)
        s = sum([customers[i] for i in range(n) if grumpy[i]==0])        

        for i in range(X):
            max_val += customers[i]
            if grumpy[i]==0:
                s -= customers[i]


        for i in range(X, n):
            ans =  max(ans, max_val + s)
            max_val += customers[i]
            max_val -= customers[i-X]

            if grumpy[i]==0:
                s -= customers[i]
            if grumpy[i-X]==0:
                s += customers[i-X]
        
        ans = max(ans, max_val + s)

        print 'ans:', ans



if __name__ == '__main__':
    customers = [1,0,1,2,1,1,7,5]
    grumpy = [0,1,0,1,0,1,0,1]
    X = 3

    # customers = [4,10,10]
    # grumpy = [1,1,0]
    # X = 2

    # customers = [10,1,7]
    # grumpy = [0,0,0]
    # X = 2

    customers = [2,6,6,9]
    grumpy = [0,0,1,1]
    X = 1

    su = Solution()
    su.maxSatisfied(customers, grumpy, X)

