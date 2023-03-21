import Game from './components/game'
import Navigation from './components/nav'
import { RulesSection } from './components/rules'

function App() {
  return (
    <div className="app">
      <Navigation />
      <Game />
      {/* <RulesSection /> */}
    </div>
  )
}

export default App
