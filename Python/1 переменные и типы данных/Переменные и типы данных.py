#Переменные и типы данных

#Переменная
name = 10 #int
lol = "Valera" #str
lol2 = 'Vladimir' # = - это присваивание 

print(name,lol,lol2) #10 Valera Vladimir

#Типы данных

str = 'dwdw' #строка буквы 
int = 102 #целые числа
bool = True #Логические True/False
float = 3.5 #дроби и числа с плавающей запятой

print(str,int,bool,float) #dwdw 102 True 3.5

print(type(str)) # str
print(type(int)) # int
print(type(bool)) # bool
print(type(float)) # float

# Переменная не может начинаться с цифр и специальных знаков : ^*$#$(@)&?

# =============================== INPUT ======================================================

#input('kommentariy')

yourname = input('Введите ваше имя: ')
print('Ваше имя:',yourname) #Ваше имя: то что вы ввели

#Input всегда str. ==>

print(type(yourname)) # ВСЕГДА БУДЕТ STR

# ========================================== Int ========================================================

dasdasd = '25'
print(type(dasdasd)) #str
intdasdasd = int(dasdasd)
print(type(intdasdasd)) #int