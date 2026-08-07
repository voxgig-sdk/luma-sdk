# Luma SDK exists test

import pytest
from luma_sdk import LumaSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = LumaSDK.test(None, None)
        assert testsdk is not None
