import asyncio
from aiogram import Bot, Dispatcher, types
from aiogram.filters import Command
from aiogram.types import Message, ReplyKeyboardMarkup, KeyboardButton

# Указываем токен напрямую (для тестов)
TOKEN = "8332063959:AAEFb_1LQT_DThCY19cX5p8pQUFoxjiAMX0"

# Создаем экземпляры бота и диспетчера
bot = Bot(token=TOKEN)
dp = Dispatcher()

# Создаем клавиатуру
def make_keyboard() -> ReplyKeyboardMarkup:
    buttons = [
        [KeyboardButton(text="💵 Конвертер валют")],
        [KeyboardButton(text="📏 Конвертер единиц")],
        [KeyboardButton(text="🔍 Случайный факт")]
    ]
    return ReplyKeyboardMarkup(keyboard=buttons, resize_keyboard=True)

# Обработчик команды /start
@dp.message(Command("start"))
async def start_handler(message: Message):
    keyboard = make_keyboard()
    await message.answer(
        "Привет! Я бот-конвертер. Выбери действие:",
        reply_markup=keyboard
    )

# Обработчик текстовых сообщений
@dp.message()
async def message_handler(message: types.Message):
    if message.text == "💵 Конвертер валют":
        valutetype = await message.input('Напишите какую валюту вы хотите конвертировать: ')
        if valutetype == "Евро":
            print('sex')
    elif message.text == "📏 Конвертер единиц":
        await message.answer("Выбран конвертер единиц (будет реализован)")
    elif message.text == "🔍 Случайный факт":
        await message.answer("Вот случайный факт: Python был создан в 1991 году!")
    else:   
        await message.answer("Пожалуйста, используйте кнопки или команду /start")

# Запуск бота
async def main():
    await dp.start_polling(bot)

if __name__ == "__main__":
    asyncio.run(main())