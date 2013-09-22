#   Problem 9
import sys
inputStr = sys.argv[1]

for i in range(0, len(inputStr) + 1):
	print inputStr[:i]