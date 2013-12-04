a = [1,2,3]
b = [21,7,12,2]
c = [1, 12, -1, 0]
d = [1, 20, 500, 6, -500, 40, -500, 67]


def findMinValue(a, i = 0, sofar = None):
    if i == len(a) - 1:
        return a[i] if a[i] < sofar else sofar
    elif sofar == None:
        sofar = a[i]
    elif a[i] < sofar:
        sofar = a[i]
    
    return findMinValue(a, i + 1, sofar)

print findMinValue(a)
print findMinValue(b)
print findMinValue(c)
print findMinValue(d)



