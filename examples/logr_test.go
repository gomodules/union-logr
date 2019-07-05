package ulogr_test

import (
	"flag"
	"fmt"

	"github.com/go-logr/glogr"
	"k8s.io/klog"
	"k8s.io/klog/klogr"

	ulogr "gomodules.xyz/union-logr"
)

func Example() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	logG := glogr.New().WithName("glog")

	klog.InitFlags(flag.NewFlagSet("klog", flag.ExitOnError))
	logK := klogr.New().WithName("klog")

	ulog := ulogr.NewUnionLogger(logG, logK).WithName("ulog").WithValues("logr", "union-logr")
	ulog.V(0).Info("Example", "Key", "Value")

	fmt.Println("Example")
	//	Output: Example
}
