#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 所谓差分数组即 diff[i] = nums[i] - nums[i-1]
# 当一段区间的所有元素都+val时，例如nums[i]~nums[j]之间的元素都加上了val,
# diff[i]-diff[j]保持不变, 
# diff[l] += val, diff[r+1] -= val, 可以表示 [l, r]区间的元素都加上了val
# 重建时： cnt[0] =  diff[0],
# for i in range(1, n):
#     cnt[i] = cnt[i-1] + diff[i]

n = 11
gap = [[1, 3], [2, 5], [9, 10]] # 表示[i, j]之间的区间+1
    
diff = [0] *(n+1)
for l, r in gap:
    diff[l] += 1
    diff[r+1] -= 1

nums = [0]*n
nums[0] = diff[0]
for i in range(1, n):
    nums[i] = nums[i-1] + diff[i]
print diff
print nums