class Solution(object):
    def prisonAfterNDays(self, cells, N):
        """
        :type cells: List[int]
        :type N: int
        :rtype: List[int]
        """
        M = N
        mem = [cells[:]]
        ln = len(cells)
        while N:
            N -= 1

            
            for i in range(1, ln-1):
                t = ((cells[i-1]&1) ^ (cells[i+1]&1)) ^ 1
                # print i, cells[i-1]&1, cells[i+1]&1, t
                if t:
                    cells[i] |= (1<<1)
                else:
                    cells[i] &= ~(1<<1)
            # print map(bin, cells)
            for i in range(1, ln-1):
                cells[i] >>= 1

            if cells[0]==1:
                cells[0] = 0
            if cells[-1]==1:
                cells[-1] = 0

            # print cells, mem[0], cells==mem[0]
            if cells in mem:
                i = mem.index(cells)
                return mem[i + N%(len(mem) - i)]

            mem.append(cells[:])
        
        # print cells
        return cells

if __name__ == '__main__':
    cells = [1,0,0,1,0,0,1,0]
    N = 36
    su = Solution()
    print su.prisonAfterNDays(cells, N)