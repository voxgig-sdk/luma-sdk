# Luma SDK utility: prepare_method

METHOD_MAP = {
    "create": "POST",
    "update": "PUT",
    "load": "GET",
    "list": "GET",
    "remove": "DELETE",
    "patch": "PATCH",
}


def prepare_method_util(ctx):
    # Luma: the OpenAPI spec is authoritative for the HTTP method. Luma uses POST
    # for every write, so METHOD_MAP would send PUT/DELETE to POST-only endpoints.
    # Prefer the method carried on the resolved point (set from the spec).
    point = getattr(ctx, "point", None)
    point_method = None
    if isinstance(point, dict):
        point_method = point.get("method")
    elif point is not None:
        point_method = getattr(point, "method", None)
    if point_method:
        return str(point_method).upper()

    opname = ctx.op.name
    return METHOD_MAP.get(opname, "GET")
