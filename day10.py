def day10():
    k = []
    lst = [int(x) for x in range(256)]
    head = 0
    skip_size =0
    fileR = open("day10.txt","r")
    for line in fileR:
        k = [int(_) for _ in line.strip().split(",")]
    for j in range(len(k)):
        container = []
        prev = head
        for _ in range(k[j]):
            container.append(lst[head])
            head += 1
            if head == len(lst):
                head = 0
        container = container[::-1]
        for m in range(len(container)):
            lst[prev] = container[m]
            prev += 1
            if prev == len(lst):
                prev = 0
        head += skip_size
        skip_size += 1
        if head >= len(lst):
            head = head % len(lst)
    fileR.close()
    return lst[0] *lst[1]
#print(day10())

def day10ii():
    # Need to finish, look at the 2nd part!
    answer_lst = []
    k = []
    lst = [int(x) for x in range(256)]
    fileR = open("day10.txt","r")
    string = "1,2,3"
    k = [ord(char) for char in string]
    default = [17,31,73,47,23]
    for item in default:
        k.append(item)
    print(k)
    # done first part, now second part (loop 64)
    for j in range(len(k)):
        head = 0
        skip_size =0
        for g in range(64):
            container = []
            prev = head
            for _ in range(k[j]):
                container.append(lst[head])
                head += 1
                if head == len(lst):
                    head = 0
            container = container[::-1]
            for m in range(len(container)):
                lst[prev] = container[m]
                prev += 1
                if prev == len(lst):
                    prev = 0
            head += skip_size
            skip_size += 1
            if head >= len(lst):
                head = head % len(lst)
    for g in range(0,256,16):
        container = []
        answer = lst[g]
        for u in range(1,16):
            container.append(lst[g+u])
            answer = answer ^ container[-1]
        answer_lst.append(answer)
    print(len(answer_lst))
    print(answer_lst[0])
    for n in range(len(answer_lst)):
        answer_lst[n] = hex(answer_lst[n])[2:]
    fileR.close()
    return "".join(answer_lst)

print(day10ii())
    
