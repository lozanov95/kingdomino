import { useEffect, useState } from "react"

function Chat() {
    const [connected, setConnected] = useState(false)

    function handleChatConnect() {
        connect()
    }

    const connect = () => {
        const ws = new WebSocket("ws://localhost:8080/ws")
        ws.onopen = () => {
            setConnected(true)
        }

        ws.onclose = () => {
            setConnected(false)
        }

        ws.onerror = () => {
            setConnected(false)
        }
    }

    return (
        <div>
            <h1>Chat</h1>
            <button hidden={connected} onClick={handleChatConnect}>Connect</button>
            <MessageBox />
        </div>
    )
}

function MessageBox() {
    return (
        <div>
            <input />
            <button>Send</button>
        </div>

    )
}

export default Chat