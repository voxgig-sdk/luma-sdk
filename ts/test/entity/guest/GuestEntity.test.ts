
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


describe('GuestEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when LUMA_TEST_LIVE=TRUE.
  afterEach(liveDelay('LUMA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.Guest()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.LUMA_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load']) {
      if (maybeSkipControl(t, 'entityOp', 'guest.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set LUMA_TEST_GUEST_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const guest_ref01_ent = client.Guest()
    let guest_ref01_data = setup.data.new.guest['guest_ref01']

    guest_ref01_data = await guest_ref01_ent.create(guest_ref01_data)
    assert(null != guest_ref01_data.id)


    // LIST
    const guest_ref01_match: any = {}

    const guest_ref01_list = await guest_ref01_ent.list(guest_ref01_match)

    assert(!isempty(select(guest_ref01_list, { id: guest_ref01_data.id })))


    // UPDATE
    const guest_ref01_data_up0: any = {}
    guest_ref01_data_up0.id = guest_ref01_data.id

    const guest_ref01_markdef_up0 = { name: 'approval_status', value: 'Mark01-guest_ref01_' + setup.now }
    ;(guest_ref01_data_up0 as any)[guest_ref01_markdef_up0.name] = guest_ref01_markdef_up0.value

    const guest_ref01_resdata_up0 = await guest_ref01_ent.update(guest_ref01_data_up0)
    assert(guest_ref01_resdata_up0.id === guest_ref01_data_up0.id)

    assert((guest_ref01_resdata_up0 as any)[guest_ref01_markdef_up0.name] === guest_ref01_markdef_up0.value)


    // LOAD
    const guest_ref01_match_dt0: any = {}
    guest_ref01_match_dt0.id = guest_ref01_data.id
    const guest_ref01_data_dt0 = await guest_ref01_ent.load(guest_ref01_match_dt0)
    assert(guest_ref01_data_dt0.id === guest_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/guest/GuestTestData.json')

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
    ['guest01','guest02','guest03'],
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
  const idmapEnvVal = process.env['LUMA_TEST_GUEST_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'LUMA_TEST_GUEST_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_GUEST_ENTID']

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
  
