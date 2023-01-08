const unAuthElements = document.querySelectorAll(".unAuth")
const authElements = document.querySelectorAll(".auth")
let joinedChat = false

const updateHiddenState = ()=>{
    console.log(joinedChat)
    if (joinedChat){
        unAuthElements.forEach((el)=>{
            el.classList.add("hidden")
        })
        authElements.forEach((el)=>{
            el.classList.remove("hidden")
        })
        
        
    }else{
        authElements.forEach((el) =>{
            el.classList.add("hidden")
        })
        unAuthElements.forEach((el)=>{
            el.classList.remove("hidden")
        })
    }
}

document.querySelector("#joinChat").addEventListener("click", ()=>{
    const ws = new WebSocket("ws://localhost:8080/ws")
    console.log("joining chat")
    ws.onopen = ()=>{
        const msgDiv = document.querySelector("#msgDiv")
        const chatForm = document.querySelector("#chatForm")
        const msg = document.querySelector("#msg")
        
        joinedChat = true
        updateHiddenState()
        msgDiv.innerHTML = ""

        chatForm.addEventListener("submit", (ev)=>{
            ev.preventDefault()
            if (msg.value === "")
                return
            
            ws.send(msg.value)
            let p = document.createElement("p")
            p.textContent = "You: " + msg.value
            msg.value = ""
            msgDiv.appendChild(p)
        })

        ws.onmessage=(ev) =>{
            let p  = document.createElement("p")
            p.textContent = "Other player: " + ev.data
            msgDiv.appendChild(p)
        }      
    }

    ws.onclose = ()=>{
        joinedChat = false
        updateHiddenState()
    }

    ws.onerror = () =>{
        joinedChat = false
        updateHiddenState()
    }
})


window.onload = ()=>{
    updateHiddenState()
}
