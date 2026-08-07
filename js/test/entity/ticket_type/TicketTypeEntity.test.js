
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


describe('TicketTypeEntity', async () => {

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.TicketType()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
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
    const ticket_type_ref01_match = {}

    const ticket_type_ref01_list = await ticket_type_ref01_ent.list(ticket_type_ref01_match)

    assert(!isempty(select(ticket_type_ref01_list, { id: ticket_type_ref01_data.id })))


    // UPDATE
    const ticket_type_ref01_data_up0 = {}
    ticket_type_ref01_data_up0.id = ticket_type_ref01_data.id

    const ticket_type_ref01_markdef_up0 = { name: 'description', value: 'Mark01-ticket_type_ref01_' + setup.now }
    ticket_type_ref01_data_up0 [ticket_type_ref01_markdef_up0.name] = ticket_type_ref01_markdef_up0.value

    const ticket_type_ref01_resdata_up0 = await ticket_type_ref01_ent.update(ticket_type_ref01_data_up0)
    assert(ticket_type_ref01_resdata_up0.id === ticket_type_ref01_data_up0.id)

    assert(ticket_type_ref01_resdata_up0[ticket_type_ref01_markdef_up0.name] === ticket_type_ref01_markdef_up0.value)


    // LOAD
    const ticket_type_ref01_match_dt0 = {}
    ticket_type_ref01_match_dt0.id = ticket_type_ref01_data.id
    const ticket_type_ref01_data_dt0 = await ticket_type_ref01_ent.load(ticket_type_ref01_match_dt0)
    assert(ticket_type_ref01_data_dt0.id === ticket_type_ref01_data.id)


    // REMOVE
    const ticket_type_ref01_match_rm0 = {}
    ticket_type_ref01_match_rm0.id = ticket_type_ref01_data.id
    await ticket_type_ref01_ent.remove(ticket_type_ref01_match_rm0)
  

    // LIST
    const ticket_type_ref01_match_rt0 = {}

    const ticket_type_ref01_list_rt0 = await ticket_type_ref01_ent.list(ticket_type_ref01_match_rt0)

    assert(isempty(select(ticket_type_ref01_list_rt0, { id: ticket_type_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

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

  const env = envOverride({
    'LUMA_TEST_TICKET_TYPE_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_TICKET_TYPE_ENTID']

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
  
