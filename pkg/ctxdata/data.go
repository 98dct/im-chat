package ctxdata

import "context"

func GetUid(c context.Context) string {
	if v, ok := c.Value(Identity).(string); ok {
		return v
	}
	return ""
}
