#include <stdio.h>

int main() {
  printf("I'm a Bot. What's your name?\n");
  char name[20];
  scanf("%s", name);
  printf("Hello, %s! How old are you?\n", name);
  int age;
  scanf("%d", &age);
  printf("You are look younger!"
         " I thought that you are %d!\n", age-3);

}
