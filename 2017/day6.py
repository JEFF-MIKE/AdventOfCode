def day6():
    listContainer = []
    answerPart=0
    firstLoop = 0
    listInQuestion = []
    # just iterate through the list and
    # then balance as required
    stateCounter = 0 # answer to question
    fileR = open("day6input.txt","r")
    for line in fileR:
        x = [int(y) for y in line.split('\t')]
    listContainer.append(x[:])
    maxi = 0 # max int
    workingIndex = 0 # iterator
    while True:
        stateCounter+=1
        # have to reset maxi to zero here
        maxi = 0
        workingIndex = 0
        #print("performing while true")
        # find the max value first
        for i in range(len(x)):
            if x[i] > maxi:
                maxi = x[i]
                workingIndex = i
        # now balance as required
        value = maxi
        x[workingIndex] = 0
        workingIndex+=1 # goes to next element
        while value != 0:
            # reset workingIndex to 0
            # if it is out of range
            if workingIndex == len(x):
                workingIndex = 0
            x[workingIndex] += 1
            value -=1
            workingIndex+=1
        # now we do a check on the arrays to
        # see if their states are the same again or not
        #print(x)
        for k in listContainer:
            if x == k:
                if answerPart == 0:
                    print("The answer to the first part is: ",stateCounter)
                    firstLoop = stateCounter
                    answerPart+=1
                    listContainer = [] # reset the list
                    break
                else:
                    print("The answer to the second part is:",(stateCounter-firstLoop))
                    return None
        listContainer.append(x[:])
day6()