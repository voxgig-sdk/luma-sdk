<?php
declare(strict_types=1);

// Luma SDK utility: result_headers

class LumaResultHeaders
{
    public static function call(LumaContext $ctx): ?LumaResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
