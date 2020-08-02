学习笔记
1.贪心算法  
每次找到当前状态的最优解 并将当前状态发生改变。
2.广度优先（bfs）//层级的感觉
que.push(root)
visted.add(root)
while(!que.empty())
{
    Tree* node = que.poll();
    visted.add(nide)
    process(node)
    nodes = generate(node);
    que.push(nodes)
}
3.深度优先(dfs)
if(vist[i])
{
   return;
}
vist.add(node)
dfs(node.left)
dfs(node.right)

4.二分查找
int lo =0,hi=len-1;
while(lo <= hi)
{
    int mid = (lo + hi)/2;
    if(v[mid] == target)
    {
        //找到返回
        return;
    }
    else if(v[mid] < target)
    {
        lo = mid + 1;
    }
    else{
        hi = mid - 1;
    }
}

