# [57. Insert Interval](https://leetcode.com/problems/insert-interval/)

## 题目
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

给定几个区间，互不相交

如果现在要插入一个新的区间，要求将区间进行合并，依然互不相交

```
Example 1:
Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].

Example 2:
Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].
This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].
```
## 解题思路

1. 插入的new.start在某一区间i内，合并后的merge.start= min(new.start,i.start)
2. 同理，插入后的new.end在某一区间j内，merge.end = max(new.end,j.end)
3. i=j、i<j、i和j只需存在一个，合并的场景都会由上面两种条件覆盖
4. 如果i和j不存在，那么需要将new整个插入到数组中，可以将数组分三段
   + 第一段, i.end < new.start, 直接插入 i
   + 第二段, i.end > new.start && j.start < new.end (i<j), 合并区间
   + 第三段, new.end < i.start, 直接插入 i
