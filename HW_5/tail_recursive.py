def findMinValue(a, i = 0, sofar = None):
    if i == len(a) - 1:
        return a[i] if a[i] < sofar else sofar
    elif sofar == None:
        sofar = a[i]
    elif a[i] < sofar:
        sofar = a[i]
    
    return findMinValue(a, i + 1, sofar)
