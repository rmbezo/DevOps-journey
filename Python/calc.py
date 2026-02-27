

num1 = float(input("Type first number: "))
op = input("Type operator: ")
num2 = float(input("Type second num: "))

if op == "+":
    result = num1 + num2
    print("Answer is ", result)
elif op == "-":
    result = num1 - num2
    print("Answer is ", result)
elif op == ":" or op == "/":
    if num2 != 0:
        result = num1 / num2
        print("Answer is ", result)
elif op == "*":
    result = num1 * num2
    print("Answer is ", result)
elif op == "**" or op == "^":
    result = num1 ** num2
    print("Answer is ", result)
else:
    print("Non an operator! Allowed operators is: + ; - ; * ; ** ; ^ ; / ; :")

input("Press ENTER to exit: ")

