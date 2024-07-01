plain_message = input("Please input the message you wish to encrypt: \n").strip().lower()  #Strip to remove whitespace at front or back
alphabet = "abcdefghijklmnopqrstuvwxyz "
key = input("Please input the key, in alphabetical order with no spaces in between \n").strip().lower()
key += " "
print(f"Plain message: {plain_message}")
print(f"Key: {key}")
coded_message = " "
counter = 0

for letter in plain_message:
    for i in range(27):
        if letter == alphabet[i]:
            coded_message += key[i]
    counter += 1

print(coded_message)

