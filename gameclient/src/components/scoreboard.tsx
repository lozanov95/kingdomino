import { getBadgeIcon, Scoreboard } from "./common";

export function ScoreSection({ scoreboards }: { scoreboards: Scoreboard[] }) {
    return (
        <div className="flex m-auto">
            {scoreboards.map((scoreboard, idx) => {
                return <PlayerScoreboard scoreboard={scoreboard} key={idx} />
            })}
        </div>
    );
}

export function PlayerScoreboard({ scoreboard }: { scoreboard: Scoreboard }) {
    return (
        <div className="container bg-gray-700 m-2 rounded-lg max-w-max">
            <p className="text-4xl font-bold m-2">{scoreboard.name}</p>
            {scoreboard.scores.map((badgeScore, idx) => {
                return <ScoreRow badgeId={badgeScore.badge} score={badgeScore.score} key={idx} />
            })}
            <div className="m-2">
                <span className="text-5xl font-bold">Total score: {scoreboard.totalScore}</span>
            </div>
        </div>
    );
}

function ScoreRow({ badgeId, score }: { badgeId: number; score: number }) {
    return (
        <div className="flex justify-center">
            <img className="cell rounded-none" src={getBadgeIcon(badgeId)} />
            <input
                className="cell text-5xl text-center rounded-none"
                value={score}
                disabled
            />
        </div>
    );
}
