#include <vector>
#include <iostream>
using namespace std;

class Solution {
  public:
    vector<int> countBits(int n) {
      vector<int> res(n+1, 0);
      for (int i = 1; i <= n; i++) {
        res[i] = res[i >> 1] + (i & 1);
      }
      return res;
    }
};

int main() {
  Solution s;
  vector<int> res = s.countBits(5);
  for (int i = 0; i < res.size(); i++) {
    cout << res[i] << '\n';
  }
  return 0;
}
