# API
# The api library has 3 api layers

# /events/
# Events api allows users to send predefine events to other systems, i.e. user login/logout, 
# If the event sent by the user matches all the required fields heraldo will translate the 
# event into the proper api and send it to the target system, events api ensures that the 
# message is relayed into the proper format. This api is designed after TCP protocol

# /notify/
# Notifications api allows users to just push forward any notifications messages to another 
# systems, it wont check/verify the message has the proper fields and just do the maximum
# effor to forward the notification to its destination. The process and check of the actual
# message is done by the target system. This api is designed after UDP protocol

# /mgmt/
# Management api allows user to CRUD new events and enable/disable/configure sender plugins

# /monitor/
# The query api allows user to get stats of how many messages are pushed by heraldo as well as 
# obtaining stats from the system


