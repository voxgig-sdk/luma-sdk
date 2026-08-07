# Luma SDK utility: make_context

from core.context import LumaContext


def make_context_util(ctxmap, basectx):
    return LumaContext(ctxmap, basectx)
