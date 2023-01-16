import Chat from './components/chat'
import Game from './components/game'
import Navigation from './components/nav'
import BonusBoard from './components/bonusboard'

function App() {
  return (
    <div className="app">
      <Navigation />
      <Game />
      {/* <BonusBoard /> */}
      <Chat />
    </div>
  )
}

export default App
