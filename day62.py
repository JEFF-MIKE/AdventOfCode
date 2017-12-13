def day6():
    lst = [10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6]
    s = [10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6]
    container = []
    counter = 0
    '''
    fileR = open("day6input.txt","r")
    for line in fileR:
        splitter = line.strip().split()
        lst = [int(_) for _ in splitter]
    for item in lst:
        s.append(item)
    '''
    container.append([10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6])
    flag = True
    while flag:
        maxi = max(s)
        i = s.index(maxi)
        s[i] = 0
        i += 1
        if i == len(s):
            i = 0 # reset the index here
        while maxi != 0:
            #print(maxi)
            s[i] += 1
            i += 1
            if i == len(s):
                i = 0 # reset the index
            maxi -= 1
        counter += 1
        meme = 0
        if s in container:
            flag = False
            counter += 1
    fileR.close()
    print(counter)
day6()
