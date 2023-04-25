package counter

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

type Counter uint32

func (c *Counter) Count() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		countValue := atomic.AddUint32((*uint32)(c), 1)
		stringCountValue := strconv.Itoa(int(countValue))
		_, _ = w.Write([]byte(stringCountValue))
	}
}
