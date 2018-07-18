# smooth_wrr

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
