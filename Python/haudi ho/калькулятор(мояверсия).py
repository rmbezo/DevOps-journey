firstnum = float(input("Введите первое число: "))
secondnum = float(input("Введите второе число: "))
whatoperation = input('Введите что вы хотите сделать с этими числами(*, /, +, -). Напишите в точности как в примере: ')

if whatoperation == "*":
    print(firstnum * secondnum)
elif whatoperation == "/":
    print(firstnum / secondnum)
elif whatoperation == "+":
    print(firstnum + secondnum)
else:
    print(firstnum - secondnum)
