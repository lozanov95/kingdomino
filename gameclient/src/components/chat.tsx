import { EventHandler, FormEventHandler, useEffect, useState } from "react"

function Chat() {
    const [connected, setConnected] = useState(false)
    const [wsConnection, setWsConnection] = useState<WebSocket | null>(null)

    function handleChatConnect() {
        connect()
    }

    const connect = () => {
        const ws = new WebSocket("ws://localhost:8080/ws")
        ws.onopen = () => {
            setConnected(true)
            setWsConnection(ws)
        }

        ws.onclose = () => {
            setConnected(false)
        }

        ws.onerror = () => {
            setConnected(false)
        }
    }

    function sendMessage(msg: string) {
        wsConnection?.send(msg)
        console.log("sent ", msg)
    }

    return (
        <div>
            <h1>Chat</h1>
            <button hidden={connected} onClick={handleChatConnect}>Connect</button>
            <MessageBox handleSend={sendMessage} hidden={!connected} />
        </div>
    )
}

function MessageBox(props: { handleSend: (msg: string) => void, hidden: boolean }) {
    const [msg, setMsg] = useState("")

    function handleSubmit(ev: any) {
        ev.preventDefault()
        props.handleSend(msg)
        setMsg("")
    }

    return (
        <form onSubmit={e => handleSubmit(e)} hidden={props.hidden}>
            <input value={msg} onChange={e => setMsg(e.target.value)} />
            <button>Send</button>
        </form>

    )
}

export default Chat

