def d16i():
    disk = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p']
    #disk=['a','b','c','d','e']
    fileR = open("d16.txt","r")
    lst = fileR.read().strip('\'').split(",")
   # print(lst[0][0])
    for i in range(len(lst)):
        if lst[i][0] == 's':
            k = int(lst[i].strip('s')) - 1
            newLst = []
            for p in range(k,-1,-1):
                element = disk[(len(disk)-1)-p]
                newLst.append(element)
            for p in range(((len(disk))-len(newLst))):
                newLst.append(disk[p])
            disk = newLst
            continue
        elif lst[i][0] == 'x':
            temp = lst[i].strip('x').split('/')
           # print(temp)
            pos1 = int(temp[0])
            pos2 = int(temp[1])
            #print(pos1,pos2)
            #print(disk)
            disk[pos1],disk[pos2] = disk[pos2],disk[pos1]
        else:
            # it's the p instruction, which
            # swaps by names of files as opposed to
            # indexes like above
            pos1 = disk.index(lst[i][1])
            pos2 = disk.index(lst[i][3])
            disk[pos1],disk[pos2] = disk[pos2],disk[pos1]
    print(disk)
    print("".join(disk))

def d16ii():
    disk = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p']
    #disk=['a','b','c','d','e']
    fileR = open("d16.txt","r")
    lst = fileR.read().strip('\'').split(",")
    # print(lst[0][0])
    for q in range(1000000000):
        for i in range(len(lst)):
            if lst[i][0] == 's':
                k = int(lst[i].strip('s')) - 1
                newLst = []
                for p in range(k,-1,-1):
                    element = disk[(len(disk)-1)-p]
                    newLst.append(element)
                for p in range(((len(disk))-len(newLst))):
                    newLst.append(disk[p])
                disk = newLst
                continue
            elif lst[i][0] == 'x':
                temp = lst[i].strip('x').split('/')
            # print(temp)
                pos1 = int(temp[0])
                pos2 = int(temp[1])
                #print(pos1,pos2)
                #print(disk)
                disk[pos1],disk[pos2] = disk[pos2],disk[pos1]
            else:
                # it's the p instruction, which
                # swaps by names of files as opposed to
                # indexes like above
                pos1 = disk.index(lst[i][1])
                pos2 = disk.index(lst[i][3])
                disk[pos1],disk[pos2] = disk[pos2],disk[pos1]
    print(disk)
    print("".join(disk))
        


#d16i()
d16ii()