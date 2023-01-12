import { useState } from 'react'
import reactLogo from './assets/react.svg'
import Chat from './components/chat'
import Game from './components/game'
import Navigation from './components/nav'

function App() {


  return (
    <div className="app">
      <Navigation />
      <Game />
      <Chat />
    </div>
  )
}

export default App
