
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { LumaSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await LumaSDK.test()
    equal(null !== testsdk, true)
  })

})
