<?php
declare(strict_types=1);

// Luma SDK exists test

require_once __DIR__ . '/../luma_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = LumaSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
