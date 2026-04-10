integer_input = int(input(""))

while True:
    following_num = int(input(""))
    
    if following_num == 0:
        break
        
    if integer_input % following_num == 0:
        print(f"{following_num} is a divisor")
    else:
        print(f"{following_num} is not a divisor")