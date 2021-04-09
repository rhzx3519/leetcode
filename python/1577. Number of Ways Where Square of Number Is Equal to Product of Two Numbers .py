class Solution(object):
    def numTriplets(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: int
        """
        def frac(n):
            if n <= 1:
                return n
            return n + frac(n-1)

        def statistic(nums1, nums2):
            ans = 0
            count1 = collections.Counter(nums1)
            count = collections.Counter(nums2)
            n1, n2 = len(nums1), len(nums2)
            for i_val, i_cnt in count1.iteritems():
                i2 = i_val * i_val
                vis = set()
                for j_val, cnt in count.iteritems():
                    if i2%j_val != 0:
                        continue
                    k_val = i2/j_val
                    if k_val not in count:
                        continue
                    if k_val in vis:
                        continue
                    if k_val == j_val:
                        ans += frac(cnt-1) * i_cnt
                    else:
                        # print j_val, k_val
                        ans += cnt * count[k_val] * i_cnt
                    vis.add(j_val)
                    vis.add(k_val)
                    # del count[j_val]

            # print '-'*20
            return ans

        return statistic(nums1, nums2) + statistic(nums2, nums1)
        