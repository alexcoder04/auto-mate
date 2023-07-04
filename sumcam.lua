--[[
 - Name: "Summarize image"
 - On: []
]]

photo = GET.camera.photo()
text = DO.api.img2text.img2text(photo)
if text ~= "" then
    DO.control.continue("This is the text: " + text + ". Continue?")
    summary = DO.api.llm.ask("Summarize the following text: " + text)
    DO.notification.show(summary)
    DO.control.menu({
        {
            title = "Copy",
            action = function() DO.clipboard.copy(summary) end
        },
        {
            title = "Nothing",
            action = function() end
        }
    })
else
    DO.notification.show("There is no text in the image")
end
