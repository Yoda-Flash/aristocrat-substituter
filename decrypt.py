coded_message = input("Please enter your coded message \n").strip().lower()
# hdeeen hen abd bem begif
alphabet = "abcdefghijklmnopqrstuvwxyz"
key = input("Please input the key, in alphabetical order with no spaces in between \n").strip().lower()
# alskdjfhgpqowieurytzmxncbv
plain_message = coded_message

for i in range(26):
    print(i)
    plain_message = plain_message.replace(key[i], alphabet[i])
print(plain_message)