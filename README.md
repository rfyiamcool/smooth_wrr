# smooth_wrr

smooth weighted round robin algo like nginx

https://github.com/nginx/nginx/commit/52327e0627f49dbda1e8db695e63a4b0af4448b1

```

func main(){
	peers := []*Peer{
		CreatePeer("a", 5),
		CreatePeer("b", 1),
		CreatePeer("c", 1),
	}

	result := []string{}
	pool := &Pool{peers: peers}
	for i := 0; i < 7; i++ {
		peer := pool.Get()
		result = append(result, peer)
	}
	resultStr := strings.Join(result, ",")
	fmt.Println(resultStr)

}
```
