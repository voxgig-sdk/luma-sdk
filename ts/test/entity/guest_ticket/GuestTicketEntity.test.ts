
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


describe('GuestTicketEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when LUMA_TEST_LIVE=TRUE.
  afterEach(liveDelay('LUMA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.GuestTicket()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.LUMA_TEST_LIVE
    for (const op of ['update']) {
      if (maybeSkipControl(t, 'entityOp', 'guest_ticket.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set LUMA_TEST_GUEST_TICKET_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let guest_ticket_ref01_data = Object.values(setup.data.existing.guest_ticket)[0] as any

    // UPDATE
    const guest_ticket_ref01_ent = client.GuestTicket()
    const guest_ticket_ref01_data_up0: any = {}

    const guest_ticket_ref01_markdef_up0 = { name: 'event_id', value: 'Mark01-guest_ticket_ref01_' + setup.now }
    ;(guest_ticket_ref01_data_up0 as any)[guest_ticket_ref01_markdef_up0.name] = guest_ticket_ref01_markdef_up0.value

    const guest_ticket_ref01_resdata_up0 = await guest_ticket_ref01_ent.update(guest_ticket_ref01_data_up0)
    assert(null != guest_ticket_ref01_resdata_up0)

    assert((guest_ticket_ref01_resdata_up0 as any)[guest_ticket_ref01_markdef_up0.name] === guest_ticket_ref01_markdef_up0.value)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/guest_ticket/GuestTicketTestData.json')

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
    ['guest_ticket01','guest_ticket02','guest_ticket03'],
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
  const idmapEnvVal = process.env['LUMA_TEST_GUEST_TICKET_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'LUMA_TEST_GUEST_TICKET_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_GUEST_TICKET_ENTID']

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
  
