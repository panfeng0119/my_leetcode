# [164. Maximum Gap](https://leetcode.com/problems/maximum-gap/)

## 题目

You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.

Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

```
Example 1:

Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.
Example 2:

Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.
```

Note:

+ You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
+ Try to solve it in linear time/space.

## 思路

+ 此题目使用桶排序.
+ 我们首先获取数组的最小值mixnum和最大值maxnum，得到数组的范围.
+ n个数有n-1个间隔，我们计算平均间隔gap：(maxnum-minnum)/n-1 向上取整.
+ 我们计算需要的桶的个数size = int((maxnum-minnum)/gap)+1个此题目的关键思想是：在一个桶内的数字之间的差值一定小于gap，如果某两个数之间的差大于平均差gap，一定会被放到两个桶内。最大的差一定大于等于gap（对一组数求平均值，平均值小于等于最大值），于是如果出现了两个数a，b，且a和b的差大于gap，那么a和b一定会被放到两个连续的桶t1，t2内，且a是t1桶的嘴后一个值（最大值），b是t2桶的第一个值（最小值）。于是我们只需要记录每个桶内的最大值和最小值，让后用当前桶内的最小值减去上一个桶内的最大值得到maxgap，取maxgap最大的一个返回即可.
+ 要注意的是，如果在计算平均距离gap时候如果得到了0，说明所有的数相等，这时可以直接返回0.

