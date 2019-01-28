# [274. H-Index](https://leetcode.com/problems/h-index/)

## 题目

Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."

一个作者发布的n篇文章，其中有h篇文章的引用不少于h，即该作者的hindex为h

```$xslt
Example:

Input: citations = [3,0,6,1,5]
Output: 3 
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had 
             received 3, 0, 6, 1, 5 citations respectively. 
             Since the researcher has 3 papers with at least 3 citations each and the remaining 
             two with no more than 3 citations each, her h-index is 3.
```

Note: If there are several possible values for h, the maximum one is taken as the h-index.

Credits:Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.

## 解题思路

见程序注释