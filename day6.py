def day6():
    counter = 0
    flag = True
    fileR = open('day6input.txt',"r")
    states = []
    lst = None
    for line in fileR: # it's one line eggs dee
        splitter = line.split("\t")
        splitter[-1] = splitter[-1].strip()
        lst = [int(x) for x in splitter]
    '''
    for i in range(len(lst)):
        state.append(lst[i])
    while True:
        #print("lst",lst)
        #print("state",state)
        current = 0
        maxi = 0
        indexer = 0
        for i in range(len(lst)):
            if lst[i] > maxi:
                maxi = lst[i]
                current = i
        #print("Maxi is: ",maxi,"at index: ",current)
        indexer = current + 1
        if indexer == len(lst):
            indexer = 0
        lst[current] = 0
        #print("current",current)
        while maxi != 0:
            lst[indexer] += 1
            indexer += 1
            if indexer == len(lst):
                indexer = 0
            maxi -= 1
        counter += 1
        if lst == state:
            break
    return counter
    '''
    while lst not in states:
        states.append(lst[:])
        print(states)
        m = max(lst)
        i = lst.index(m)
        lst[i] = 0
        while m:
            i =(i+1)%len(lst)
            lst[i] += 1
            m-= 1
        counter += 1
    return counter
print(day6())
                
                    
