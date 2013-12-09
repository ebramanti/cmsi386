def withoutEvens(a):
    def helper(a, i, result):
        result.append(a[i])
        return result if i + 2 >= len(a) else helper(a, i + 2, result)
    return helper(a, 1, [])


print withoutEvens([0,1,2,3,4,5,6])
print withoutEvens([4,6,True,None,2.3,"xyz"])
print withoutEvens(['h', 'e', 'y'])