plain_message = input("Please input your message \n").strip().lower()  #Strip to remove whitespace at front or back
alphabet = "abcdefghijklmnopqrstuvwxyz"
key = input("Please input the key, in alphabetical order with no spaces in between \n").strip().lower()
print(f"Plain message: {plain_message}")
print(f"Key: {key}")
coded_message = plain_message

for i in range(26):
    print(i)
    coded_message = coded_message.replace(alphabet[i], key[i])
print(coded_message)

