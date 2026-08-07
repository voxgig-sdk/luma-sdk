

class LumaError extends Error {

  isLumaError = true

  sdk = 'Luma'

  constructor(code, msg, ctx) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

module.exports = {
  LumaError
}

