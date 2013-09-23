#   Problem 8
import sys, string

lineCounter = 0
textFile = sys.argv[1]
with open(textFile) as f:
    for line in f:
        line = line.strip()
        if line and not line.startswith("#"):
            lineCounter += 1
print lineCounter