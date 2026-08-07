
import { Context } from './Context'


class LumaError extends Error {

  isLumaError = true

  sdk = 'Luma'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  LumaError
}

