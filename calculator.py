
operator = input("Masukkan operator (+ - * /): ")
num1 = float(input("masukkan angka pertama: "))
num2 = float(input("masukkan angka kedua: "))

if operator == "+": 
    result = num1 + num2
    print(round(result, 3))
elif operator == "-":
    result = num1 - num2
    print(round(result, 3))
elif operator == "*":
    result = num1 * num2
    print(round(result, 3))
elif operator == "/":
    result = num1 / num2
    print(round(result, 3))
else:
    print(f"{operator}bukan operator yang tepat")