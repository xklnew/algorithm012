//191. 位1的个数
class Solution {
public:
int hammingWeight(uint32_t n) {
	int count  = 0;
	while(n )
	{
		count++;
		n &= n-1;
	}
	return count;
}
};
//231. 2的幂
class Solution {
public:
bool isPowerOfTwo(int n) {
	if(n == 0)
		return false;
	n &= n-1;
	return n  == 0;
}
};