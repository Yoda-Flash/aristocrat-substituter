coded_message = input("Please enter your coded message \n").strip().lower()
# hdooe hen ayd bem kegif
alphabet = "abcdefghijklmnopqrstuvwxyz "
key = input("Please input the key, in alphabetical order with no spaces in between \n").strip().lower()
key += " "
# alskdjfhgpqowieurytzmxncbv
plain_message = " "
counter = 0

for letter in coded_message:
    for i in range(27):
        if letter == key[i]:
            plain_message += alphabet[i]
    counter += 1

print(plain_message)