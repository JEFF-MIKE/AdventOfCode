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

#print(day4i())

def day4ii():
    fileR = open("input4.txt","r")
    total = 0
    for line in fileR:
        counter = True
        splitter = line.split(" ")
        splitter[-1] = splitter[-1].strip()
        print(splitter)
        for i in range(len(splitter)):
            for j in range(i+1,len(splitter)):
                part1 = [x for x in splitter[i]]
                part2 = [_ for _ in splitter[j]]
                part1.sort()
                part2.sort()
                if part1 == part2:
                    counter = False               
        if counter:
            total += 1
    return total

print(day4ii())
                    

                    
