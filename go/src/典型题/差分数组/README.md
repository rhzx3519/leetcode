所谓差分数组即 diff[i] = nums[i] - nums[i-1]    
当一段区间的所有元素都+val时，例如nums[i]~nums[j]之间的元素都加上了val,   
diff[i]-diff[j]保持不变, diff[l] += val, diff[r+1] -= val,  
可以表示 [l, r]区间的元素都加上了val

重建时： cnt[0] =  diff[0],
for i in range(1, n):
    cnt[i] = cnt[i-1] + diff[i]