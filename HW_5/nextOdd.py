def nextOdd():
	state = {'x': -1}
	return lambda x: state['x'] += 2
