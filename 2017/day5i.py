def day5i():
    counter = 0
    lst = []
    fileR = open("input5.txt","r")
    for line in fileR:
        line = line.strip()
        lst.append(int(line))
    # array done
    print(lst)
    indexer = 0
    while indexer < len(lst):
        # print(indexer)
        prev = indexer
        indexer += lst[prev]
        lst[prev] += 1
        counter += 1
        if indexer >= len(lst):
            break
    return (indexer,counter)
#print(day5i()) 

def day5ii():
    counter = 0
    lst = []
    fileR = open("input5.txt","r")
    for line in fileR:
        line = line.strip()
        lst.append(int(line))
    # array done
    print(lst)
    indexer = 0
    while indexer < len(lst):
        print(indexer)
        prev = indexer
        indexer += lst[prev]
        if lst[prev] >= 3:
            lst[prev] -= 1
        else:
            lst[prev] += 1
        counter += 1
        if indexer >= len(lst):
            break
    return (indexer,counter)
print(day5ii())        
