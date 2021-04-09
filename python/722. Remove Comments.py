class Solution(object):
    def removeComments(self, source):
        """
        :type source: List[str]
        :rtype: List[str]
        """
        isBlock = False
        ans = []
        newLine = []
        for line in source:
            if not isBlock:
                newLine = []
            i = 0
            while i < len(line):
                if line[i:i+2]=='/*' and not isBlock:
                    isBlock = True
                    i += 1
                elif line[i:i+2]=='*/' and isBlock:
                    isBlock = False
                    i += 1
                elif line[i:i+2]=='//' and not isBlock:
                    break
                elif not isBlock:
                    newLine.append(line[i])
                i += 1

            if not isBlock and newLine:
                ans.append(''.join(newLine))
        return ans

            
if __name__ == '__main__':
    source = ["a/*comment", "line", "more_comment*/b"]
    # source = ["/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"]
    # source = ["void func(int k) {", "// this function does nothing /*", "   k = k*2/4;", "   k = k/2;*/", "}"]
    su = Solution()
    su.removeComments(source)