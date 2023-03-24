import Game from "./components/game";
import Navigation from "./components/nav";

function App() {
  return (
    <div className="mont-mono bg-zinc-800 text-white lg:h-screen box-border">
      <Navigation />
      <Game />
    </div>
  );
}

export default App;
