package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// very naive implementation of a counter.
func incrementHitCounter(ctx context.Context, cli *clientv3.Client) int64 {
	kv := clientv3.NewKV(cli)
	for {
		r, err := kv.Get(ctx, "hit-counter")
		if err != nil {
			log.Printf("WARN failed to get hit-counter: %v", err)
			continue
		}
		previousHitCounterValue := "0"
		if len(r.Kvs) != 0 {
			previousHitCounterValue = string(r.Kvs[0].Value)
		}
		hitCounter, err := strconv.ParseInt(previousHitCounterValue, 10, 64)
		if err != nil {
			log.Printf("WARN failed to get hit-counter: %v", err)
			continue
		}
		hitCounter = hitCounter + 1
		_, err = kv.Put(ctx, "hit-counter", strconv.FormatInt(hitCounter, 10))
		if err != nil {
			log.Printf("WARN failed to set hit-counter: %v", err)
			continue
		}
		return hitCounter
	}
}

func main() {
	log.SetFlags(0)

	var listenAddress = flag.String("listen", ":8888", "Listen address.")

	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		log.Fatalf("\nYou MUST NOT pass any positional arguments")
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"etcd:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer cli.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s%s\n", r.Method, r.Host, r.URL)

		if r.URL.Path != "/" {
			w.WriteHeader(404)
			return
		}

		hitCounter := incrementHitCounter(context.TODO(), cli)

		fmt.Fprintf(
			w,
			`Hello World!

etcd hit-counter: %d
HTTP Request: %s %s%s
Server IP address: %s
Client IP address: %s
`,
			hitCounter,
			r.Method,
			r.Host,
			r.URL,
			r.Context().Value(http.LocalAddrContextKey).(net.Addr).String(),
			r.RemoteAddr)
		for _, e := range os.Environ() {
			fmt.Fprintf(w, "Environment Variable: %s\n", e)
		}
	})

	fmt.Printf("Listening at http://%s\n", *listenAddress)

	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
