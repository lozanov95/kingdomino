import { Cell, getBadgeIcon } from "./common";
import { Scoreboard } from "../helpers/types";
import { ModalWithChildren } from "./modal";
import { SetStateAction } from "react";

export function ScoreSection({ scoreboards, setScoreboards }: { scoreboards: Scoreboard[], setScoreboards: React.Dispatch<SetStateAction<Scoreboard[]>> }) {
  if (scoreboards.length === 0) {
    return <></>
  }



  return (
    <ModalWithChildren>
      <div className="flex justify-evenly flex-col">
        <div className="flex">
          {scoreboards.map((scoreboard, idx) => {
            return <PlayerScoreboard scoreboard={scoreboard} key={idx} />;
          })}
        </div>
        <button className="p-1 px-3 rounded-lg bg-gray-700 hover:bg-gray-600" onClick={() => setScoreboards([])}>Close</button>
      </div>
    </ModalWithChildren>
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
      <input
        className="text-2xl font-bold text-center w-[40px] lg:w-[90px]"
        value={score}
        disabled
      />
    </div>
  );
}
