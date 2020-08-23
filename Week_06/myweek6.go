//最小路径和
class Solution {
public:
void printAll(vector<int> &result)
{
	for(int i=0;i<result.size();i++)
	{
		printf("%d ",result[i]);
	}
	printf("\n");
}
int minPathSum(vector<vector<int>>& grid)
{
	if(grid.size() <= 0 || grid[0].size() <=0 )
		return -1;
	int m= grid.size(),n=grid[0].size();
	vector<int> res(n);

	res[0] = grid[0][0];
	for(int k=1;k<n;k++)
	{
		res[k] = res[k-1] + grid[0][k];
	}
	printAll(res);//[1,3]
	for(int i=1;i<m;i++)
	{
		for(int j=0;j<n;j++)
		{
			int min = res[j];
			if(j-1 >= 0)
			{
				min = min > res[j-1] ? res[j-1]:min;
			}
			res[j] = min + grid[i][j];
		}
		return res[n-1];
	}
};

//最大正方形
class Solution {
public int maximalSquare(char[][] matrix) {
	int n = matrix.length;
	if(n == 0) return 0;
	int m = matrix[0].length;
	if(m == 0) return 0;
	int[][] dp = new int[n][m];
	int result = 0;
	//初始化base case
	for(int i = 0 ;i<n;i++) {
	dp[i][0] = matrix[i][0] - '0';
	result = Math.max(dp[i][0],result);
	}
	for(int i = 0 ;i<m;i++){
		dp[0][i] = matrix[0][i] - '0';
		result = Math.max(dp[0][i],result);
	}
	for(int i = 1; i<n; i++){
		for(int j = 1; j<m; j++){
//计算dp[i][j];
			if(matrix[i][j] == '1'){
			dp[i][j] = Math.min(Math.min(dp[i-1][j-1],dp[i-1][j]),dp[i][j-1]) + 1;
			result = Math.max(dp[i][j],result);
			}
		}
	}
	return result*result;
}
}