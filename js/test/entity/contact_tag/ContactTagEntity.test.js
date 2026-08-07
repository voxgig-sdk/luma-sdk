
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


describe('ContactTagEntity', async () => {

  test('instance', async () => {
    const testsdk = LumaSDK.test()
    const ent = testsdk.ContactTag()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const contact_tag_ref01_ent = client.ContactTag()
    let contact_tag_ref01_data = setup.data.new.contact_tag['contact_tag_ref01']

    contact_tag_ref01_data = await contact_tag_ref01_ent.create(contact_tag_ref01_data)
    assert(null != contact_tag_ref01_data.id)


    // LIST
    const contact_tag_ref01_match = {}

    const contact_tag_ref01_list = await contact_tag_ref01_ent.list(contact_tag_ref01_match)

    assert(!isempty(select(contact_tag_ref01_list, { id: contact_tag_ref01_data.id })))


    // UPDATE
    const contact_tag_ref01_data_up0 = {}
    contact_tag_ref01_data_up0.id = contact_tag_ref01_data.id

    const contact_tag_ref01_markdef_up0 = { name: 'name', value: 'Mark01-contact_tag_ref01_' + setup.now }
    contact_tag_ref01_data_up0 [contact_tag_ref01_markdef_up0.name] = contact_tag_ref01_markdef_up0.value

    const contact_tag_ref01_resdata_up0 = await contact_tag_ref01_ent.update(contact_tag_ref01_data_up0)
    assert(contact_tag_ref01_resdata_up0.id === contact_tag_ref01_data_up0.id)

    assert(contact_tag_ref01_resdata_up0[contact_tag_ref01_markdef_up0.name] === contact_tag_ref01_markdef_up0.value)


    // REMOVE
    const contact_tag_ref01_match_rm0 = {}
    contact_tag_ref01_match_rm0.id = contact_tag_ref01_data.id
    await contact_tag_ref01_ent.remove(contact_tag_ref01_match_rm0)
  

    // LIST
    const contact_tag_ref01_match_rt0 = {}

    const contact_tag_ref01_list_rt0 = await contact_tag_ref01_ent.list(contact_tag_ref01_match_rt0)

    assert(isempty(select(contact_tag_ref01_list_rt0, { id: contact_tag_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/contact_tag/ContactTagTestData.json')

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
    ['contact_tag01','contact_tag02','contact_tag03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'LUMA_TEST_CONTACT_TAG_ENTID': idmap,
    'LUMA_TEST_LIVE': 'FALSE',
    'LUMA_TEST_EXPLAIN': 'FALSE',
    'LUMA_APIKEY': 'NONE',
  })

  idmap = env['LUMA_TEST_CONTACT_TAG_ENTID']

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
  
