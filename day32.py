import math
def day3(n):
    # manhattan distance
    # will be our square number
    length = math.ceil(math.sqrt(n))  + 2 # I add two here to provide an outline for append
    container = [None for _ in range(length)]
    grid = [container for m in range(length)]
    half = ((len(grid) // 2) + 1) # midpoint for this
    grid[half][half] = 1
    d1 = half # d1 and d2 are used to 
    d2 = half # navigate our grid
    # d1 will go up/down,d2 will go left/right on our grid
    #print(grid[d1][d2])
    value = 2 # this is what we'll add to our 2d table 
    d = {"r":1,"u":1,"l":2,"d":2} #right left up down. We add 2 everytime we're done with this key
    while value <= n:
        right = d["r"]
        while right != 0:
            d2 += 1
            grid[d1][d2] = value
            value += 1
            if value > n:
                break
            right -= 1
        if value > n:
            break
        d["r"] += 2
        up = d["u"]
        while up != 0:
            d1 += 1
            grid[d1][d2] = value
            value += 1
            if value > n:
                break
            up -= 1
        if value > n:
            break
        d["u"] += 2
        left = d["l"]
        while left != 0:
            d2 -= 1
            grid[d1][2] = value
            value += 1
            if value > n:
                break
            left -= 1
        if value > n:
            break
        d["l"] += 2
        down = d["d"]
        while down != 0:
            d1 -= 1
            grid[d1][d2] = value
            value += 1
            if value > n:
                break
            down -= 1
        if value > n:
            break
        d["d"] += 2
    print("Math answer: ",abs(half - d1) + abs(half - d2))
#print(day3(312051))

def day3i(n):
    # manhattan distance
    # will be our square number
    length = math.ceil(math.sqrt(n))  + 2 # I add two here to provide an outline for append
    #container = [0 for _ in range(length)]
    grid = [[0 for _ in range(length)] for m in range(length)]
    half = ((len(grid) // 2) + 1) # midpoint for this
    grid[half][half] = 1
    d1 = half # d1 and d2 are used to 
    d2 = half # navigate our grid
    # d1 will go up/down,d2 will go left/right on our grid
    #print(grid[d1][d2])
    d = {"r":1,"u":1,"l":2,"d":2} #right left up down. We add 2 everytime we're done with this key
    while True:
        right = d["r"]
        while right != 0:
            d2 += 1
            r = grid[d1][d2 + 1]
            tr = grid[d1-1][d2+1]
            t=grid[d1-1][d2]
            tl=grid[d1-1][d2-1]
            l=grid[d1][d2-1]
            bl=grid[d1+1][d2-1]
            b=grid[d1+1][d2]
            br=grid[d1+1][d2+1]
            grid[d1][d2] = r+tr+t+tl+l+bl+b+br
            print("Value in grid", grid[d1][d2])
            if grid[d1][d2] > n:
                break
            right -= 1
        if grid[d1][d2] > n:
            break
        d["r"] += 2
        up = d["u"]
        while up != 0:
            d1 += 1
            r = grid[d1][d2 + 1]
            tr = grid[d1-1][d2+1]
            t=grid[d1-1][d2]
            tl=grid[d1-1][d2-1]
            l=grid[d1][d2-1]
            bl=grid[d1+1][d2-1]
            b=grid[d1+1][d2]
            br=grid[d1+1][d2+1]
            grid[d1][d2] = r+tr+t+tl+l+bl+b+br
            print("Value in grid", grid[d1][d2])
            if grid[d1][d2] > n:
                break
            up -= 1
        if grid[d1][d2] > n:
            break
        d["u"] += 2
        left = d["l"]
        while left != 0:
            value = 0
            d2 -= 1
            r = grid[d1][d2 + 1]
            tr = grid[d1-1][d2+1]
            t=grid[d1-1][d2]
            tl=grid[d1-1][d2-1]
            l=grid[d1][d2-1]
            bl=grid[d1+1][d2-1]
            b=grid[d1+1][d2]
            br=grid[d1+1][d2+1]
            grid[d1][d2] = r+tr+t+tl+l+bl+b+br
            print("Value in grid", grid[d1][d2])
            if grid[d1][d2] > n:
                break
            left -= 1
        if grid[d1][d2] > n:
            break
        d["l"] += 2
        down = d["d"]
        while down != 0:
            d1 -= 1
            r = grid[d1][d2 + 1]
            tr = grid[d1-1][d2+1]
            t=grid[d1-1][d2]
            tl=grid[d1-1][d2-1]
            l=grid[d1][d2-1]
            bl=grid[d1+1][d2-1]
            b=grid[d1+1][d2]
            br=grid[d1+1][d2+1]
            grid[d1][d2] = r+tr+t+tl+l+bl+b+br
            print("Value in grid", grid[d1][d2])
            if grid[d1][d2] > n:
                break
            down -= 1
        if grid[d1][d2] > n:
            break
        d["d"] += 2
    return grid[d1][d2]
print(day3i(312051))

# Please note: These are the brute force methods. The grids are generated, and
# then the answers are received from said generated grid

