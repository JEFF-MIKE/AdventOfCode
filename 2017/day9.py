class Customstack:
    def __init__(self):
        self._length = 0 # empty
        self._storage = []
        self._gStart = None
    def getTop(self):
        return self._storage[0]
    def getLength(self):
        return self._length
    def pop(self):
        if self.isEmpty():
            return None
        element = self._storage[-1]
        self._storage = self._storage[:-1]
        self._length -= 1
        return element
    def top(self):
        if self.isEmpty():
            return None
        return self._storage[-1]
    def isEmpty(self):
        if self._length == 0:
            return True
        return False
    def push(self,item):
        self._storage.append(item)
        self._length += 1
    def __str__(self):
        print("-----------------------\n")
        mloc_string = []
        print("Stack contents: (First item here is top of stack)")
        for i in range((len(self._storage)-1),-1,-1):
            mloc_string.append(str(self._storage[i]))
        return "\n".join(mloc_string) + "\n-----------------------"
        
def day9i():
    fileR = open("day9.txt","r")
    string = fileR.read()
    garbage = 0
    totalSum = 0
    factor = 0
    lst = [_ for _ in string]
    s = Customstack()
    i = 0
    while i < len(lst): 
        if lst[i] == "!":
            i+=2
            continue
        elif lst[i] == "<":
            n = True
            i += 1
            while n:
                if lst[i] not in [">","!"]:
                    garbage += 1
                if lst[i] == "!":
                    i += 2
                elif lst[i] == ">":
                    n = False
                else:
                     i += 1
        elif lst[i] == "{":
            factor += 1
            s.push(lst[i])
        elif lst[i] == "}":
            b = True
            while b:
                tempStack = Customstack()
                #print("doing stuff")
                if s.top() == "<":
                    tempStack.push(s.top())
                    s.pop()
                if s.top() == "{":
                    totalSum += factor
                    s.pop()
                    factor -= 1
                    for _ in range(tempStack.getLength()):
                        s.push(tempStack.top())
                        tempStack.pop()
                    b = False
                else:
                    s.pop()
        else:
            s.push(lst[i])
        i += 1
    print("Total score of day9.txt is: ",totalSum)
    print("Amount of collected garbage is: ",garbage)
    fileR.close()
    return

print(day9i())
