def nextOdd():
	state = {'x': -1}
	def f(): state['x'] += 2; return state['x']
	return f
	# NOT A MODULEEEEE WHYYYYYYYYYYYYY