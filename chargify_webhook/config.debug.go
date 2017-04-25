// +build debug

package chargify_webhook

import "fmt"

func debugf(s string, vals ...interface{}) {
	fmt.Printf(s, vals...)
}
