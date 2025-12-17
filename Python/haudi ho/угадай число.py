import random
# +, -, *, **, //, %,  унарный минус, округление, число Пи



#print(random.randint(1,5))


randomcount = (random.randint(1,10))

guess = int(input('Угадайте загаданное число от 1 до 10: '))

if randomcount == guess:
    print('Вы угадали "число" ')
else:
    print("Неверно! Число было:",randomcount)

print("Завершение программы.... ")

input('')