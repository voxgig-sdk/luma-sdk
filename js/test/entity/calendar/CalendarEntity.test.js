
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

const Path = require('node:path')
const Fs = require('node:fs')

const { test, describe } = require('node:test')
const assert = require('node:assert')


const { LumaSDK, BaseFeature, stdutil, config } = require('../../..')

const {
  envOverride,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
} = require('../../utility')


describe('CalendarEntity', async () => {

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.Calendar()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let calendar_ref01_data = Object.values(setup.data.existing.calendar)[0]

    // UPDATE
    const calendar_ref01_ent = client.Calendar()
    const calendar_ref01_data_up0 = {}
    calendar_ref01_data_up0.id = calendar_ref01_data.id

    const calendar_ref01_markdef_up0 = { name: 'calendar_id', value: 'Mark01-calendar_ref01_' + setup.now }
    calendar_ref01_data_up0 [calendar_ref01_markdef_up0.name] = calendar_ref01_markdef_up0.value

    const calendar_ref01_resdata_up0 = await calendar_ref01_ent.update(calendar_ref01_data_up0)
    assert(calendar_ref01_resdata_up0.id === calendar_ref01_data_up0.id)

    assert(calendar_ref01_resdata_up0[calendar_ref01_markdef_up0.name] === calendar_ref01_markdef_up0.value)


    // LOAD
    const calendar_ref01_match_dt0 = {}
    calendar_ref01_match_dt0.id = calendar_ref01_data.id
    const calendar_ref01_data_dt0 = await calendar_ref01_ent.load(calendar_ref01_match_dt0)
    assert(calendar_ref01_data_dt0.id === calendar_ref01_data.id)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/calendar/CalendarTestData.json')

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
    ['calendar01','calendar02','calendar03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'LUMA_TEST_CALENDAR_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_CALENDAR_ENTID']

  if ('TRUE' === env.LUMA_TEST_LIVE) {
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
    now: Date.now(),
  }

  return setup
}
  
