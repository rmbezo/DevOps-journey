## Arrays, strokes, vector

### array:
int a[5] = {10, 20, 30, 40, 50};


### vector is dynamic array

vector<int> a = {1, 2, 3};

vector features:

```
a.size() // size of vector - len()

a.push_back(10) // add number 10 in the end of vector

a.pop_back() // Removes final element and reduce size by 1 

a.empty() // Check if it's empty
```


### Strokes / strings
```
string s = "hello";
cout << s[0]; // h
cout << s[1]; // e
```

size of stroke:
`cout << s.size() << '\n'; // 5`


To scan stroke:
```
string name;
cin >> name; // hello -> hello, hello world -> hello ! 

==> for whole stroke:
getline:
string s;
getline(cin, s); // hello world --> hello world // whole stroke with spaces

```


### String is similiar to vector<char>

```
string s = "abc";
s.push_back('d');
cout << s << "\n"; // abcd

// We can change symbols
s[0] = 'z';
cout << s << "\n"; // zbcd
```




