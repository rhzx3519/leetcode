class Solution(object):
    def __init__(self):
        self.k = 1
        self.files = set()

    def getFolderNames(self, names):
        """
        :type names: List[str]
        :rtype: List[str]
        """
        ans = []
        for name in names:
            if name not in self.files:
                self.files.add(name)
                ans.append(name)
            else:
                tmp = "%s(%d)" % (name, self.k)
                self.k += 1
                self.files.add(tmp)
                ans.append(tmp)
        return ans
                    

if __name__ == '__main__':
    names = ["gta","gta(1)","gta","avalon"]
    su = Solution()
    print su.getFolderNames(names)