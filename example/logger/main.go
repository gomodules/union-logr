package main

import (
	"flag"

	"github.com/nats-io/stan.go"

	"github.com/go-logr/glogr"
	natslogr "gomodules.xyz/nats-logr"
	natslog "gomodules.xyz/nats-logr/nats-log"
	ulogr "gomodules.xyz/union-logr"
	"k8s.io/klog"
	"k8s.io/klog/klogr"
)

func main() {
	// glog
	logG := glogr.New().WithName("glog")
	flag.Parse()

	flag.CommandLine.Parse([]string{})

	// klogr
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	flag.CommandLine.VisitAll(func(f1 *flag.Flag) {
		if f1.Name == "logtostderr" {
			return
		}
		f2 := klogFlags.Lookup(f1.Name)
		if f2 != nil {
			value := f1.Value.String()
			f2.Value.Set(value)
		}
	})
	logK := klogr.New().WithName("klog")

	flag.CommandLine.Parse([]string{})

	// nats-logr
	natslogFlags := flag.NewFlagSet("nats-log", flag.ExitOnError)
	natslog.InitFlags(natslogFlags)
	flag.CommandLine.VisitAll(func(f1 *flag.Flag) {
		f2 := natslogFlags.Lookup(f1.Name)
		if f2 != nil {
			value := f1.Value.String()
			f2.Value.Set(value)
		}
	})
	logN := natslogr.New().WithName("natslog").WithValues(natslogr.ClusterID, "nats-logr", natslogr.ClientID, "nats-logger", natslogr.NatsURL, stan.DefaultNatsURL, natslogr.ConnectWait, 5, natslogr.Subject, "nats-log-example")

	// union-logr
	ulog := ulogr.NewUnionLogger(logG, logK, logN).WithName("Union Log")
	ulog.V(2).Info("Example", "Name", "Masudur Rahman")
}
