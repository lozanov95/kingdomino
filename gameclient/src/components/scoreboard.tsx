import { Cell, getBadgeIcon, Scoreboard } from "./common";

export function ScoreSection({ scoreboards }: { scoreboards: Scoreboard[] }) {
  return (
    <div className="flex justify-evenly">
      {scoreboards.map((scoreboard, idx) => {
        return <PlayerScoreboard scoreboard={scoreboard} key={idx} />;
      })}
    </div>
  );
}

export function PlayerScoreboard({ scoreboard }: { scoreboard: Scoreboard }) {
  return (
    <div>
      <p className="text-3xl">{scoreboard.name}</p>
      {scoreboard.scores.map((badgeScore, idx) => {
        return (
          <ScoreRow
            badgeId={badgeScore.badge}
            score={badgeScore.score}
            key={idx}
          />
        );
      })}
      <div className="text-2xl font-bold">
        Total score: {scoreboard.totalScore}
      </div>
    </div>
  );
}

function ScoreRow({ badgeId, score }: { badgeId: number; score: number }) {
  return (
    <div className="flex">
      <Cell id="" imgSrc={getBadgeIcon(badgeId)} />
      <input className="text-2xl font-bold text-center w-[40px] lg:w-[90px]" value={score} disabled />
    </div>
  );
}
