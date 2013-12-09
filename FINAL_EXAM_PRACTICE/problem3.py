#   Our way
def f(b):
    return lambda a: a / (b * b)

a = [f(x) for x in range(10)]

#   we CAN DO THIS BETTER..
a = (lambda x: x / (b*b) for b in range(10))
print b
