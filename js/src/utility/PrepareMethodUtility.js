
function prepareMethod(ctx) {
  const op = ctx.op
  const opname = op.name

  let key = opname

  // Luma: the OpenAPI spec is authoritative for the HTTP method. Luma uses POST
  // for every write, so the op-name map below would send PUT and DELETE to
  // endpoints that only accept POST. Prefer the resolved point's method.
  const pointMethod = ctx.point && ctx.point.method
  if (pointMethod) {
    return String(pointMethod).toUpperCase()
  }

  const methodMap = {
    create: 'POST',
    update: 'PUT',
    load: 'GET',
    list: 'GET',
    remove: 'DELETE',
    patch: 'PATCH',
  }

  return methodMap[key]
}

module.exports = {
  prepareMethod
}
