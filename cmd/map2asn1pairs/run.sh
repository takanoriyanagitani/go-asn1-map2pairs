#!/bin/sh

input=./input.json
output=./output.asn1.der.dat

jsonmap(){
	printf '{
		"helo":"wrld",
		"mount":"fuji"
	}'
}

genjsonmap(){
	jsonmap > "${input}"
}

test -f "${input}" || genjsonmap

cat "${input}" |
	./map2asn1pairs |
	dd \
		if=/dev/stdin \
		of="${output}" \
		bs=1048576 \
		status=none

ls -l "${input}" "${output}"

cat "${output}" |
	python3 sample.py |
	jq -c
