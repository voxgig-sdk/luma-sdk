
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


describe('GuestEntity', async () => {

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.Guest()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
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
    const guest_ref01_match = {}

    const guest_ref01_list = await guest_ref01_ent.list(guest_ref01_match)

    assert(!isempty(select(guest_ref01_list, { id: guest_ref01_data.id })))


    // UPDATE
    const guest_ref01_data_up0 = {}
    guest_ref01_data_up0.id = guest_ref01_data.id

    const guest_ref01_markdef_up0 = { name: 'approval_status', value: 'Mark01-guest_ref01_' + setup.now }
    guest_ref01_data_up0 [guest_ref01_markdef_up0.name] = guest_ref01_markdef_up0.value

    const guest_ref01_resdata_up0 = await guest_ref01_ent.update(guest_ref01_data_up0)
    assert(guest_ref01_resdata_up0.id === guest_ref01_data_up0.id)

    assert(guest_ref01_resdata_up0[guest_ref01_markdef_up0.name] === guest_ref01_markdef_up0.value)


    // LOAD
    const guest_ref01_match_dt0 = {}
    guest_ref01_match_dt0.id = guest_ref01_data.id
    const guest_ref01_data_dt0 = await guest_ref01_ent.load(guest_ref01_match_dt0)
    assert(guest_ref01_data_dt0.id === guest_ref01_data.id)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

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

  const env = envOverride({
    'LUMA_TEST_GUEST_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_GUEST_ENTID']

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
  
