# Write a program that asks for an initial integer. The program then repeatedly prompts for integers. If the integer i is a divisor of the initial integer, print 'i is a divisor'.
# If the integer i is not a divisor of the initial integer, print 'i is not a divisor'. The program should continue until it is given the value 0 as input.

integer_input = int(input("")) # store user input in variable

while True: # initiate while loop 
    following_num = int(input(""))
    
    if following_num == 0: # if 0, end program
        break
        
    if integer_input % following_num == 0: # if integer_input divided by following_num = 0, print "is a divisor". otherwise, print "is not a divisor".
        print(f"{following_num} is a divisor")
    else:
        print(f"{following_num} is not a divisor")
