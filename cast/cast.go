package cast

import "github.com/spf13/cast"

func ToBool(v any) bool {
	return cast.ToBool(v)
}

func ToString(v any) string {
	return cast.ToString(v)
}
