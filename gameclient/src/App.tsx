import Game from "./components/game";
import Navigation from "./components/nav";

function App() {
  return (
    <div className="mont-mono bg-zinc-800 text-white h-full lg:h-screen border-box">
      <Navigation />
      <Game />
    </div>
  );
}

export default App;
