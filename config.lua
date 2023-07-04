--[[
 - Name: "My Sequence"
 - On:
   - Action: network
     Args:
      - Status: enabled
      - Name: "FRITZ!Box 7520 JH"
]]

if GET.battery.Level() < 60 then
    DO.notification.show({
        title = "Battery under 60%"
    })
else
    DO.notification.show({
        title = "Battery above 60%"
    })
end
