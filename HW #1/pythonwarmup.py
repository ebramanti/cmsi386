#   HW #1
#   pythonwarmup.py
#	Containing Functions for Questions 1-7
#	Edward Bramanti & Zane Kansil 

from random import shuffle

#   Problem 1
def change(cents):
    if cents < 0: raise ValueError("Not an int > 0.")
    usCoinDenominations = [25,10,5,1]
    results = []
    remaining = cents
    for s in usCoinDenominations:
        partialSolution = divmod(remaining, s)
        results.append(partialSolution[0])
        remaining = partialSolution[1]
    
    return tuple(results)

#   Problem 2
def strip_vowels(s):
    return s.translate(None, "aeiouAEIOU")

#   Problem 3
def scramble(s):
    word = list(s)
    shuffle(word)
    return ''.join(word)

#   Problem 4
def powers_of_two(max):
    power, current = 0, 1
    while True:
        current = 2 ** power
        if current > max:
            break;
        yield current
        power += 1

#   Problem 5
def powers(base, max):
    power, current = 0, 1
    while True:
        current = base ** power
        if current > max:
            break;
        yield current
        power += 1

#   Problem 6
def interleave(a, b):
    return [y for x in zip(a,b) for y in x] + b[len(a):] + a[len(b):]

#   Problem 7
def stutter(list1):
	result = []
	degreeOfStutter = 2
	for i in list1:
		for j in range(degreeOfStutter):
			result.append(i)
	return result



