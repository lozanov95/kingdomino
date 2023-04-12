import badgeEmpty from "../assets/empty.svg";
import badgeCastle from "../assets/castle.svg";
import badgeChecked from "../assets/checked.svg";
import badgeFilled from "../assets/filled.svg";
import badgeDot from "../assets/dot.svg";
import badgeDoubleDot from "../assets/doubledot.svg";
import badgeLine from "../assets/line.svg";
import badgeDoubleLine from "../assets/doubleline.svg";
import badgeQuestion from "../assets/question.svg";
import { MouseEventHandler } from "react";

export type Domino = {
  name: Badge;
  nobles: number;
};

export type Bonus = {
  name: number;
  requiredChecks: number;
  currentChecks: number;
  eligible: boolean;
};

export type BadgeScore = {
  badge: number;
  score: number;
};

export type Scoreboard = {
  name: string;
  scores: BadgeScore[];
  totalScore: number;
};

export type DiceResult = {
  dice: Domino;
  isSelected: boolean;
  playerId: number;
}

export type GameState = {
  bonusCard: Bonus[];
  message: string;
  board: Domino[][];
  dices: DiceResult[];
  playerPower: PlayerPower;
  scoreboards: Scoreboard[];
};

export type ServerPayload = {
  name?: string;
  boardPosition?: {
    row: number;
    cell: number;
  };
  selectedDie?: number;
  playerPower?: PlayerPower;
};

export enum Badge {
  EMPTY = 0,
  CASTLE,
  DOT,
  LINE,
  DOUBLEDOT,
  DOUBLELINE,
  FILLED,
  CHECKED,
  QUESTIONMARK,
}

export type PlayerPower = {
  type: number;
  description: string;
  use: boolean;
  confirmed: boolean;
};

export function getBadgeIcon(id: number) {
  switch (id) {
    case Badge.EMPTY:
      return badgeEmpty;
    case Badge.CASTLE:
      return badgeCastle;
    case Badge.DOT:
      return badgeDot;
    case Badge.LINE:
      return badgeLine;
    case Badge.DOUBLEDOT:
      return badgeDoubleDot;
    case Badge.DOUBLELINE:
      return badgeDoubleLine;
    case Badge.FILLED:
      return badgeFilled;
    case Badge.CHECKED:
      return badgeChecked;
    case Badge.QUESTIONMARK:
      return badgeQuestion;
    default:
      return "";
  }
}

export function Cell({
  id,
  imgSrc,
  onClick,
  className,
}: {
  id: string;
  imgSrc: string;
  onClick?: MouseEventHandler;
  className?: string;
}) {
  return (
    <div
      className={[
        "w-[40px] h-[40px] lg:w-[90px] lg:h-[90px] bg-zinc-600",
        className,
      ].join(" ")}
      id={id}
      onClick={onClick}
    >
      <img src={imgSrc} alt="badge icon" className="" />
    </div>
  );
}

export function Nobles({ amount, color }: { amount: number, color?: string }) {
  function renderNobles() {
    switch (amount) {
      case 0:
        return;
      case 1:
        return <Noble />;
      case 2:
        return (
          <>
            <Noble />
            <Noble />
          </>
        );
      default:
        return null;
    }
  }
  return (
    <div className={["w-[16px] lg:w-[20px] text-center", color ?? "bg-gray-600"].join(" ")}>
      {renderNobles()}
    </div>
  );
}

export function Noble() {
  return (
    <div className="w-[14px] h-[14px] bg-black rounded-full m-0.5 lg:w-[18px] lg:h-[18px]"></div>
  );
}
