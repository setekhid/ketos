#include <stdio.h>

extern void say_hello(char* p0, ...);
extern void say_bye();

int main() {

	printf("this is a c application.\n");

	say_hello("jack", 4);
	say_bye();

	return 0;
}
