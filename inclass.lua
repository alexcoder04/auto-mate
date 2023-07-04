--[[
 - Name: "In Class"
 - On:
   - Action: time
     Args:
      - day: [1, 1, 1, 1, 1, 1, 1]
      - time: "08:00"
]]

DO.notification.doNotDistrub(true)
DO.volume(0)
DO.brightness(34)
DO.network.airplaneMode(false)
DO.network.wifi(true)
DO.network.bluetooth(true)
