# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

## 题目

Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array `[-2,1,-3,4,-1,2,1,-5,4]`,

the contiguous subarray `[4,-1,2,1]` has the largest sum = `6`.

More practice:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## 解题思路

利用好`连续`这个约束

### 用动态规划

已知数组的前k个元素的和最大子的序列的和为maxSub，以及前k个数组的临时和sum。

对于第k+1个元素

如果sum<0, 对k+1的求和是没有贡献，不如直接从k+1来开始算，此时应sum置0

注意，只要sum不减到负数，中间出现小于0的元素是没关系的，sum仍然可以继续累加。

    1）若sum[k]>=0,sum[k+1]=sum[k]+A[k+1]
    2）若sum[k]<0,sum[k+1]=0+A[k+1]

那么

    sum[i] = max(sum[i-1]+array[i], array[i])
    res = max(res，sum[i])

### 贪心算法

其从头开始一直保持将数组加入进来，直到sum < 0，则将sum置为0，ans一直保存最大的和。（和dp差不多）