## Vector, pair, tuple

```
vector: 
a.push_back(10);   // добавить в конец
a.pop_back();      // удалить последний
a.size();          // текущий размер
a.empty();         // пустой ли
a[0];              // доступ по индексу
a.at(0);           // доступ с проверкой границ
```


### Pair:
 Object of a 2 elements, even different types

```
pair<int, int> p = {3, 7};
cout << p.first << " " << p.second << "\n";
```

Pair compare:

first with first, second with second
```
pair<int, int> a = {1, 5};
pair<int, int> b = {2, 1};

cout << (a < b) << "\n"; // 1 (true)
```

### Vector pair

```
vector<pair<int, int>> a;
a.push_back({3, 10});
a.push_back({2, 11});
a.push_back({4, 12});

cout << a[0].first << " " << a[0].second << "\n";
```


### Tuple 

Same as pair, but works with more elements

```
#include <tuple>
#include <iostream>

using namepsace std;

int main() {
  tuple<int, string, double> t = {5, "abc", 3.14};

  cout << get<0>(t) << "\n"; // 5
  cout << get<1>(t) << "\n"; // abc
  cout << get<2>(t) << "\n"; // 3.14
}
```



Examples:
pair<int,int > - координата (x, y)
tuple<int,int,int> - (x, y, cost)


## Structured bindings
```
pair<int, int> p = {4, 9};
auto [x, y] = p;

tuple<int, int, int> t = {1, 2, 3};
auto [a, b, c] = t;
```

```
vector<pair<int,int>> v = {{1,2}, {3,4}};

for (auto [x, y] : v) {
    cout << x << " " << y << "\n";
}
```


