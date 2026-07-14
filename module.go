package xk6modules

import "go.k6.io/k6/v2/js/modules"

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/zitadel", new(Mod))
}

type Mod struct{}
