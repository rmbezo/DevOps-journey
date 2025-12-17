result = 5 + 10 # slojenie
result2 = 10 - 5 # vichitanie
result3 = 5 * 5 # umnojenie
result4 = 5 / 2 # деление вернет float

print(result4, type(result4))

#Получаем целое число а не float
result5 = 5 // 2 #цело-численное деление ( возвращает int, а не float )
print(result5, type(result5))

#Получаем остаток деление
result6 = 5 % 2
print(result6, type(result6))

print((5 + 5 * 5 // 10) % 5) # 10