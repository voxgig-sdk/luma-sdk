
import { Context } from '../types'

function prepareMethod(ctx: Context) {
  const op = ctx.op
  const opname = op.name

  let key = opname

  // Luma: the OpenAPI spec is authoritative for the HTTP method. Luma is an
  // RPC-style API that uses POST for every write (/v1/events/update,
  // /v1/webhooks/delete), so the op-name map below would send PUT and DELETE to
  // endpoints that only accept POST. Prefer the method carried on the resolved
  // point (set from the spec); fall back to the map only if the point has none.
  const pointMethod = ctx.point && ctx.point.method
  if (pointMethod) {
    return String(pointMethod).toUpperCase()
  }

  const methodMap: any = {
    create: 'POST',
    update: 'PUT',
    load: 'GET',
    list: 'GET',
    remove: 'DELETE',
    patch: 'PATCH',
  }

  return methodMap[key]
}


export {
  prepareMethod
}
