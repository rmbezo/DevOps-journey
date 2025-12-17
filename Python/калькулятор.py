

# fnum = float(input('Введите первое число: '))
# action = input('Введите действие(*, /, **, +, -): ')
# snum = float(input('Введите второе число: '))
# ask = input('Хотели бы вы добавить ещё действие?(да/нет): ')

# if ask.lower() == 'да':
#     while ask.lower() == 'да':
#         action2 = input('Введите действие(*, /, **, +, -): ')
#         tnum = float(input('Введите третье число: '))
#         ask = input('Хотели бы вы добавить ещё действие?(да/нет): ')
# elif ask.lower() == 'нет':
#     if action == "*":
#         print(round(fnum * snum))
#     elif action == "/":
#         print(round(fnum / snum))
#     elif action == "**":
#         print(round(fnum ** snum))
#     elif action == "+":
#         print(round(fnum + snum))
#     elif action == "-":
#         print(round(fnum - snum))
# else: 
#     print('error') 

    

#ВВерсия дипсик


# result = float(input('Введите первое число: '))

# while True:
#     action = input('Введите действие (*, /, **, +, -) или "стоп" для завершения: ')
#     if action.lower() == 'стоп':
#         break
    
#     if action not in ('*', '/', '**', '+', '-'):
#         print('Ошибка: неизвестное действие!')
#         continue
    
#     num = float(input('Введите следующее число: '))
    
#     if action == '*':
#         result *= num
#     elif action == '/':
#         if num == 0:
#             print('Ошибка: деление на ноль!')
#             continiue
#         result /= num
#     elif action == '**':
#         result **= num
#     elif action == '+':
#         result += num
#     elif action == '-':
#         result -= num
    
#     print(f'Промежуточный результат: {result}')

# print(f'Итоговый результат: {result}')


# result = float(input('Введите первое число: '))

# while True:
#     action = input('Введите действие (*, /, **, +, -): ')
#     num = float(input('Введите следующее число: '))
    
#     if action == '*':
#         result *= num
#     elif action == '/':
#         if num == 0:
#             print('Ошибка: деление на ноль!')
#             continue
#         result /= num
#     elif action == '**':
#         result **= num
#     elif action == '+':
#         result += num
#     elif action == '-':
#         result -= num
#     else:
#         print('Ошибка: неизвестное действие!')
#         continue
    
#     print(f'Текущий результат: {result}')
    
#     ask = input('Добавить ещё действие? (да/нет): ')
#     if ask.lower() != 'да':
#         break

# print(f'Итоговый результат: {result}')

