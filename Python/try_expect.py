try:
    x = float(input("число: "))
    if x > 100:
        print('Ваше число больше ста')
    else:
        print('Ваше число меньше ста')
except ValueError:
    print("Это не число!")