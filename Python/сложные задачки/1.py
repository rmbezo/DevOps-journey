# import random

# randomnum = random.randint(1,100)

# guess = int(input('Угадай число от 1 до 100: '))
# i = 0

# while guess != randomnum:
#     i += 1
#     if i == 5:
#         if (randomnum % 2) == 0:
#             print('Подсказка: число - четное')
#         else:
#             print('Подсказка: число не четное')
#     if i == 5:
#         bait = input('Может вы готовы сдаться? (напишите "да" или "нет")')
#         if bait == 'нет':
#             continue
#         elif bait == 'да':
#             print(f'загаданное число было - {randomnum}! Ваше количество попыток {i}')
#             print(input('Нажмите enter что бы продолжить'))
#     if guess > randomnum and 0 < guess < 101:
#         print('Ваше число больше чем загаданное')
#         guess = int(input('Угадай число от 1 до 100 (или если вы сдаетесь введите - 0): '))
#         continue
#     elif guess < randomnum and 0 < guess < 101:
#         print('Ваше число меньше загаданного')
#         guess = int(input('Угадай число от 1 до 100 (или если вы сдаетесь введите - 0): '))
#         continue
#     elif guess == 0:
#         print(f'загаданное число было - {randomnum}! Ваше количество попыток {i}')
#         print(input('Нажмите enter что бы продолжить'))
#     else:
#         print('Введенное число находится вне диапазона')
# 
# 
# 
# if guess == randomnum and i == 0:
#     print(f'Вы угадали число - {randomnum}! Ваше количество попыток - 1')
# elif guess == randomnum:
#     print(f'Вы угадали число - {randomnum}! Ваше количество попыток - {i}')


import random

randomnum = random.randint(1, 100)
attempts = []  # Список для хранения всех попыток

while True:
    guess_input = input('Угадай число от 1 до 100 (или "сдаюсь"): ').strip().lower()
    
    # Проверка на "сдаюсь"
    if guess_input == "сдаюсь":
        print(f'Загаданное число было: {randomnum}! Количество попыток: {len(attempts)}')
        print("Ваши попытки:", sorted(attempts))
        input('Нажмите Enter чтобы выйти...')
        break
    
    # Проверка, что введено число
    if not guess_input.isdigit():
        print("Ошибка! Введите число.")
        continue
    
    guess = int(guess_input)
    
    # Проверка диапазона
    if guess < 1 or guess > 100:
        print("Число должно быть от 1 до 100!")
        continue
    
    attempts.append(guess)  # Добавляем попытку в список
    
    if guess == randomnum:
        print(f'Поздравляю! Ты угадал за {len(attempts)} попыток.')
        print("Твои попытки:", sorted(attempts))
        break
    elif guess > randomnum:
        print("Слишком много!")
    else:
        print("Слишком мало!")
    
    # Подсказка после 5 попыток
    if len(attempts) == 5:
        if randomnum % 2 == 0:
            print("Подсказка: число четное!")
        else:
            print("Подсказка: число нечетное!")