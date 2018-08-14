def d23i(part):
    # day 23 solution
    # don't usually do the below, but it's told
    # how many regsisters there are
    d = {"h":0}
    lineArray = []
    fileR = open("d23.txt","r")
    for line in fileR:
        splitter = line.strip("\n").split(" ")
        lineArray.append(splitter)
        if len(d) != 8:
            if splitter[1].isdigit() == False and splitter[1] not in d:
                d[splitter[1]] = 0
            if splitter[2].isdigit() == False and splitter[2] not in d:
                d[splitter[2]] = 0
    i = 0
    if part == 1:
        d['a'] = 0
    else:
        d['a'] = 1
    counter = 0
    while i < (len(lineArray)) and (i > -1):
        #print(d)
        # again,check operands to see if they
        # are infact registers or numbers
        if lineArray[i][0] != "jnz":
            #print(d)
            #print(lineArray[i][0])
            if lineArray[i][2].strip('-').isdigit():
                value = int(lineArray[i][2])
            else:
                value = d[lineArray[i][2]]
            if lineArray[i][0] == "set":
                d[lineArray[i][1]] = value
            elif lineArray[i][0] == "sub":
                d[lineArray[i][1]] -= value
            else:
                # it's mul
                d[lineArray[i][1]] *= value
                counter += 1 
        else:  
            if lineArray[i][1].isdigit():
                value = int(lineArray[i][1])
            else:
                value = d[lineArray[i][1]]
            if value != 0:
                i+= int(lineArray[i][2])
                continue
        i+=1
    if part == 1:
        print(counter)
    else:
        print(d['h'])

d23i(1)
d23i(2)

