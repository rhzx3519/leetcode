import collections

class Solution(object):
    def minMutation(self, start, end, bank):
        """
        :type start: str
        :type end: str
        :type bank: List[str]
        :rtype: int
        """
        bank = set(bank)
        if end not in bank: return -1
        start in bank and bank.remove(start)
        end in bank and bank.remove(end)
        que = []
        que.append(start)
        vis = collections.defaultdict(int)
        step = 0
        seq = 'ACGT'
        while que:
            sz = len(que)
            step += 1
            while sz:
                sz -= 1
                cur = que.pop(0)
                cur = list(cur)
                for i in range(8):
                    ch = cur[i]
                    for j in seq:
                        if j==ch:continue
                        cur[i] = j
                        str_cur = ''.join(cur)
                        if str_cur==end: return step
                        if str_cur not in bank: continue
                        bank.remove(str_cur)
                        que.append(str_cur)
                    cur[i] = ch

        return -1  

if __name__ == '__main__':
    start = "AACCGGTT"
    end = "AAACGGTA"
    bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
    su = Solution()
    print su.minMutation(start, end, bank)