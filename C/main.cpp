#include <iostream>
#include <vector>
using namespace std;

int f(int x) {
  if (x > 0) {
    return x * f(x-1);
  } else {
    return 1;
  }
}

int fmax(vector<int>& a) {
  int mx = a[0];
  int lenA = (int)a.size();
  for (int i = 1; i < lenA; i++) {
    if (a[i] > mx) { 
      mx = a[i];
    }
    if (i+1 >= (int)a.size()) {
      a.push_back(mx);
    }
  }
  return mx;
}

void printV(const vector<int>& a) {
  for (int i = 0; i < (int)a.size(); i++ ) {
    cout << (int)(a[i]) << " ";
  }
}

int main() {
  int x;
  cin >> x;
  cout << f(x) << "\n" ;
  vector<int> l = {9, 2, 4, 1923, 999};
  vector<int>* pl = &l;
  cout << fmax(*pl);
  cout << pl << "\n";
  printV(*pl);
  cout << "\n"; 
  printV(l);
  return 0;
}
