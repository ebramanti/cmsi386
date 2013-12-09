def counter():
    n = {'value': 0}
    def increment():
        n['value'] = n['value'] + 1
        print n['value']
    def decrement():
        n['value'] = n['value'] - 1
        print n['value']
    return (increment, decrement)