
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { LumaSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('EventCouponEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when LUMA_TEST_LIVE=TRUE.
  afterEach(liveDelay('LUMA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.EventCoupon()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.LUMA_TEST_LIVE
    for (const op of ['create', 'list', 'update']) {
      if (maybeSkipControl(t, 'entityOp', 'event_coupon.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set LUMA_TEST_EVENT_COUPON_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const event_coupon_ref01_ent = client.EventCoupon()
    let event_coupon_ref01_data = setup.data.new.event_coupon['event_coupon_ref01']

    event_coupon_ref01_data = await event_coupon_ref01_ent.create(event_coupon_ref01_data)
    assert(null != event_coupon_ref01_data.id)


    // LIST
    const event_coupon_ref01_match: any = {}

    const event_coupon_ref01_list = await event_coupon_ref01_ent.list(event_coupon_ref01_match)

    assert(!isempty(select(event_coupon_ref01_list, { id: event_coupon_ref01_data.id })))


    // UPDATE
    const event_coupon_ref01_data_up0: any = {}
    event_coupon_ref01_data_up0.id = event_coupon_ref01_data.id

    const event_coupon_ref01_markdef_up0 = { name: 'code', value: 'Mark01-event_coupon_ref01_' + setup.now }
    ;(event_coupon_ref01_data_up0 as any)[event_coupon_ref01_markdef_up0.name] = event_coupon_ref01_markdef_up0.value

    const event_coupon_ref01_resdata_up0 = await event_coupon_ref01_ent.update(event_coupon_ref01_data_up0)
    assert(event_coupon_ref01_resdata_up0.id === event_coupon_ref01_data_up0.id)

    assert((event_coupon_ref01_resdata_up0 as any)[event_coupon_ref01_markdef_up0.name] === event_coupon_ref01_markdef_up0.value)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/event_coupon/EventCouponTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = LumaSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['event_coupon01','event_coupon02','event_coupon03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['LUMA_TEST_EVENT_COUPON_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'LUMA_TEST_EVENT_COUPON_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_EVENT_COUPON_ENTID']

  const live = 'TRUE' === env.LUMA_TEST_LIVE

  if (live) {
    client = new LumaSDK(merge([
      {
        apikey: env.LUMA_APIKEY,
      },
      extra
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.LUMA_TEST_EXPLAIN,
    live,
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
