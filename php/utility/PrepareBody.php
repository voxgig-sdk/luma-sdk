<?php
declare(strict_types=1);

// Luma SDK utility: prepare_body

class LumaPrepareBody
{
    public static function call(LumaContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
