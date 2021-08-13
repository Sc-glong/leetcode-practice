### 解题思路
第一次写直接累加除数，最后超时了。因此考虑位运算速度较快，而因为位运算每次是2倍，因此我们考虑这样的算法：

【第一步：】
    除数tempDiv依次乘以2（左移一位），直到大于被除数dividend。
    tempDiv除以2（右移一位）得到最大除数，进行下一步。

【第二步：】
    被除数dividend减去第一步中tempDiv，得到diff。
    将diff作为被除数dividend（dividend=diff），重复第一步。

【终止条件：】
    如果diff小于除数，则终止。

【商是这样计算的：】
    每次递归时设置临时商tempRes=1，除数每左移一位，临时的商tempRes也左移一位。
    每次递归结束后，总商result加上临时商tempRes

上面的只是思路，具体还有很多细节需要看代码处理的方式。

### 代码

```c

#define INT_MIN -2147483648
#define INT_MAX 2147483647

//递归求
void div0(int diff, int divisor0, int *res){
    int tempDIV = divisor0,tempRes = 0;
    while(diff >= tempDIV){
        if(tempDIV>(INT_MAX>>2)){//溢出
            if(tempRes==0) tempRes = 1;
            diff -= tempDIV;
            break;
        }
        tempDIV = tempDIV<<1;
        if(diff < tempDIV){//大于差
            if(tempRes==0) tempRes = 1;
            diff -= tempDIV>>1;//yuanshu
            break;
        }else{
            if(tempRes == 0) tempRes = 2;
            else tempRes = tempRes<<1;
        }
    }
    *res += tempRes;
    if(diff < divisor0) return;
    else div0(diff, divisor0, res);
}

int divide(int dividend, int divisor){
    int sign = (dividend>0) == (divisor>0);//dir: 1+, 0-
    int result = 0;
    //特殊值处理
    if(dividend==0) return 0;
    if (divisor == INT_MIN) {
        if (dividend == INT_MIN) {
            return 1;
        } else {
            return 0;
        }
    }
    if (dividend == INT_MIN) {
        if (divisor == -1) {
            return INT_MAX;
        } else if (divisor == 1) {
            return INT_MIN;
        }
        dividend += abs(divisor);
        result++;
    }
    if(divisor == 1){
        return dividend;
    }else if(divisor == -1){
        return -dividend;
    }
    //转换为正数
    if(dividend<0){
        dividend = -dividend;
    }
    if(divisor<0){
        divisor = -divisor;
    }
    //正整数除法运算    
    div0(dividend, divisor, &result);

    //符号处理
    result = sign ? result:-result;
    return result;
}
```