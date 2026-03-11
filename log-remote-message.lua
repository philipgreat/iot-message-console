-- 导入 socket 库
local socket = require "socket"

-- 初始化 UDP 客户端
local iotlogclient = assert(socket.udp())
assert(iotlogclient:settimeout(100), "iotlogclient:settimeout(0) fail")
assert(iotlogclient:setsockname('*', 0), "iotlogclient:setsockname('*', 0)")

-- 设置远程主机和端口
local remoteHost = "iotlog.doublechaintech.com"
local remotePort = 54321
local remoteIp = socket.dns.toip(remoteHost)

-- 获取命令行参数
local messageArg = arg[1] or "No message provided"

-- 编码函数：将字符串进行简单的转换
function codec(str)
    local hexTable = {}
    for index = 1, #str do
        local hex = 255 - string.byte(str, index, index)
        table.insert(hexTable, string.char(hex))
    end
    return table.concat(hexTable)
end

-- 记录日志函数
local lastTime = os.time()
function log(msg)
    if (os.time() - lastTime) > 1000 then
        remoteIp = socket.dns.toip(remoteHost) -- 每 1000 秒更新一次 IP 地址，以防 IP 变更
    end
    
    local message = codec(msg)
    iotlogclient:sendto(message, remoteIp, remotePort)
    print("Sent message: " .. msg .. " to " .. remoteIp .. ":" .. remotePort)
end

-- 主程序逻辑
if messageArg and messageArg ~= "" then
    log(messageArg)
else
    print("Usage: lua script_name.lua \"your_message_here\"")
    print("Please provide a message as a command line argument.")
end

