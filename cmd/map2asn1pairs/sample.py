import asn1tools
import sys
import json

pairs = asn1tools.compile_files("./pairs.asn")
encoded = sys.stdin.buffer.read()
decoded = pairs.decode("RawPairs", encoded)
mapd = map(json.dumps, decoded)
prints = map(print, mapd)
sum(1 for _ in prints)
