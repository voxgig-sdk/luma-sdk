<?php
declare(strict_types=1);

// Luma SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class LumaFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new LumaBaseFeature();
            case "test":
                return new LumaTestFeature();
            default:
                return new LumaBaseFeature();
        }
    }
}
