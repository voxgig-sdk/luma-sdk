-- Luma SDK utility: prepare_method

local METHOD_MAP = {
  create = "POST",
  update = "PUT",
  load = "GET",
  list = "GET",
  remove = "DELETE",
  patch = "PATCH",
}

local function prepare_method_util(ctx)
  -- Luma: the OpenAPI spec is authoritative for the HTTP method. Luma uses POST
  -- for every write, so METHOD_MAP would send PUT/DELETE to POST-only endpoints.
  -- Prefer the method carried on the resolved point (set from the spec).
  local point = ctx.point
  if point ~= nil and point.method ~= nil and point.method ~= "" then
    return string.upper(point.method)
  end

  local opname = ctx.op.name
  return METHOD_MAP[opname] or "GET"
end

return prepare_method_util
