#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def openLock(self, deadends, target):
        """
        :type deadends: List[str]
        :type target: str
        :rtype: int
        """
        deadends = set(deadends)
        que = []
        que.append('0000')
        level = 0
        while que:
            sz = len(que)
            
            # print que, sz
            while sz > 0:
                sz -= 1
                s = que.pop(0)
                if s in deadends:
                    continue
                if s == target:
                    return level

                deadends.add(s)
                candidates = []
                for i in range(4):
                    tmp = list(s)
                    tmp[i] = str((int(s[i]) + 1) % 10)
                    candidates.append(''.join(tmp))
                    tmp[i] = str((int(s[i]) - 1) % 10)
                    candidates.append(''.join(tmp))

                for cand in candidates:
                    if cand in deadends:
                        continue
                    que.append(cand)

            level += 1    

        return -1


## 广度优先遍历，注意将校验过的字符添加到deadends避免重复校验   



if __name__ == '__main__':
    deadends = ["0201","0101","0102","1212","2002"]
    target = "0202"   

    deadends = ["8888"]
    target = "0009"   

    deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"]
    target = "8888"

    su = Solution()
    print su.openLock(deadends, target)


# 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

# 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

# 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

# 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

'''
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。

示例 3:

输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
'''