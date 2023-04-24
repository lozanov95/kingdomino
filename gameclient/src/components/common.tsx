import badgeEmpty from "../assets/empty.svg";
import badgeCastle from "../assets/castle.svg";
import badgeChecked from "../assets/checked.svg";
import badgeFilled from "../assets/filled.svg";
import badgeDot from "../assets/dot.svg";
import badgeDoubleDot from "../assets/doubledot.svg";
import badgeLine from "../assets/line.svg";
import badgeDoubleLine from "../assets/doubleline.svg";
import badgeQuestion from "../assets/question.svg";
import { MouseEventHandler, ReactElement } from "react";
import { Badge } from "../helpers/types";

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

export function Nobles({ amount, color }: { amount: number; color?: string }) {
  return (
    <div
      className={[
        "w-[16px] lg:w-[20px] text-center",
        color ?? "bg-gray-600",
      ].join(" ")}
    >
      {renderNobles(amount)}
    </div>
  );
}

export function Noble() {
  return (
    <div className="w-[14px] h-[14px] bg-black rounded-full mt-[2px] ml-[1px] lg:w-[18px] lg:h-[18px]"></div>
  );
}

function renderNobles(amount: number) {
  if (amount === 0) {
    return;
  }

  let nobles: ReactElement[] = [];
  for (let i = 0; i < amount; i++) {
    nobles.push(<Noble key={i} />);
  }

  return <>{nobles}</>;
}
