package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	mp "github.com/takanoriyanagitani/go-asn1-map2pairs"
	. "github.com/takanoriyanagitani/go-asn1-map2pairs/util"
)

var envValByKey func(string) IO[string] = Lift(
	func(key string) (string, error) {
		val, found := os.LookupEnv(key)
		switch found {
		case true:
			return val, nil
		default:
			return "", fmt.Errorf("env var %s unknown", key)
		}
	},
)

var mapSizeLimit IO[int64] = Bind(
	envValByKey("ENV_MAP_FILESIZE_LIMIT"),
	Lift(func(s string) (int64, error) {
		i, e := strconv.Atoi(s)
		return int64(i), e
	}),
).Or(Of(int64(1048576)))

func reader2bytesLimited(limit int64) func(io.Reader) IO[[]byte] {
	return Lift(func(rdr io.Reader) ([]byte, error) {
		limited := &io.LimitedReader{
			R: rdr,
			N: limit,
		}
		var buf bytes.Buffer
		_, e := io.Copy(&buf, limited)
		return buf.Bytes(), e
	})
}

func stdin2bytesLimited(limit int64) IO[[]byte] {
	var rdr2bytes func(io.Reader) IO[[]byte] = reader2bytesLimited(limit)
	return rdr2bytes(os.Stdin)
}

var jsonMapBytesStdin IO[[]byte] = Bind(
	mapSizeLimit,
	stdin2bytesLimited,
)

func jsonBytesToMap(jbytes []byte) (map[string]string, error) {
	var ret map[string]string
	e := json.Unmarshal(jbytes, &ret)
	return ret, e
}

var jsonMap IO[map[string]string] = Bind(
	jsonMapBytesStdin,
	Lift(jsonBytesToMap),
)

var rawMap IO[mp.RawMap] = Bind(
	jsonMap,
	Lift(mp.RawMapNew),
)

var derBytes IO[[]byte] = Bind(
	rawMap,
	Lift(func(r mp.RawMap) ([]byte, error) { return r.ToDerBytes() }),
)

var bytes2stdout func([]byte) IO[Void] = Lift(
	func(b []byte) (Void, error) {
		_, e := os.Stdout.Write(b)
		return Empty, e
	},
)

var sub IO[Void] = Bind(
	derBytes,
	bytes2stdout,
)

func main() {
	_, e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
