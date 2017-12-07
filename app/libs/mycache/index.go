package mycache

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

const (
	compress = true
)

var prefix = "mycache."

var memcachedHostStr string

func init() {
	mhost := getenvWithDefault("MEMCACHED_HOST", "localhost")
	mport := getenvWithDefault("MEMCACHED_POST", "11211")
	memcachedHostStr = mhost + ":" + mport
	mc = memcache.New(memcachedHostStr)

	dbname := getenvWithDefault("V3_DB_NAME", "websta")
	if dbname == "websta_dev" {
		prefix += "d."
	}

}

// GetHostStr ...
func GetHostStr() string {
	return memcachedHostStr
}

// Get ...
func Get(suffix string) (string, error) {
	var key string
	if compress {
		key = prefix + ".c." + suffix
	} else {
		key = prefix + suffix
	}

	//log.Printf("mycache.Get: key=%s", key)
	it, err := mc.Get(key)
	if err != nil {
		//log.Printf("%s , key=%s", err.Error(), key)
		return "", err
	}
	if compress {
		return gzuncompress(it.Value)
	}
	//log.Print(it.Value)
	return string(it.Value), nil
}

// Set ...
func Set(suffix string, val string, ttl int64) (bool, error) {
	var key string
	var e error
	if compress {
		key = prefix + ".c." + suffix
		//log.Printf("Set:%s", key)
		e = mc.Set(&memcache.Item{
			Key:        key,
			Value:      gzcompress(val),
			Expiration: int32(ttl),
		})
	} else {
		key = prefix + suffix
		//log.Printf("Set:%s", key)
		e = mc.Set(&memcache.Item{
			Key:        key,
			Value:      []byte(val),
			Expiration: int32(ttl),
		})
	}

	if e != nil {
		return false, e
	}
	return true, nil
}

// Delete ...
func Delete(suffix string) (bool, error) {
	key := prefix + suffix
	if compress {
		key = prefix + ".c." + suffix
	}

	e := mc.Delete(key)

	if e != nil {
		return false, e
	}

	return true, nil
}

// GetenvWithDefault ...
func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func gzcompress(str string) []byte {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)

	//bufWrite := bufio.NewWriter(bufzip)
	if _, err := gz.Write([]byte(str)); err != nil {
		//panic(err)
		return []byte("")
	}
	if err := gz.Flush(); err != nil {
		//panic(err)
		return []byte("")
	}
	if err := gz.Close(); err != nil {
		//panic(err)
		return []byte("")
	}
	return b.Bytes()

}

func gzuncompress(b []byte) (string, error) {
	bb := bytes.NewBuffer(b)
	zipread, _ := gzip.NewReader(bb)

	defer zipread.Close()
	reader := bufio.NewReader(zipread)

	var (
		part []byte
		err  error
	)
	ret := ""

	for {
		if part, _, err = reader.ReadLine(); err != nil {
			break
		}

		ret += string(part)

	}
	return ret, nil

}
