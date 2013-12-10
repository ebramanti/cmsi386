#   Our way
def f(b):
    return lambda a: a / (b * b)

a = [f(x) for x in range(10)]
