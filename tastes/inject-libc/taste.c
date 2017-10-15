#include <string.h>
#include <stdio.h>


int main() {

	char const *hello = "hello, world!";
	char *ind = strrchr(hello, 'w');
	printf("last char is at %li\n", ind - hello);

	return 0;
}
