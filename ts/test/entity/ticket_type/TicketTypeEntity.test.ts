
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


describe('TicketTypeEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when LUMA_TEST_LIVE=TRUE.
  afterEach(liveDelay('LUMA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.TicketType()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.LUMA_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (maybeSkipControl(t, 'entityOp', 'ticket_type.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set LUMA_TEST_TICKET_TYPE_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const ticket_type_ref01_ent = client.TicketType()
    let ticket_type_ref01_data = setup.data.new.ticket_type['ticket_type_ref01']

    ticket_type_ref01_data = await ticket_type_ref01_ent.create(ticket_type_ref01_data)
    assert(null != ticket_type_ref01_data.id)


    // LIST
    const ticket_type_ref01_match: any = {}

    const ticket_type_ref01_list = await ticket_type_ref01_ent.list(ticket_type_ref01_match)

    assert(!isempty(select(ticket_type_ref01_list, { id: ticket_type_ref01_data.id })))


    // UPDATE
    const ticket_type_ref01_data_up0: any = {}
    ticket_type_ref01_data_up0.id = ticket_type_ref01_data.id

    const ticket_type_ref01_markdef_up0 = { name: 'description', value: 'Mark01-ticket_type_ref01_' + setup.now }
    ;(ticket_type_ref01_data_up0 as any)[ticket_type_ref01_markdef_up0.name] = ticket_type_ref01_markdef_up0.value

    const ticket_type_ref01_resdata_up0 = await ticket_type_ref01_ent.update(ticket_type_ref01_data_up0)
    assert(ticket_type_ref01_resdata_up0.id === ticket_type_ref01_data_up0.id)

    assert((ticket_type_ref01_resdata_up0 as any)[ticket_type_ref01_markdef_up0.name] === ticket_type_ref01_markdef_up0.value)


    // LOAD
    const ticket_type_ref01_match_dt0: any = {}
    ticket_type_ref01_match_dt0.id = ticket_type_ref01_data.id
    const ticket_type_ref01_data_dt0 = await ticket_type_ref01_ent.load(ticket_type_ref01_match_dt0)
    assert(ticket_type_ref01_data_dt0.id === ticket_type_ref01_data.id)


    // REMOVE
    const ticket_type_ref01_match_rm0: any = { id: ticket_type_ref01_data.id }
    await ticket_type_ref01_ent.remove(ticket_type_ref01_match_rm0)
  

    // LIST
    const ticket_type_ref01_match_rt0: any = {}

    const ticket_type_ref01_list_rt0 = await ticket_type_ref01_ent.list(ticket_type_ref01_match_rt0)

    assert(isempty(select(ticket_type_ref01_list_rt0, { id: ticket_type_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/ticket_type/TicketTypeTestData.json')

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
    ['ticket_type01','ticket_type02','ticket_type03'],
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
  const idmapEnvVal = process.env['LUMA_TEST_TICKET_TYPE_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'LUMA_TEST_TICKET_TYPE_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_TICKET_TYPE_ENTID']

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
  
