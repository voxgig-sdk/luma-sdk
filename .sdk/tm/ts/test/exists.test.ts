
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { LumaSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await LumaSDK.test()
    equal(null !== testsdk, true)
  })

})
