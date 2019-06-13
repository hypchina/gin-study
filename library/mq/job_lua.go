package mq

const publishLuaScript = `
		local exists
		local insertFlag
		local jobDetailKey = tostring(KEYS[1])
		local jobDetailVal = tostring(ARGV[1])
		local jobDetailTTL = tonumber(ARGV[2])
		local jobListKey   = tostring(KEYS[2])
		local jobListVal   = tostring(ARGV[3])
		local jobDelayKey    = tostring(KEYS[3])
		local jobDelayVal    = tostring(ARGV[4])
		local jobDelayTime   = tonumber(ARGV[5])
		local jobAt			 = tonumber(ARGV[6])
		exists = redis.call("SET",jobDetailKey,jobDetailVal,"NX","EX",jobDetailTTL)
		if  type(exists) == 'table' then
			if exists["ok"] == "OK" then
				if jobDelayTime < 1 then
					insertFlag = redis.call("LPUSH",jobListKey,jobListVal)
				else
					insertFlag = redis.call("ZADD",jobDelayKey,jobAt,jobDelayVal)
				end
				if insertFlag >0 then
					return "ok"
				end
				redis.call("DEL",jobDetailKey)
				return "error:jobInsertFlag is error"
			end
		end
			return "error:jobId is exists"
	`
