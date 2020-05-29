package filter

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	err := Init("../data/filter.txt")
	if err != nil {
		t.Errorf("load filter data failed, err:%v\n", err)
		return
	}

	data := `
	作者的观点在 issue 中得到了非常多的 👎，但是这件事情很难说对错；在社区中保证一致的编程规范是一件非常有益的事情,公安
	不过对于很多公司内部的服务或者项目，可能在业务服务上就会发生一些比较棘手的情况，使用这种过强的约束没有太多明显地收益。
	`
	result, hit := Replace(data, "***")
	fmt.Printf("isReplace:%#v, str:%v\n", hit, result)
}
