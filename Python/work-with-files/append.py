file = open("data/test2.txt", "a")
count = int(input("Type the count of strings that you want to write: "))

if type(count) == int:
    for i in range(count):
        file.write("Privet mir\n")
        print("The count of wroten strings is: ", i)
    else: 
        print("The cycle is end!")
else:
    print("Wrong type of count!!")


file.close()
input("Press Enter to close programm: ")

