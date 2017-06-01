#!/usr/bin/env python
# CodeEval: Message Decoding (hard).
import sys


def IndexFromKey(key):
  """Compute the header index from the given key string."""
  # First compute the base index of the 0-value of a key whose length is
  # the same as the given key.
  length = len(key)
  base = 2 ** length - (length + 1)

  # Return the offset from the computed base.
  return base + int(key, base=2)


def ProcessSegment(header, key_len, data):
  """Process a segment of key_len beginning at data[0]."""
  terminal = '1' * key_len
  while True:
    key, data = data[:key_len], data[key_len:]
    if key == terminal:
      break
    sys.stdout.write(header[IndexFromKey(key)])

  return data


def ProcessLine(line):
  """Process the given line."""
  part = min(line.find('0'), line.find('1'))
  header, data = line[:part], line[part:]

  while True:
    key_len, data = int(data[:3], base=2), data[3:]
    if key_len == 0:
      break
    data = ProcessSegment(header, key_len, data)
  sys.stdout.write('\n')


def main(argv):
  """Program entry."""
  with open(argv[1]) as f:
    for line in f.readlines():
      ProcessLine(line.strip())


if __name__ == '__main__':
  main(sys.argv)
