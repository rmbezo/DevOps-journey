class user():
    def __init__(self, name : str, age : int, kastratsiya : str):
        self.name = name
        self.age = age
        self.kostratsiya = kastratsiya
    
    def cat_info(self):
        print(f'Имя кота: {self.name}, его возраст: {self.age}, кастрирован ли: {self.kostratsiya}')

user1 = user('Валера', 9, "да")
user1.cat_info()