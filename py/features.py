# Luma SDK feature factory

from feature.base_feature import LumaBaseFeature
from feature.test_feature import LumaTestFeature


def _make_feature(name):
    features = {
        "base": lambda: LumaBaseFeature(),
        "test": lambda: LumaTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
