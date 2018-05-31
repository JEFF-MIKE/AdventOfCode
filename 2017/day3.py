import math
def day3i(num):
    box = [] # holds all our other lists
    side= math.ceil(math.sqrt(num))
    mid = math.ceil(side / 2)
    print(mid)
    for i in range(side):
        lst = [None for _ in range(side)]
        box.append(lst)
    print("Done")
    print(len(box[0]))
    print(len(box))
day3i(312051)

        
    
