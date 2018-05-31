class Plumber:
    def __init__(self,element):
        self._element = element
        self._connectedTo = []
        self._status = False
    @property
    def element(self):
        return self._element
    @element.setter
    def element(self,element):
        self._element = element
    @property
    def connectedTo(self):
        return self._connectedTo
    def addChild(self,adder):
        self._connectedTo.append(adder)
        
    def __eq__(self,other):
        if self._element == other._element and self._connectedTo == other._connectedTo:
            return True
        return False
    @property
    def status(self):
        return self._status
    @status.setter
    def status(self,status):
        self._status = status
    def __str__(self):
        child = ""
        for i in range(len(self._connectedTo)):
            child += str(self._connectedTo[i].element) + "\n"
        return "Element: %s\nConnected to: %s" % (self._element,child)
    
def day12():
    oLst = [] # array for objects
    fileR = open("day12.txt","r")
    for line in fileR:
        splitter = line.strip().split(" ")
        p = Plumber(int(splitter[0]))
        for i in range(2,len(splitter)):
            splitter[i] = splitter[i].strip(",") # remove the commas from the children
            p.addChild(int(splitter[i]))
        oLst.append(p)
    fileR.close()
    # Re-attach the children together
    for m in range(len(oLst)): # this is a sorted array btw
        for n in range(len(oLst[m].connectedTo)): # children of the object as an int
            oLst[m]._connectedTo[n] = oLst[oLst[m]._connectedTo[n]]
    # All hooked up, now time for the answer. I will place the object names in an array and get their length
    ansLst = []
    oLst[0].status = True
    ansLst.append(oLst[0])
    indexer = 0
    while indexer < len(ansLst):
        for g in range(len(ansLst[indexer]._connectedTo)):
            if (ansLst[indexer]._connectedTo[g] not in ansLst) and ansLst[indexer]._connectedTo[g].status == False:
                ansLst[indexer]._connectedTo[g].status = True
                ansLst.append(ansLst[indexer]._connectedTo[g])
        indexer += 1
    return len(ansLst)
            
#print(day12())

def day12i():
    oLst = [] # array for objects
    fileR = open("day12.txt","r")
    for line in fileR:
        splitter = line.strip().split(" ")
        p = Plumber(int(splitter[0]))
        for i in range(2,len(splitter)):
            splitter[i] = splitter[i].strip(",") # remove the commas from the children
            p.addChild(int(splitter[i]))
        oLst.append(p)
    fileR.close()
    # Re-attach the children together
    for m in range(len(oLst)): # this is a sorted array btw
        for n in range(len(oLst[m].connectedTo)): # children of the object as an int
            oLst[m]._connectedTo[n] = oLst[oLst[m]._connectedTo[n]]
    ansLst = [] # for part 2, will contain the object lists "container"
    for g in range(len(oLst)):
        if oLst[g].status == False: # is not in a group
            container = []
            container.append(oLst[g])
            indexer = 0
            while indexer < len(container):
                for m in range(len(container[indexer]._connectedTo)):
                    if (container[indexer]._connectedTo[m] not in container) and container[indexer]._connectedTo[m].status == False:
                        container[indexer]._connectedTo[m].status = True
                        container.append(container[indexer]._connectedTo[m])
                indexer += 1
            #exit loop here, now append container to the ansLst
            ansLst.append(container)
        else:
            continue # is true, is  already in a group
            
    return len(ansLst)
print(day12i())
