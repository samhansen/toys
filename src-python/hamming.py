#!/usr/bin/python
# Compute the hamming weight for a 32 bit integer.

from sys import argv

def main(argv):
	arg = int(argv[1])

	# handle int32 overflow
	if arg > 0xFFFFFFFF:
		raise ValueError

	# ultimately we need 1/2 of the total bits masked off, this example assumes
	# 32 bits
	m1  = 0x55555555	# 0b01...
	m2  = 0x33333333	# 0b0011...
	m4  = 0x0f0f0f0f	# 0b00001111...
	m8  = 0x00ff00ff	# 0b0000000011111111...
	m16 = 0x0000ffff	# 0b00000000000000001111111111111111...

	# the operations over 32 bits
	x = arg
	x = (x & m1)  + ((x >> 1)  & m1)
	x = (x & m2)  + ((x >> 2)  & m2)
	x = (x & m4)  + ((x >> 4)  & m4)
	x = (x & m8)  + ((x >> 8)  & m8)
	x = (x & m16) + ((x >> 16) & m16)

	print 'hamming weight of %d is %d' % (arg, x)


if __name__ == '__main__':
	try:
		main(argv)
	except ValueError:
		print 'Error, value larger than 32 bits!!'
	except IndexError:
		print 'Usage: hamming.py int32'
