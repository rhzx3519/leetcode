#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def canTransform(self, start, end):
        """
        :type start: str
        :type end: str
        :rtype: bool
        """
        i = j = 0
        n = len(start)
        while i<n and j <n:
            # 跳过前面的 X
            while i<n and start[i]=='X':
                i += 1
            while j<n and end[j]=='X':
                j += 1

            if i==n and j==n:
                return True                

            if i>=n or j>=n:
                return False

            if start[i]!=end[j]:
                return False

            # 第三步：检测start中'L'左边'X'的个数是否不少于end中对应位置的'L'左边的'X'个数
            # 第四步：检测start中'R'左边'X'的个数是否不多余end中对应位置的'R'左边的'X'个数                
            if (start[i]=='L' and i<j) or (start[i]=='R' and i>j):
                return False

            i += 1
            j += 1

        return True



if __name__ == '__main__':
    su = Solution()
    start = "RXXLRXRXL"
    end = "XRLXXRRXR"
    print su.canTransform(start, end)


# 思路分析： 题目的意思是说 ‘R’只能向右移动，并且只能移向’X’，‘L’只能向左移动，并且只能移向’X’。

# 第一：如果将start、end中的‘X’全部去掉得到的newStart 和 newEnd相等才有可能转换成功。
# 第二：如果start中'R'的左边'X'在en的个数超过d中对应位置的'R'的左边'X'的个数，则不能转换成功，因为start中的'R'只能向右移动，右边的'X'只能增加不能减少
# 第三：如果end中'L'的左边'X'的个数超过在start中对应位置的'L'的左边'X'的个数，则不能转换成功，因为start中的'L'只能向左移动，左边的'X'只能减少不能增加

# 在一个由 'L' , 'R' 和 'X' 三个字符组成的字符串（例如"RXXLRXRXL"）中进行移动操作。一次移动操作指用一个"LX"替换一个"XL"，或者用一个"XR"替换一个"RX"。
# 现给定起始字符串start和结束字符串end，请编写代码，当且仅当存在一系列移动操作使得start可以转换成end时， 返回True。
# XL -> LX
# RX -> XR
# 示例 :

# 输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
# 输出: True
# 解释:
# 我们可以通过以下几步将start转换成end:
# RXXLRXRXL ->
# XRXLRXRXL ->
# XRLXRXRXL ->
# XRLXXRRXL ->
# XRLXXRRLX

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/swap-adjacent-in-lr-string
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。