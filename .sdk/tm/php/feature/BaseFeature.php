<?php
declare(strict_types=1);

// Luma SDK base feature

class LumaBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(LumaContext $ctx, array $options): void {}
    public function PostConstruct(LumaContext $ctx): void {}
    public function PostConstructEntity(LumaContext $ctx): void {}
    public function SetData(LumaContext $ctx): void {}
    public function GetData(LumaContext $ctx): void {}
    public function GetMatch(LumaContext $ctx): void {}
    public function SetMatch(LumaContext $ctx): void {}
    public function PrePoint(LumaContext $ctx): void {}
    public function PreSpec(LumaContext $ctx): void {}
    public function PreRequest(LumaContext $ctx): void {}
    public function PreResponse(LumaContext $ctx): void {}
    public function PreResult(LumaContext $ctx): void {}
    public function PreDone(LumaContext $ctx): void {}
    public function PreUnexpected(LumaContext $ctx): void {}
}
