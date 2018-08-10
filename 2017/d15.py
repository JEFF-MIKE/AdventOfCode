# should really tidy this code later on
# extreme brute force was used in the 
# writing of this code please be aware


class Generator:
    def __init__(self,factor,prev):
        self.prev = prev
        self.factor = factor
    


def d15i():
    judge = 0
    # day 15 of advent of code
    values = []
    fileR = open("d15.txt","r")
    for line in fileR:
        values.append(int(line.split(" ")[4]))
    genA = Generator(16807,values[0])
    genB = Generator(48271,values[1])
    modulator = 2147483647 # as in question
    for i in range(40000000):
        resultA = genA.prev * genA.factor
        resultB = genB.prev * genB.factor
        finalA = resultA % modulator
        finalB = resultB % modulator
        #print("genA",finalA)
        #print("genB",finalB)
        #print("------------")
        binA = format(finalA,'032b')[16:]
        binB = format(finalB,'032b')[16:]
        if binA == binB:
            judge+=1
        #print(binA)
        #print(binB)
        genA.prev = finalA
        genB.prev = finalB
    print(judge)


def d15ii():
    judge = 0
    # day 15 of advent of code
    values = []
    fileR = open("d15.txt","r")
    for line in fileR:
        values.append(int(line.split(" ")[4]))
    genA = Generator(16807,values[0])
    genB = Generator(48271,values[1])
    modulator = 2147483647 # as in question
    for i in range(5000000):
        # in this part, the generators run independently
        while True:
            # generator for a
            resultA = genA.prev * genA.factor
            finalA = resultA % modulator
            if finalA % 4 == 0:
                a = finalA
                break
            else:
                genA.prev = finalA
                continue
        while True:
            resultB = genB.prev * genB.factor
            finalB = resultB % modulator
            if finalB % 8 == 0:
                b = finalB
                break
            else:
                genB.prev = finalB
                continue
        #print("genA",a)
        #print("genB",b)
        #print("------------")
        binA = format(a,'032b')[16:]
        binB = format(b,'032b')[16:]
        if binA == binB:
            judge+=1
        #print(binA)
        #print(binB)
        genA.prev = a
        genB.prev = b
    print(judge)
d15i()
d15ii()