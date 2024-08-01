const xmpp = require("simple-xmpp");
//hussein
xmpp.on ("online", data => {
    console.log("Hey you are online! ")
    console.log(`Connected as ${data.jid.user}`)
    send();
})

function send () {
    setTimeout(send, 5000);
    xmpp.send("admin@domin", `hi! ${Date.now()}`)
}
xmpp.on("error", error => 
    console.log(`something went wrong!${error} `))

xmpp.on("chat", (from, message)=>{
    console.log(`Got a message! ${message} from ${from}`)
})

xmpp.connect({
    "jid": "user@domain",
    "password": "user",
    "host": "localhost",
    "port": 5222
})