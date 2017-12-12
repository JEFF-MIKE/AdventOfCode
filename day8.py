#day 8
def day8i():
    # register systems inc boiii
    reg_dic = {} # dictionary with registers as strings, all at 0
    high = 0
    maximum = 0
    fileR = open("day8.txt","r")
    for line in fileR:
        splitter1 = line.strip().split()
        if splitter1[0] not in reg_dic:
            reg_dic[splitter1[0]] = 0
        if splitter1[3] not in reg_dic:
            reg_dic[splitter1[3]] = 0
    # this is the setup loop, now we gotta go once more for instructions
    fileR.close()
    fileX = open("day8i.txt","r")
    for line in fileX:
        flag = False
        splitter = line.strip().split()
        print(splitter)
        operand = splitter[5]
        if operand == ">":
            if reg_dic[splitter[4]] > int(splitter[6]):
                flag = True
        elif operand == "<":
            if reg_dic[splitter[4]] < int(splitter[6]):
                flag = True
        elif operand == ">=":
            if reg_dic[splitter[4]] >= int(splitter[6]):
                flag = True
        elif operand == "<=":
            if reg_dic[splitter[4]] <= int(splitter[6]):
                flag = True
        elif operand == "==":
            if reg_dic[splitter[4]] == int(splitter[6]):
                flag = True
        elif operand == "!=":
            if reg_dic[splitter[4]] != int(splitter[6]):
                flag = True
        else:
            pass
            print("passed")
        print(flag)
        if flag:
            if splitter[1] == "inc":
                reg_dic[splitter[0]] += int(splitter[2])
                if reg_dic[splitter[0]] > maximum:
                    maximum = reg_dic[splitter[0]]
            if splitter[1] == "dec":
                reg_dic[splitter[0]] -= int(splitter[2])
    fileX.close()
    for key in reg_dic.keys():
        if reg_dic[key] > high:
            high = reg_dic[key]
    print("highest value at the end,:",high)
    print("highest value during the entire process:",maximum)

print(day8i())
