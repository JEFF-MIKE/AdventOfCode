class treeNode:
    def __init__(self,name,weight):
        self._name = name
        self._parent = None
        self._children = []
        self._weight = weight
    @property
    def parent(self):
        return self._parent
    @parent.setter
    def parent(self,parent):
        self._parent = parent
    @property
    def children(self):
        return self._children
    @property
    def name(self):
        return self._name
    @name.setter
    def name(self,name):
        self._name = name
    @property
    def weight(self):
        return self._weight
    @weight.setter
    def weight(self,weight):
        self._weight = weight
    def getSum(self):
        # returns the sum of this node, and it's children for use in part 2
        g1 = 0
        for _ in range(len(self._children)):
            g1 += self._children[_].weight
            g1 += self._weight
        return g1
    def addChild(self,child):
        self._children.append(child)
        return
    def __str__(self):
        return "Parent: %s\nChildren: %s\nweight:%i" % (self._parent,self._children,self._weight)
        

def day7i():
    fileResult = []
    mini = 10000
    current = 0
    arrow = []
    oLst = [] # list to hold objects
    leaves = []
    fileR = open("day7i.txt","r")
    for line in fileR:
        '''
        splitter = line.split(" ")
        splitter[-1] = splitter[-1].strip()
        splitter[1] = splitter[1].strip("()")
        '''
        splitter = line.strip().split(" ")
        '''
        container = []
        container.append(splitter[0])
        container.append(int(splitter[1]))
        fileResult.append(container)
        '''
        if "->" in splitter:
            arrow.append(splitter) # we only care about nodes with children (not leaves)
            p = treeNode(splitter[0])
            childLst = splitter[3:]
            #print(childLst)
            for j in range(len(childLst)):
                #print(childLst[j])
                p.addChild(childLst[j].strip(","))
            oLst.append(p)
        else: # it's a child node
            leaves.append(splitter[0])
    for k in range(len(oLst)): # lst loop
        for m in range(len(oLst[k].children)):
            for g in range(len(oLst)):
                #print(oLst[k].children[m])
                #print(oLst[g].name)
                if oLst[k].children[m] == oLst[g].name:
                    oLst[g].parent = oLst[k]
                    #break # exit the g loop. go to the next child
    for i in range(len(oLst)):
        if oLst[i].parent == None:
            print
            print(oLst[i].name)
#print(day7i())

def day7ii():
    weights = []
    fileResult = []
    current = 0
    arrow = []
    oLst = [] # list to hold objects
    leaves = []
    fileR = open("day7i.txt","r")
    for line in fileR:
        '''
        splitter = line.split(" ")
        splitter[-1] = splitter[-1].strip()
        splitter[1] = splitter[1].strip("()")
        '''
        splitter = line.strip().split(" ")
        splitter[1] = splitter[1].strip("()")
        splitter[1] = int(splitter[1])
        '''
        container = []
        container.append(splitter[0])
        container.append(int(splitter[1]))
        fileResult.append(container)
        '''
        if "->" in splitter:
            arrow.append(splitter) # we only care about nodes with children (not leaves)
            p = treeNode(splitter[0],splitter[1])
            childLst = splitter[3:]
            #print(childLst)
            for j in range(len(childLst)):
                #print(childLst[j])
                p.addChild(childLst[j].strip(","))
            oLst.append(p)
        else: # it's a child node
            leaf = treeNode(splitter[0],splitter[1])
            leaves.append(leaf)
    for k in range(len(oLst)): # lst loop
        for m in range(len(oLst[k].children)):
            for g in range(len(oLst)):
                #print(oLst[k].children[m])
                #print(oLst[g].name)
                if oLst[k].children[m] == oLst[g].name:
                    oLst[g].parent = oLst[k]
                    oLst[k].children[m] = oLst[g]
                    #break # exit the g loop. go to the next child
            for _ in range(len(leaves)):
                if oLst[k].children[m] == leaves[_].name:
                    leaves[_].parent = oLst[k]
                    oLst[k].children[m] = leaves[_]
    '''  
    for i in range(len(oLst)):
        if oLst[i].parent == None:
            print
            print(oLst[i].name)
    for h in range(len(leaves)):
        print(leaves[h].parent)
    '''
    # Tree is built properly, now for meme time
    for i in range(len(oLst)):
        if oLst[i].parent == None:
            root = oLst[i]
    flag = True
    while flag:
        # finish this part. I want to check that a the sum of the weight of a node and its children are the same on each level
        
    
            
print(day7ii())
