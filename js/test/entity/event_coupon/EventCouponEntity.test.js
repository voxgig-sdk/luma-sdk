
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


describe('EventCouponEntity', async () => {

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.EventCoupon()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
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
    const event_coupon_ref01_match = {}

    const event_coupon_ref01_list = await event_coupon_ref01_ent.list(event_coupon_ref01_match)

    assert(!isempty(select(event_coupon_ref01_list, { id: event_coupon_ref01_data.id })))


    // UPDATE
    const event_coupon_ref01_data_up0 = {}
    event_coupon_ref01_data_up0.id = event_coupon_ref01_data.id

    const event_coupon_ref01_markdef_up0 = { name: 'code', value: 'Mark01-event_coupon_ref01_' + setup.now }
    event_coupon_ref01_data_up0 [event_coupon_ref01_markdef_up0.name] = event_coupon_ref01_markdef_up0.value

    const event_coupon_ref01_resdata_up0 = await event_coupon_ref01_ent.update(event_coupon_ref01_data_up0)
    assert(event_coupon_ref01_resdata_up0.id === event_coupon_ref01_data_up0.id)

    assert(event_coupon_ref01_resdata_up0[event_coupon_ref01_markdef_up0.name] === event_coupon_ref01_markdef_up0.value)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

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

  const env = envOverride({
    'LUMA_TEST_EVENT_COUPON_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_EVENT_COUPON_ENTID']

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
  
