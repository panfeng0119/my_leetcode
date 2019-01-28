# [75. Sort Colors](https://leetcode.com/problems/sort-colors/)

## 题目

关键字：编程技巧、空间利用、双指针

难度：中等

前置知识：二维数组、矩阵

题目大意：对于给定的值域处于0~2的数组，对其排序，要求仅扫描一遍数组并且只使用常数空间。

要求：不能使用sort包

Example:

```text
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

Follow up:

- A rather straight forward solution is a two-pass algorithm using counting sort.
- First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
- Could you come up with a one-pass algorithm using only constant space?

