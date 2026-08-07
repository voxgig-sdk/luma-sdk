<?php
declare(strict_types=1);

// Luma SDK utility: prepare_method

class LumaPrepareMethod
{
    private const METHOD_MAP = [
        'create' => 'POST',
        'update' => 'PUT',
        'load' => 'GET',
        'list' => 'GET',
        'remove' => 'DELETE',
        'patch' => 'PATCH',
    ];

    public static function call(LumaContext $ctx): string
    {
        // Luma: the OpenAPI spec is authoritative for the HTTP method. Luma uses
        // POST for every write, so METHOD_MAP would send PUT/DELETE to POST-only
        // endpoints. Prefer the method carried on the resolved point.
        $point = $ctx->point ?? null;
        $pointMethod = null;
        if (is_array($point)) {
            $pointMethod = $point['method'] ?? null;
        } elseif (is_object($point)) {
            $pointMethod = $point->method ?? null;
        }
        if (!empty($pointMethod)) {
            return strtoupper((string)$pointMethod);
        }

        return self::METHOD_MAP[$ctx->op->name] ?? 'GET';
    }
}
