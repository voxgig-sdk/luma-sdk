
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


describe('WebhookEntity', async () => {

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.Webhook()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const webhook_ref01_ent = client.Webhook()
    let webhook_ref01_data = setup.data.new.webhook['webhook_ref01']

    webhook_ref01_data = await webhook_ref01_ent.create(webhook_ref01_data)
    assert(null != webhook_ref01_data.id)


    // LIST
    const webhook_ref01_match = {}

    const webhook_ref01_list = await webhook_ref01_ent.list(webhook_ref01_match)

    assert(!isempty(select(webhook_ref01_list, { id: webhook_ref01_data.id })))


    // UPDATE
    const webhook_ref01_data_up0 = {}
    webhook_ref01_data_up0.id = webhook_ref01_data.id

    const webhook_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-webhook_ref01_' + setup.now }
    webhook_ref01_data_up0 [webhook_ref01_markdef_up0.name] = webhook_ref01_markdef_up0.value

    const webhook_ref01_resdata_up0 = await webhook_ref01_ent.update(webhook_ref01_data_up0)
    assert(webhook_ref01_resdata_up0.id === webhook_ref01_data_up0.id)

    assert(webhook_ref01_resdata_up0[webhook_ref01_markdef_up0.name] === webhook_ref01_markdef_up0.value)


    // LOAD
    const webhook_ref01_match_dt0 = {}
    webhook_ref01_match_dt0.id = webhook_ref01_data.id
    const webhook_ref01_data_dt0 = await webhook_ref01_ent.load(webhook_ref01_match_dt0)
    assert(webhook_ref01_data_dt0.id === webhook_ref01_data.id)


    // REMOVE
    const webhook_ref01_match_rm0 = {}
    webhook_ref01_match_rm0.id = webhook_ref01_data.id
    await webhook_ref01_ent.remove(webhook_ref01_match_rm0)
  

    // LIST
    const webhook_ref01_match_rt0 = {}

    const webhook_ref01_list_rt0 = await webhook_ref01_ent.list(webhook_ref01_match_rt0)

    assert(isempty(select(webhook_ref01_list_rt0, { id: webhook_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/webhook/WebhookTestData.json')

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
    ['webhook01','webhook02','webhook03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'LUMA_TEST_WEBHOOK_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_WEBHOOK_ENTID']

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
  
