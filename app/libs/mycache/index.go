package mycache

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

const compress = true

var mc *memcache.Client
var memcachedHostStr string
var prefix = "cache"

func init() {
	memcacheHost := getenvWithDefault("MEMCACHED_HOST", "localhost")
	memcachePort := getenvWithDefault("MEMCACHED_POST", "11211")

	memcachedHostStr = memcacheHost + ":" + memcachePort
	mc = memcache.New(memcachedHostStr)
}

func GetHostStr() string {
	return memcachedHostStr
}

func Get(suffix string) (string, error) {
	var key string

	if compress {
		key = prefix + ".c." + suffix
	} else {
		key = prefix + "." + suffix
	}

	it, err := mc.Get(key)

	if err != nil {
		return "", err
	}

	if compress {
		return gzuncompress(it.Value)
	}

	return string(it.Value), nil
}

func Set(suffix string, val string, ttl int64) (bool, error) {
	var key string
	var err error

	if compress {
		key = prefix + ".c." + suffix
		err = mc.Set(&memcache.Item{
			Key:        key,
			Value:      gzcompress(val),
			Expiration: int32(ttl),
		})
	} else {
		key = prefix + "." + suffix
		err = mc.Set(&memcache.Item{
			Key:        key,
			Value:      []byte(val),
			Expiration: int32(ttl),
		})
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func Delete(suffix string) (bool, error) {
	key := prefix + "." + suffix

	if compress {
		key = prefix + ".c." + suffix
	}

	err := mc.Delete(key)

	if err != nil {
		return false, err
	}

	return true, nil
}

func gzcompress(str string) []byte {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)

	if _, err := gz.Write([]byte(str)); err != nil {
		return []byte("")
	}

	if err := gz.Flush(); err != nil {
		return []byte("")
	}

	if err := gz.Close(); err != nil {
		return []byte("")
	}

	return b.Bytes()
}

func gzuncompress(b []byte) (string, error) {
	bb := bytes.NewBuffer(b)
	zipread, _ := gzip.NewReader(bb)

	defer zipread.Close()

	reader := bufio.NewReader(zipread)
	ret := ""

	var part []byte
	var err error

	for {
		if part, _, err = reader.ReadLine(); err != nil {
			break
		}

		ret += string(part)
	}

	return ret, nil
}
