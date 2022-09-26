package day001

import (
	"testing"

	"github.com/spf13/cast"
)

func TestCast(t *testing.T) {
	t.Log("cast.ToString")
	t.Log(cast.ToString("https://oscome.cn"))
	t.Log(cast.ToString(8.88))
	t.Log(cast.ToString([]byte("https://oscome.cn")))
	var abc interface{} = "https://oscome.cn"
	t.Log(cast.ToString(abc))

	t.Log("cast.ToInt")
	t.Log(cast.ToInt("8"))
	t.Log(cast.ToInt64E("8.99"))

	t.Log("cast.ToInt")
	t.Log(cast.ToBool("1"))
	t.Log(cast.ToBool("8.99"))
}

func TestCastMore(t *testing.T) {
	t.Log("cast.ToTime")
	t.Log(cast.ToTime("2022-01-02 01:01:01 +0000 UTC"))

	t.Log("cast.ToDuration")
	t.Log(cast.ToDuration(1e9))

	t.Log("cast.ToStringSlice")
	t.Log(cast.ToStringSlice([]int{1, 3}))
}
