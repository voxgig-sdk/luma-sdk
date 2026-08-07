-- Luma SDK error

local LumaError = {}
LumaError.__index = LumaError


function LumaError.new(code, msg, ctx)
  local self = setmetatable({}, LumaError)
  self.is_sdk_error = true
  self.sdk = "Luma"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function LumaError:error()
  return self.msg
end


function LumaError:__tostring()
  return self.msg
end


return LumaError
