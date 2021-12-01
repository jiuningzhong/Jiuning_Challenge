# Enter your code here. Read input from STDIN. Print output to STDOUT

import re

check_credit_number = re.compile(
    r"^" # starts with
    r"(?!.*(\d)(-?\1){3})" # no consecutive digits repeating 4 times
    r"[456]" # start with 4, 5, 6
    r"\d{3}" # with 3 digits following 4, 5, 6
    r"(?:-?\d{4}){3}" # zero or one occurrences of - and 4 digits appear 3 times
    r"$") # ends with
for _ in range(int(input().strip())):
    print("Valid" if check_credit_number.search(input().strip()) else "Invalid")