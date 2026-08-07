<?php
declare(strict_types=1);

// Luma SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

LumaUtility::setRegistrar(function (LumaUtility $u): void {
    $u->clean = [LumaClean::class, 'call'];
    $u->done = [LumaDone::class, 'call'];
    $u->make_error = [LumaMakeError::class, 'call'];
    $u->feature_add = [LumaFeatureAdd::class, 'call'];
    $u->feature_hook = [LumaFeatureHook::class, 'call'];
    $u->feature_init = [LumaFeatureInit::class, 'call'];
    $u->fetcher = [LumaFetcher::class, 'call'];
    $u->make_fetch_def = [LumaMakeFetchDef::class, 'call'];
    $u->make_context = [LumaMakeContext::class, 'call'];
    $u->make_options = [LumaMakeOptions::class, 'call'];
    $u->make_request = [LumaMakeRequest::class, 'call'];
    $u->make_response = [LumaMakeResponse::class, 'call'];
    $u->make_result = [LumaMakeResult::class, 'call'];
    $u->make_point = [LumaMakePoint::class, 'call'];
    $u->make_spec = [LumaMakeSpec::class, 'call'];
    $u->make_url = [LumaMakeUrl::class, 'call'];
    $u->param = [LumaParam::class, 'call'];
    $u->prepare_auth = [LumaPrepareAuth::class, 'call'];
    $u->prepare_body = [LumaPrepareBody::class, 'call'];
    $u->prepare_headers = [LumaPrepareHeaders::class, 'call'];
    $u->prepare_method = [LumaPrepareMethod::class, 'call'];
    $u->prepare_params = [LumaPrepareParams::class, 'call'];
    $u->prepare_path = [LumaPreparePath::class, 'call'];
    $u->prepare_query = [LumaPrepareQuery::class, 'call'];
    $u->result_basic = [LumaResultBasic::class, 'call'];
    $u->result_body = [LumaResultBody::class, 'call'];
    $u->result_headers = [LumaResultHeaders::class, 'call'];
    $u->transform_request = [LumaTransformRequest::class, 'call'];
    $u->transform_response = [LumaTransformResponse::class, 'call'];
});
