'''
record each register in a dictionary
for fast access time
'''

def d18i():
    lineArray = [] # an array of our instructions
    d = {} # register dictionary
    fileR = open("d18.txt","r")
    for line in fileR:
        splitter = line.strip('\n').split(" ")
        if splitter[1] not in d:
            d[splitter[1]] = 0
        lineArray.append(splitter)
    flag = True
    frequency = 0
    index = 0 # how we move through our instructions
    while flag:
        if len(lineArray[index]) == 3:
            # should do an initial check 
            # check whether it only contains digits or not
            if lineArray[index][2].strip('-').isdigit(): 
                value = int(lineArray[index][2])
            else :
                # it's a register
                value = d[lineArray[index][2]]                        
            if lineArray[index][0] == "set":
                d[lineArray[index][1]] = value
            elif lineArray[index][0] == "add":
                d[lineArray[index][1]] += value
            elif lineArray[index][0] == "mul":
                d[lineArray[index][1]] *= value
            elif lineArray[index][0] == "mod":
                d[lineArray[index][1]] %= value
            else:
                # it's jgz, need to do a check
                # whether 2nd part is register or int here
                if lineArray[index][1].isdigit():
                    k = int(lineArray[index][1])
                else:
                    k = d[lineArray[index][1]]
                if k > 0:
                    index += value
                    continue
            index+=1
        else:
            if lineArray[index][0] == "snd":
                frequency = d[lineArray[index][1]]
            else: # it's rcv
                if d[lineArray[index][1]] > 0:
                    print("Recovering sound!")
                    print(frequency)
                    flag = False
            index+=1

d18i()



