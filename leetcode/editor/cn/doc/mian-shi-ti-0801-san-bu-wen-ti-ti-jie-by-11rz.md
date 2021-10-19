### 解题思路
此题用**动态规划**去做。
dp 是一个行数为 1 的矩阵（用一个**二维数组**表示）
dp[0][i]表示不同的状态。
# 具体步骤：
- **首先，要明确每一阶可以有哪些状态。（一步可以走1阶、2阶或3阶）**
1. 状态0：当前是这一步的第一阶
![image.png](https://pic.leetcode-cn.com/1629356954-leHSZX-image.png)

2. 状态1：当前是这一步的第二阶
![image.png](https://pic.leetcode-cn.com/1629356975-UCzHyr-image.png)

3. 状态2：当前是这一步的第三阶
![image.png](https://pic.leetcode-cn.com/1629356993-aLAHFc-image.png)

- **当前可由前一阶转移过来**
1. 当前状态0 = (之前) 状态0 + 状态1 + 状态2
2. 当前状态1 = (之前) 状态0
3. 当前状态2 = (之前) 状态1
可以发现：每一次的状态转移的系数都是不变的，运用线性代数的知识，我们可以利用**矩阵相乘**来表示状态转移

- **明确系数矩阵（工具矩阵）和初始化结果矩阵（dp）**
1. 系数矩阵为
![image.png](https://pic.leetcode-cn.com/1629357404-IcZoGN-image.png)

2. 初始化：当仅有一阶台阶时，很明显只能是状态 0 ,不存在状态1 和状态2 ，因此 dp[0] = { 1, 0, 0 };

- **取模处理**
1. 因为n越大，数据会快速增长，很快就会超过int型大小造成溢出，因此要利用快速乘法和快速幂算法，避免数据溢出。

- emmm，表达能力一般般，希望可以帮助到你，欢迎留言！
### 代码

```cpp
class Solution {
public:
    int waysToStep(int n) {
        int ans = 0;
        vector<vector<int> > dp = {{1,0,0}};
        vector<vector<int> > tmp = {{1,1,0},{1,0,1},{1,0,0}};
        tmp = power(tmp,n-1);
        dp = multiply(tar,tmp);
        for(int i : dp[0]){
            ans += i;
            ans %= mod;
        }
        return ans;
    }
private:
    const int mod = 1e9 + 7;
    //快速乘法 (a*b)%mod
    int mul(int a,int b){
        int ans = 0;
        while(b){
            if(b&1) ans = (ans+a)%mod;
            b >>= 1;
            a = (a<<1)%mod;
        }
        return ans % mod;
    }
    //矩阵乘法
    vector<vector<int> > multiply(vector<vector<int> > a,vector<vector<int> > b){
        const int row = a.size() , col = b[0].size() , len = a[0].size();
        vector<vector<int> > ans(row,vector<int>(col,0));
        for(int i=0 ; i<row ; ++i){
            for(int j=0 ; j<col ; ++j){
                for(int k=0 ; k<len ; ++k){
                    ans[i][j] += mul(a[i][k] , b[k][j]) ;
                    ans[i][j] %= mod; //取模
                }
            }
        }
        return ans;
    }
    //矩阵快速幂
    vector<vector<int> > power(vector<vector<int> > a , int n){
        const int len = a.size();
        vector<vector<int> > ans(len,vector<int>(len,0));
        for(int i=0 ; i<len ; ++i) ans[i][i] = 1; //初始化为单位矩阵
        while(n){
            if(n&1) ans = multiply(ans , a);
            a = multiply(a,a);
            n >>= 1;
        }
        return ans;
    }
};
```