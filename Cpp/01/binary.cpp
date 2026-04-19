#include <iostream>
#include <vector>
using namespace std;

int binary_search(const vector<int>& a, int s) {
  int left = 0;
  int right = a.size() - 1;
  int res;
  int steps = 0;
  int mid = (left + right) / 2;
  while (left <= right) {
    steps += 1;
    mid = left + (right - left) / 2;
    if (s > a[mid]) {
      left = mid + 1;
    } else if (s < a[mid]) {
      right = mid - 1;
    } else {
      cout << "binary steps: " << steps << "\n";
      return res = mid;
    }
  }
  return -1;
}

int simple_search(const vector<int>& a, int s) {
  int steps = 0;
  for (int i = 0; i < (int)a.size(); i++) {
    steps += 1;
    if (s == a[i]) {
      cout << '\n' << "simple search steps: " << steps << '\n';
      return i;
    }
  }

  return -1;
}

int main() {
  vector<int> l = {10, 20, 30, 40, 50, 60, 70, 80, 90, 100};
  cout << binary_search(l, 20);
  cout << simple_search(l , 90);
  return 0;
}
