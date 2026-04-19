#include <iostream>
#include <vector>
#include <string>
using namespace std;

bool is_valid(string& s) {
  vector<char> l;
  for (int i = 0; i < s.size(); i++) {
    if (s[i] == '(' || s[i] == '{' || s[i] == '[') {
      l.push_back(s[i]);
    } 
     
    if (s[i] == ')') {
      if (l[l.size()-1] != '(') {
        return false;
      } else {
        l.pop_back();
      }
    } else if (s[i] == '}')  {
      if (l[l.size()-1] != '{') {
        return false;
      } else {
        l.pop_back();
      }

    } else if (s[i] == ']') {
      if (l[l.size()-1] != '[') {
        return false;
      } else {
        l.pop_back();
      }
    }
  }
  return l.empty();
}

int main() {
  string test = "()(())[]{]}[][]{}{}{}";
  cout << (is_valid(test) ? "True" : "False") << "\n";
  return 0;

}
