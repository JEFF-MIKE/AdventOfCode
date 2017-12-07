def day4i():
    fileR = open("input4.txt","r")
    total = 0
    for line in fileR:
        counter = True
        splitter = line.split(" ")
        splitter[-1] = splitter[-1].strip()
        print(splitter)
        for i in range(len(splitter)):
            for j in range(i + 1,len(splitter)):
                if splitter[i] == splitter[j]:
                    counter = False
                    break
        if counter:
            total += 1
    return total

print(day4i())