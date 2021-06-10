
#### 思路
![image.png](https://pic.leetcode-cn.com/1605523210-uyStLV-image.png)

- 逐段考察时间间隔 gap，逐段累加:
  - 如果毒效较长，gap 内都起效——则累加 gap。
    - 如果毒效长于 gap，新的攻击的毒效不会叠加。
  - 如果毒效较短，在 gap 内结束了——则累加 duration。

- 注意，最后一次攻击会带来一段 duration 的毒效，和 gap 无关。

### 代码

```javascript []
var findPoisonedDuration = function (timeSeries, duration) {
  if (timeSeries.length == 0) { // 没有攻击发生
    return 0;
  }
  let res = 0;
  for (let i = 1; i < timeSeries.length; i++) {
    const gap = timeSeries[i] - timeSeries[i - 1]; // 时间间隔
    if (duration > gap) { // 覆盖了gap
      res += gap;
    } else {
      res += duration;
    }
  }
  return res + duration; // 补上最后一次攻击的一个duration
};
```

```go []
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	res := 0
	for i := 1; i < len(timeSeries); i++ {
		gap := timeSeries[i] - timeSeries[i-1]
		if duration > gap {
			res += gap
		} else {
			res += duration
		}
	}
	return res + duration
} 
```
#### 感谢阅读，觉得不错的点个赞就更棒了。