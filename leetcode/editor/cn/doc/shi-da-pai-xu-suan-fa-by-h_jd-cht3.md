# 十大排序算法，包含冒泡、插入、二分插入、希尔、快速、选择、堆排、归并，以及计数、桶排、基数
## 冒泡,O(n^2)，改进可通过引入bool型变量控制
```
class Solution {
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        for(int i=0;i<n;i++)
        {
            bool flag=false;
            for(int j=n-1;j>i;j--)
            {
                if(nums[j]<nums[j-1])
                {
                    swap(nums[j],nums[j-1]);
                    flag=true;
                }
            }
            if(flag==false) return nums;
        }
        return nums;
    }
};
```
## 插入,将未排序数组的元素遍历已排序数组，插入到合适位置中，往后的元素依次顺延
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        for(int i=1;i<n;i++)
        {
            int cur=nums[i];
            int index=i-1;
            while(index>=0 && cur<nums[index])
            {
                nums[index+1]=nums[index];
                index--;
            }
            nums[index+1]=cur;
        }
        return nums;
    }
};
```
## 二分插入排序
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        for(int i=1;i<n;i++)
        {
            int cur=nums[i];
            int left=0;
            int right=i-1;
            while(left<=right)
            {
                int mid=(left+right)/2;
                if(cur<nums[mid])
                {
                    right=mid-1;
                }
                else
                {
                    left=mid+1;
                }
            }
            for(int j=i-1;j>=left;--j)
            {
                nums[j+1]=nums[j];
            }  
            nums[left]=cur;
        }
        return nums;
    }
};
```
## 希尔排序,插入排序的改进，解决插入由于经常移动元素不适用于大数据的缺陷
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        int gap=n;
        while(gap>1)
        {
            gap=gap/3+1;
            for(int i=gap;i<n;i++)
            {
                int cur=nums[i];
                int index=i-gap;
                while(index>=0 && cur<nums[index])
                {
                    nums[index+gap]=nums[index];
                    index=index-gap;
                }
                nums[index+gap]=cur;
            }
        }
        return nums;
    }
};
```
## 快速排序,常用 不稳定
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        srand((unsigned)time(nullptr));
        quickSort(nums,0,n-1);
        return nums;
    }
    void quickSort(vector<int>& nums,int start,int end)
    {
        if(start<end)
        {
            int mid=divide(nums,start,end);
            quickSort(nums,start,mid-1);
            quickSort(nums,mid+1,end);
        }
    }
    int divide(vector<int>& nums, int start, int end)
    {
        int i=rand()%(end-start+1)+start;
        swap(nums[start],nums[i]);
        int cur=nums[start];
        while(start<end)
        {
            while(start<end && nums[end]>=cur)--end;
            nums[start]=nums[end];
            while(start<end && nums[start]<=cur)++start;
            nums[end]=nums[start];
        }
        nums[start]=cur;
        return start;
    }
};
```
## 选择排序，每次从n-i个元素里选择最小的元素，作为有序数组的第i个，不稳定
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)return nums;
        for(int i=0;i<n;i++)
        {
            int loc=i;
            for(int j=i;j<n;j++)
            {
                if(nums[j]<nums[loc])
                {
                    loc=j;
                }
            }
            swap(nums[loc],nums[i]);
        }
        return nums;
    }
};
```
## 堆排序，选择排序的优化，不稳定，O(nlogn)
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)return nums;
        return heapSort(nums,n);
    }
    vector<int>heapSort(vector<int>& nums,int n)
    {
        for(int i=n/2-1;i>=0;--i)
        siftDown(nums,i,n-1);
        for(int i=n-1;i>=1;--i)
        {
            swap(nums[0],nums[i]);
            --n;
            siftDown(nums,0,n-1);
        }
        return nums;
    }
    void siftDown(vector<int>& nums,int start,int end)
    {
        int j=2*start+1;
        while(j<=end)
        {
            if(j<end && nums[j]<nums[j+1])++j;
            if(nums[start]>=nums[j])break;
            else
            {
                swap(nums[start],nums[j]);
                start=j;j=2*j+1;
            }
        }
    }
};
```
## 归并排序，分治（与快排思想一样），稳定,需要用到一个辅助数组，每次合并时利用辅助数组排好序再传回原数组。
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)return nums;
        vector<int> temp(n);
        mergeSort(nums,0,n-1,temp);
        return nums;
    }
    void mergeSort(vector<int>& nums,int first,int last,vector<int>& temp)
    {
        if(first<last)
        {
            int mid=(first+last)/2;
            mergeSort(nums,first,mid,temp);
            mergeSort(nums,mid+1,last,temp);
            merge(nums,temp,first,mid,last);
        }
    }
    void merge(vector<int>& nums,vector<int>& temp,int first,int mid,int last)
    {
        int index1=first,index2=mid+1;
        int t=first;
        while(index1<=mid && index2<=last)
        {
            if(nums[index1]<=nums[index2])temp[t++]=nums[index1++];
            else temp[t++]=nums[index2++];
        }
        while(index1<=mid) temp[t++]=nums[index1++];
        while(index2<=last)temp[t++]=nums[index2++];
        for(t=first;t<=last;++t)
        {
            nums[t]=temp[t];
        }
    }
};
```
## 计数排序，记录min&max，通过创建跨度这么大的数组用于计数，然后进行还原。稳定
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)return nums;
        int min=INT_MAX,max=INT_MIN;
        for(int i=0;i<n;++i)
        {
            if(nums[i]>max)max=nums[i];
            if(nums[i]<min)min=nums[i];
        }
        int k=max-min+1;//计数数组长度
        vector<int> count(k,0);
        for(int i=0;i<n;++i)
        {
            count[nums[i]-min]+=1;
        }
        int loc=0;
        for(int i=0;i<k;++i)
        {
            while(count[i]>0)
            {
                nums[loc++]=i+min;
                count[i]-=1;
            }
        }
        return nums;
    }
};
```
## 桶排序，就是分区间，然后元素对应放在每个区间中再排序（插入排序，因为这里假设每个桶的元素都很少），最后顺序输出。
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)return nums;
        int min=INT_MAX,max=INT_MIN;
        for(int i=0;i<n;++i)
        {
            if(nums[i]>max)max=nums[i];
            if(nums[i]<min)min=nums[i];
        }
        int bucketNum=(max-min)/n+1;//确定桶数
        vector<vector<int>> bucket(bucketNum);
        for(int i=0;i<n;++i)//数据放入桶
        {
            int num=(nums[i]-min)/n;
            bucket[num].push_back(nums[i]);
        }
        int count=0;
        for(int i=0;i<bucketNum;++i)
        {
            if(!bucket[i].empty())
            {
                sort(bucket[i].begin(),bucket[i].end());
                for(int j=0;j<bucket[i].size();++j)
                {
                    nums[count++]=bucket[i][j];
                }
            }  
        }
        return nums;
    }
};
```
## 基数排序,以位数开始进行元素排序，时间复杂度为O(n*k)
```
class Solution{
public:
    vector<int> sortArray(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)return nums;
        int min=INT_MAX,max=INT_MIN;
        for(int i=0;i<n;++i)
        {
            if(nums[i]>max)max=nums[i];
            if(nums[i]<min)min=nums[i];
        }
        max=max>(-min)?max:-min;//考虑到负数影响最大值和最小值的位数,选取绝对值最大的数
        int digit=0;
        while(max>0)
        {
            max /= 10;
            
            ++digit;
        }
        vector<vector<int>> bucket(19);
        int pos;
        int cur;
        for(int i=0,mod=1;i<digit;++i,mod*=10)
        {
            for(int j=0;j<n;++j)
            {
                pos=(nums[j]/mod)%10;
                bucket[pos+9].push_back(nums[j]);
            }
            cur=0;
            for(int j=0;j<19;++j)
            {
                for(int k=0;k<bucket[j].size();++k)
                {
                    nums[cur++]=bucket[j][k];
                }
                bucket[j].clear();
            }
        }
        return nums;
    }
};
```
