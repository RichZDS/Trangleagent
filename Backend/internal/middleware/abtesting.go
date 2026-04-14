package middleware

import (
	"hash/crc32"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	// CtxABTestGroup is the context key for storing the assigned A/B test group
	CtxABTestGroup gctx.StrKey = "ab_test_group"
	
	// GroupA represents the original/control group
	GroupA = "A"
	
	// GroupB represents the new/treatment group
	GroupB = "B"
)

// ABTesting is a middleware that assigns users to A/B test groups based on their token/username
func ABTesting(r *ghttp.Request) {
	username := r.GetCtxVar(CtxUsername).String()
	
	group := GroupA // Default group for unauthenticated or errors
	
	if username != "" {
		// Use a stable hash function to assign users consistently
		hash := crc32.ChecksumIEEE([]byte(username))
		if hash%2 == 0 {
			group = GroupA
		} else {
			group = GroupB
		}
	}
	
	// Set the group in context for controllers/services to use
	r.SetCtxVar(CtxABTestGroup, group)
	r.Middleware.Next()
}
