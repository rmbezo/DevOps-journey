from calc import calc

def test_add():
    assert calc(2, "+", 2) == 4

def test_div():
    assert calc(10, "/", 2) == 5

def test_mul():
    assert calc(3, "*", 4) == 12
