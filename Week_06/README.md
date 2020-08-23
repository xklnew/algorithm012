学习笔记
将复杂问题分解成子问题 并寻找重复性
递归：
```
void recur(int level,int param)
 if(level > MAX_LEVEL)
    return
 procces(level,param);
 recr(level:level+1,newparam);
 //restore curent status
```

分治：(递归的特殊形式)
void divide_conquer(problem,parm1,param2,param3...)
    if problem is None
        print_result
        return
    data = prepare_data(prblem)
    subproblems = split_problem(problem,data)
    
    subResult1 = divide_conquer(subproblem[0],p1,...)
    subResult2 = divide_conquer(subproblem[1],p1,...)
    ...
    result = process_result(subresult1,subreslut2...)
    //revert the current level states
回溯：DFS
贪心：只看眼前最优解，结果不一定是全局最优解
动态规划：分支 + 最优子结构（记录最优解，中途进行淘汰次优解）