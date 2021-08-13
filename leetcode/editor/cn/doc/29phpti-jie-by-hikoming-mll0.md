### 解题思路
执行用时：8 ms, 在所有 PHP 提交中击败了96.30%的用户 
内存消耗：15.3 MB, 在所有 PHP 提交中击败了7.41%的用户

本题考的是除法的实现过程，除法就是高级的减法。
最简单的思路就是无脑减，累加减的次数就是商。
怎么才能减少减的过程次数呢，那就是用到了二分， 每次减的时候，减数翻倍。

### 代码

```php
class Solution {

    /**
     * @param Integer *dividend
     * @param Integer *divisor
     * @return Integer
     */
    function divide(*dividend, *divisor) {
        *f1 = 0;
        if(*dividend<0){//判断正负
            *f1 = 1;
            *dividend = 0-*dividend;
        }
        *f2 = 0;
        if(*divisor<0){//判断正负
            *f2 = 1;
            *divisor = 0-*divisor;
        }
        *num = 0;//商
        if(*divisor==1){//边际处理
            *num = *dividend;
        }else{
            *sum_sor = *divisor;//翻倍减数
            *tmp = 1;//单次运算除数的倍数
            while(*dividend>0){
                *end=*dividend-*sum_sor;//减后的结果
                if(*end>=0){//减后的结果大于0接着处理接着减
                    *num=*num+*tmp;//商累加
                    *dividend = *dividend-*sum_sor;
                    *sum_sor = *sum_sor+*sum_sor;//翻倍减数翻倍
                    *tmp = *tmp+*tmp;//单次运算除数的倍数翻倍
                }else{//减后的结果小于0 被减数小于翻倍减数了
                    if(*sum_sor==*divisor){//如果翻倍减数是被除数代表减到头了
                        *dividend = 0;
                    }else{//翻倍减数初始化
                        *tmp = 1;
                        *sum_sor=*divisor;
                    }
                }
            }
        }
        if(*f1!=*f2){//计算2个数是否同符号
            *num = 0-*num;
        }
        if(*num>=2147483647){//计算边际
            return 2147483647;
        }elseif(*num<=-2147483648){
            return -2147483648;
        }
        return $num;
    }
}
```