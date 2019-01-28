# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## 题目
Given a collection of intervals, merge all overlapping intervals.

```text
Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

## 解题思路

1. 插入的new.start在某一区间i内，合并后的merge.start= min(new.start,i.start)
2. 同理，插入后的new.end在某一区间j内，merge.end = max(new.end,j.end)
3. i=j、i<j、i和j只需存在一个，合并的场景都会由上面两种条件覆盖
4. 如果i和j不存在，那么需要将new整个插入到数组中，可以将数组分三段
   + 第一段, i.end < new.start, 直接插入 i
   + 第二段, i.end > new.start && j.start < new.end (i<j), 合并区间
   + 第三段, new.end < i.start, 直接插入 i
