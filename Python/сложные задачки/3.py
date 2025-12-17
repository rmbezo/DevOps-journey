#1 
# этот код выводит числа из списка умноженные на два

# for i in [1, 2, 3, 4, 5]:
#     print(i ** 2)
    

#2 
# я не знаю что делает функция list, но могу предположить что выведется список в котором будут числа от 3 до 7 с шагом 2

# print(list(range(10, 21, 2)))

#3
# Список будет выглядить так: [1, 2, 3, 4, 5, 6] т.к. extend добавляется в список без скобок

# a = [10, 20, 30]
# a.append(40)
# a.remove(20)

# print(a)

#4
#Код вернет список ['Привет', 'мир!']

# str1 = ['Python']
# str1.sort(reverse = True)
# print(str1)
# Я не знаю как это сделать

#5 
# Эта функция принимает значение name и выводит приветсвие name

# number = int(input('Введите число: '))

# def if_even(num):
#     if (num % 2) == 0:
#         print(True)
#     else:
#         print(False)

# if_even(number)

#6
#Этот код создает класс который принимает имя животного, а затем запускает функцию, которая принимает имя и создает строку с этим именем и его действием

# class Cat:
#     def __init__(self, name):
#         self.name = name
    
#     def Meow(self):
#         print(f'{self.name} говорит "Мяууу"')

# cat = Cat('Дипсик')
# # cat.Meow()
# # Или
# Cat.Meow(cat)

#7

# numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

# def filter_odd(numbers):
#     for i in numbers:
#         if (i % 2) != 0:
#             print(i)
#         else:
#             continue

# filter_odd(numbers)


#Задача 1: Из строки "Hello, World!" получите "World".

# str1 = "Hello, World!"
# print(f'"{str1[-6: -1]}"')

#Задача 2: Переверните список [10, 20, 30, 40] без использования reverse().

# list = [10, 20, 30, 40]
# print(list[::-1])