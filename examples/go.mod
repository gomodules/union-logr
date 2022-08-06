module gomodules.xyz/union-logr/examples

go 1.18

require (
	github.com/go-logr/glogr v1.2.2
	gomodules.xyz/union-logr v0.0.0-00010101000000-000000000000
	k8s.io/klog v1.0.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
)

replace gomodules.xyz/union-logr => ./..
