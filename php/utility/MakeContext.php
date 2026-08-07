<?php
declare(strict_types=1);

// Luma SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class LumaMakeContext
{
    public static function call(array $ctxmap, ?LumaContext $basectx): LumaContext
    {
        return new LumaContext($ctxmap, $basectx);
    }
}
