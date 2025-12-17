#Задача 1.1: Сумма чётных чисел
# import math

# num1 = int(input('Введите число: '))
# result = []

# for i in range(2,num1,2):
#     result.append(i)
    
# total = sum(result)
# print(total)



#Задача 1.2: Лесенка из символов

# lol = '#'

# while lol != "###############################################################################":
#     if lol == '#':
#         print('#')
#         lol += '#'
#     else:
#         print(lol)
#         lol += '#'



#Задача 2.1: Фильтр чисел

# list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
# right = []

# for i in list:
#     if i % 3 == 0:
#         continue
#     else:
#         right.append(i)

# result = right
# right = result.append(sum(right))
# print(result, "длинна списка:", len(result))


#Задача 3.1: Калькулятор площади


# import math

# figure = input('Выберите фигуру(круг/прямоугольник/треугольник)').lower()



# def calculate_area(shape):
#     if shape == 'круг':
#         ask1 = float(input('Введите радиус треугольника: '))
#         res1 = ask1 ** 2 * math.pi
#         print(res1)
#     elif shape == 'прямоугольник':
#         ask2 = float(input('Введите ширину прямоугольника: '))
#         ask3 = float(input('Введите боковую сторону(высоту) прямоугольника: '))
#         res2 = ask2 * ask3
#         print(res2)
#     elif shape == 'треугольник':
#         ask4 = float(input('Введите основание треугольника: '))
#         ask5 = float(input('Введите высоту треугольника: '))
#         print(1/2 * ask4 * ask5)
#     else:
#         print('Вы написали некоректную фигуру.')


# calculate_area(figure)




#Задача 3.2: Проверка пароля 

# paswrd = input('Введите пароль: ')

# def validate_password(password):
#     if len(password) >= 8 and any(ch.isdigit() for ch in password):
#         print(True)
#     else:
#         print(False)
    
# validate_password(paswrd)




#Задача 4.1: Класс "Студент"


# class student():
#     def __init__(self, name, grade, average):
#         self.name = name
#         self.grade = grade
#         self.average = average
#     def add_grade(grades):
#         list_grades = []
#         list_grades.append(grades)
#         print(list_grades)
#     def get_average(list_of_grades):
#         print(list_of_grades)

# student1 = student('Max',5, 5)
# student1.add_grade(5)