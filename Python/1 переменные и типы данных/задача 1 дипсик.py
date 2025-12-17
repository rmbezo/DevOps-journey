#Задача 1: Конвертер времени
#Пользователь вводит количество секунд (целое число). Программа должна перевести это время в формат дни:часы:минуты:секунды.#

seconds = int(input('Введите количество секунд: '))


#Вычисляем дни
days = seconds // 86400
remaining_seconds = seconds % 86400

# Вычисляем часы
hours = remaining_seconds // 3600
remaining_seconds %= 3600

# Вычисляем минуты и секунды
minutes = remaining_seconds // 60
seconds_final = remaining_seconds % 60

print(f"{days}:{hours}:{minutes}:{seconds_final}")