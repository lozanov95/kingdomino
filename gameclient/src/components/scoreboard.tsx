import { getBadgeIcon, Scoreboard } from "./common";

export function ScoreSection({ scoreboards }: { scoreboards: Scoreboard[] }) {
  return (
    <div>
      {scoreboards.map((scoreboard, idx) => {
        return <PlayerScoreboard scoreboard={scoreboard} key={idx} />;
      })}
    </div>
  );
}

export function PlayerScoreboard({ scoreboard }: { scoreboard: Scoreboard }) {
  return (
    <div>
      <p>{scoreboard.name}</p>
      {scoreboard.scores.map((badgeScore, idx) => {
        return (
          <ScoreRow
            badgeId={badgeScore.badge}
            score={badgeScore.score}
            key={idx}
          />
        );
      })}
      <div>
        <span>Total score: {scoreboard.totalScore}</span>
      </div>
    </div>
  );
}

function ScoreRow({ badgeId, score }: { badgeId: number; score: number }) {
  return (
    <div>
      <img src={getBadgeIcon(badgeId)} />
      <input value={score} disabled />
    </div>
  );
}
