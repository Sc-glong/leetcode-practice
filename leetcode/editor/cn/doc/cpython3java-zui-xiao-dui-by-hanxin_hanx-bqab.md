思路和心得：


# （一）最小堆

1.==排名第k个选手（数值大小）

2.维护一个长度为k的最小堆。每次输出堆顶
==k个人中最小的那个==排行榜的第k个


```python3 []
class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.minHeap = []
        for x in nums:
            heapq.heappush(self.minHeap, x)

    def add(self, val: int) -> int:
        heapq.heappush(self.minHeap, val)

        while len(self.minHeap) > self.k:
            heapq.heappop(self.minHeap)
            
        return self.minHeap[0] 


# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)
```

```c++ []
class KthLargest 
{
public:
    int k;
    priority_queue<int, vector<int>, greater<int>> minHeap; 

    KthLargest(int k, vector<int>& nums) 
    {
        this->k = k;
        for (int x : nums)
            minHeap.push(x);
    }
    
    int add(int val) 
    {
        minHeap.push(val);

        while (minHeap.size() > k)
            minHeap.pop();

        return minHeap.top();
    }
};

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest* obj = new KthLargest(k, nums);
 * int param_1 = obj->add(val);
 */
```

```java []
class KthLargest 
{
    int k;
    PriorityQueue<Integer> minHeap = new PriorityQueue<>();

    public KthLargest(int k, int[] nums) 
    {
        this.k = k;
        for (int x : nums)
            minHeap.offer(x);
    }
    
    public int add(int val) 
    {
        minHeap.offer(val);

        while (minHeap.size() > k)
            minHeap.poll();

        return minHeap.peek();
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest obj = new KthLargest(k, nums);
 * int param_1 = obj.add(val);
 */
```