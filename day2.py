def day2i():
    # checksum
    checksum = 0
    fileR = open("day2i.txt","r")
    fileZ = open("day2a.txt","r")
    for line in fileR:
        print(line)
        splitter = line.split('\t')
        splitter[-1] = splitter[-1].strip()
        print(splitter)
        for p in range(len(splitter)):
            splitter[p] = int(splitter[p])
        maxi = 0
        mini = splitter[0]
        for j in range(len(splitter)):
            if splitter[j] > maxi:
                maxi = splitter[j]
            if splitter[j] < mini:
                mini = splitter[j]
        difference = maxi - mini
        print("Maximum is: ",maxi)
        print("Minimum is: ",mini)
        print("Max() is", max(splitter))
        print("Min() is", min(splitter))
        print("Difference:", difference)
        checksum += difference
        print("CheckSum is: ",checksum)
        if len(splitter) != 16:
            print("Error! Length of splitter incorrect")
    return checksum
            
# print(day2i())

def day2ii():
    # Does the numbers divide evenly? if so, checksum +=
    testlst = []
    checksum = 0
    fileS = open("day2i.txt","r")
    for line in fileS:
        splitter = line.split('\t')
        splitter[-1] = splitter[-1].strip()
        for g in range(len(splitter)):
            splitter[g] = int(splitter[g])
        print(splitter)
        result = 0
        for j in range(len(splitter)):
            for x in range(len(splitter)):
                if splitter[j] % splitter[x] == 0 and (j != x):
                    print("j is ",splitter[j])
                    print("x is ",splitter[x])
                    result = splitter[j] / splitter[x]
                    print("result is ", result)
                    checksum += result
                    testlst.append(result)
                    break # only one even , therefore we get outta there!
    for h in range(len(testlst)):
        print((h+1),testlst[h])
    print(sum(testlst))
    return checksum
print(day2ii())
