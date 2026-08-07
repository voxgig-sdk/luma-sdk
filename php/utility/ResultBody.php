<?php
declare(strict_types=1);

// Luma SDK utility: result_body

class LumaResultBody
{
    public static function call(LumaContext $ctx): ?LumaResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
