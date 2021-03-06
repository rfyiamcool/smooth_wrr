# smooth_wrr

smooth weighted round robin algo like nginx

https://github.com/nginx/nginx/commit/52327e0627f49dbda1e8db695e63a4b0af4448b1

## Usage

```
package main

import (
	"fmt"
	"strings"

	"github.com/rfyiamcool/smooth_wrr"
)

func main() {
	peers := []*smoothwrr.Peer{
		smoothwrr.CreatePeer("a", 5),
		smoothwrr.CreatePeer("b", 1),
		smoothwrr.CreatePeer("c", 1),
	}

	pool := smoothwrr.NewPool(peers)
	for index := 0; index < 5; index++ {
		fmt.Println("Loop: ", index)

		result := []string{}
		for i := 0; i < 7; i++ {
			peer := pool.Get()
			result = append(result, peer)
		}

		resultStr := strings.Join(result, ",")
		fmt.Println(resultStr)
	}
}
```

**stdout:**

```
Loop:  0
a,a,b,a,c,a,a
Loop:  1
a,a,b,a,c,a,a
Loop:  2
a,a,b,a,c,a,a
Loop:  3
a,a,b,a,c,a,a
Loop:  4
a,a,b,a,c,a,a
```
