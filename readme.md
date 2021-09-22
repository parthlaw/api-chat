## Connect to websocket server
### URL: /ws
### Protocol: wss/ws
### Headers: {Authorization:"access-token"}

<br/>

## Join Room
### Actions:
- "join-room" <br/>
- "join-room-private"
- "join-room-two-way"
### Syntax:
> {\
    action:"action"\
    message:"Room-name (can be the id stored in database or being used by client)"\
}

<br/>

## Send Message
### Actions:
- "send-message"
- "typing"
### Syntax:
> {\
    action:"send-message"\
    message:"Message-Text"\
    target:{\
        ID:"websocket id returned after joining the room"\
        name:"Name sent while joining the room"
    }\
}