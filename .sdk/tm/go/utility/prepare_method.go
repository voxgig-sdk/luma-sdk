package utility

import (
	"strings"

	"GOMODULE/core"
)

func prepareMethodUtil(ctx *core.Context) string {
	// Luma: the OpenAPI spec is authoritative for the HTTP method. Luma uses POST
	// for every write, so methodMap would send PUT/DELETE to POST-only endpoints.
	// Prefer the method carried on the resolved point (set from the spec).
	if ctx.Point != nil {
		if pm, ok := ctx.Point["method"].(string); ok && pm != "" {
			return strings.ToUpper(pm)
		}
	}

	opname := ctx.Op.Name

	methodMap := map[string]string{
		"create": "POST",
		"update": "PUT",
		"load":   "GET",
		"list":   "GET",
		"remove": "DELETE",
		"patch":  "PATCH",
	}

	if m, ok := methodMap[opname]; ok {
		return m
	}
	return "GET"
}
