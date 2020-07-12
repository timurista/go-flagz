// Copyright 2015 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package flagz

import (
	"hash/fnv"
	"sync"

	"github.com/spf13/pflag"
)

// @todo Temporary fix for race condition happening in the test:
// spf13/pflag.FlagSet.VisitAll seems to be prone to a race condition
// This fixes that but I'm not sure how much slower does it make the codebase
var visitAllMutex = &sync.Mutex{}

// ChecksumFlagSet will generate a FNV of the *set* values in a FlagSet.
func ChecksumFlagSet(flagSet *pflag.FlagSet, flagFilter func(flag *pflag.Flag) bool) []byte {
	h := fnv.New32a()

	visitAllMutex.Lock()
	defer visitAllMutex.Unlock()
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flagFilter != nil && !flagFilter(flag) {
			return
		}
		h.Write([]byte(flag.Name))
		h.Write([]byte(flag.Value.String()))
	})
	return h.Sum(nil)
}
