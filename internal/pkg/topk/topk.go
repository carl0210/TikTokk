package topk

import "github.com/go-kratos/aegis/topk"

var MyTopK topk.Topk

func init() {
	MyTopK = topk.NewHeavyKeeper(128, 10000, 5, 0.925, 0)
}
