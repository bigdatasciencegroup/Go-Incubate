#include <stdio.h>

static int a=20;
void xyz()
{
	printf("%d, ", a);
	a = 100;
}
int main()
{
	printf("%d, ", a);
	
		static int a = 10;
		printf("%d, \n", a);
		xyz();
	
	printf("%d", a);
}