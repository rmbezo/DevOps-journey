#include <iostream>
#include <string>
using namespace std;

int main() {
  string name;
  cout << "Type your name: ";
  cin >> name;

  cout << "Hello, " << name << "!\n";
  cout << "I am your personal assistant Jarvis\n";

  int num1, num2;
  
  cout << "Type first number: ";
  cin >> num1;

  cout << "Choose an operator + and -: ";
  string op;
  cin >> op;

  cout << "Type second number: ";
  cin >> num2;

  if (op == "+") { 
    cout << num1 + num2 << endl;
  } else if (op == "-") {
    cout << num1 - num2 << endl;
  } else {
    cout << "Not an operator!";
  }

  return 0;
}

