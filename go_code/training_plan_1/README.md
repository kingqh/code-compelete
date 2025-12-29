
## LeetCode 刷题计划（Go，14天，每天 2 题）
### 使用说明
- 每天完成两题：第一题快速提交，第二题深入复盘（多解法、复杂度分析、边界测试）。
- 推荐每天耗时：1.5–2 小时（快速题 30–45 分钟，复盘题 45–60 分钟）。
- Go 代码模板：
  - 链表：`type ListNode struct { Val int; Next *ListNode }`
    - 树：`type TreeNode struct { Val int; Left, Right *TreeNode }`
    - 提交到 LeetCode 时按题目要求实现对应函数签名；仓库中可保留 `main` 测试代码。
    ### 目标与方法
    - 覆盖数组、字符串、链表、树、哈希、双指针、二分、动态规划、贪心、回溯与图等常见考点。
    - 每题复盘内容包括：多种解法、复杂度分析、边界用例与易错点。
    - 建议每题写 5–8 个测试用例并在本地用 `go test` 或简单 `main` 调试。
    ### 目录（按天）
    #### 第 1 天
    - 题目 A: LeetCode 1 — Two Sum
    考点：哈希表、数组
    做法：一次遍历，用 `map` 存 value->index（O(n) 时间，O(n) 空间）
    - 题目 B: LeetCode 206 — Reverse Linked List
    考点：链表指针操作
    做法：迭代三指针（prev, cur, next）；复盘递归实现
    #### 第 2 天
    - 题目 A: LeetCode 3 — Longest Substring Without Repeating Characters
    考点：滑动窗口、哈希
    - 题目 B: LeetCode 21 — Merge Two Sorted Lists
    考点：链表合并、哨兵节点
    #### 第 3 天
    - 题目 A: LeetCode 11 — Container With Most Water
    考点：双指针、贪心
    - 题目 B: LeetCode 15 — 3Sum
    考点：排序 + 双指针去重
    #### 第 4 天
    - 题目 A: LeetCode 20 — Valid Parentheses
    考点：栈、字符串
    - 题目 B: LeetCode 35 — Search Insert Position
    考点：二分查找模板
    #### 第 5 天
    - 题目 A: LeetCode 53 — Maximum Subarray
    考点：动态规划 / Kadane
    - 题目 B: LeetCode 94 — Binary Tree Inorder Traversal
    考点：树遍历（递归 & 迭代）
    #### 第 6 天
    - 题目 A: LeetCode 101 — Symmetric Tree
    考点：树、镜像比较（递归/队列）
    - 题目 B: LeetCode 33 — Search in Rotated Sorted Array
    考点：二分变体
    #### 第 7 天
    - 题目 A: LeetCode 153 — Find Minimum in Rotated Sorted Array
    考点：二分、边界处理
    - 题目 B: LeetCode 36 — Valid Sudoku
    考点：哈希 / 布尔矩阵
    #### 第 8 天
    - 题目 A: LeetCode 238 — Product of Array Except Self
    考点：前缀后缀乘积，O(n) 无除法
    - 题目 B: LeetCode 347 — Top K Frequent Elements
    考点：哈希 + 堆 / 桶排序
    #### 第 9 天
    - 题目 A: LeetCode 56 — Merge Intervals
    考点：区间合并、排序扫描
    - 题目 B: LeetCode 57 — Insert Interval
    考点：区间插入、边界判断
    #### 第 10 天
    - 题目 A: LeetCode 200 — Number of Islands
    考点：DFS/BFS/并查集、网格遍历
    - 题目 B: LeetCode 207 — Course Schedule
    考点：图、拓扑排序 / DFS 判环
    #### 第 11 天
    - 题目 A: LeetCode 70 — Climbing Stairs
    考点：动态规划、斐波那契
    - 题目 B: LeetCode 5 — Longest Palindromic Substring
    考点：中心扩展 / Manacher
    #### 第 12 天
    - 题目 A: LeetCode 146 — LRU Cache
    考点：哈希 + 双向链表（O(1) get/put）
    - 题目 B: LeetCode 297 — Serialize and Deserialize Binary Tree
    考点：树序列化与反序列化
    #### 第 13 天
    - 题目 A: LeetCode 139 — Word Break
    考点：动态规划、记忆化搜索
    - 题目 B: LeetCode 39 — Combination Sum
    考点：回溯、剪枝
    #### 第 14 天
    - 题目 A: LeetCode 560 — Subarray Sum Equals K
    考点：前缀和 + 哈希表
    - 题目 B: LeetCode 239 — Sliding Window Maximum
    考点：单调队列（deque），O(n)
    ### 每题复盘要点（通用）
    - 明确输入输出与边界条件（空输入、极端值）。
    - 写出暴力/直观解法并分析复杂度，再思考能否优化到更优解法。
    - 列出并验证至少 5 个测试用例（含边界和随机样例）。
    - 注意 Go 语言细节：切片索引越界、nil 指针、map 零值等。
    - 对数据结构题先手工模拟 1–2 个小例子再实现。
