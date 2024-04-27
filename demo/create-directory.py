# Python3.
#
# Arguments are text filenames
# that got copied into the coypu serving
# directory.
#
# Write 16 64-byte records with links
# to the text files followed by ^Z dummy
# records.

import sys

args = sys.argv[1:]
for filename in args:
	s = "31------------" + filename
	s += (64 - len(s)) * " "
	print(s, end='')

# There must be 16 64-byte records.
# Fill the remainder with control-Zs.
for i in range(16 - len(args)):
	print(64 * '\032', end='')
