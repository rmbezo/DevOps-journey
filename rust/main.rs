use std::io;

fn main() {
    println!("--- Простой калькулятор на Rust ---");

    // Чтение первого числа
    println!("Введите первое число:");
    let mut input1 = String::new();
    io::stdin().read_line(&mut input1).expect("Ошибка чтения");
    let num1: f64 = input1.trim().parse().expect("Пожалуйста, введите число");

    // Чтение оператора
    println!("Введите операцию (+, -, *, /):");
    let mut operator = String::new();
    io::stdin().read_line(&mut operator).expect("Ошибка чтения");
    let operator = operator.trim();

    // Чтение второго числа
    println!("Введите второе число:");
    let mut input2 = String::new();
    io::stdin().read_line(&mut input2).expect("Ошибка чтения");
    let num2: f64 = input2.trim().parse().expect("Пожалуйста, введите число");

    // Вычисление
    let result = match operator {
        "+" => num1 + num2,
        "-" => num1 - num2,
        "*" => num1 * num2,
        "/" => {
            if num2 == 0.0 {
                println!("Ошибка: Деление на ноль!");
                return;
            }
            num1 / num2
        }
        _ => {
            println!("Неизвестная операция");
            return;
        }
    };

    println!("Результат: {}", result);
}