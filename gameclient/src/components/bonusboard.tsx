import { Cell } from "./common";
import { getBadgeIcon } from "./common";
import { Bonus } from "../helpers/types";
import { useEffect, useState, memo } from "react";

const BonusBoard = memo(function BonusBoard({
  bonusCard,
}: {
  bonusCard: Bonus[] | null;
}) {
  return (
    <>
      {bonusCard !== null ? (
        <div className="col-start-5 w-fit mx-auto">
          <h2 className="text-2xl font-bold text-center">Bonuses</h2>
          {bonusCard
            .sort((a: Bonus, b: Bonus) => (a.name > b.name ? 1 : -1))
            .map(({ name, currentChecks, requiredChecks, eligible }, idx) => {
              return (
                <BonusCell
                  key={idx}
                  imgSrc={getBadgeIcon(name)}
                  currentChecks={currentChecks}
                  requiredChecks={requiredChecks}
                  eligible={eligible}
                />
              );
            })}
        </div>
      ) : (
        ""
      )}
    </>
  );
});

function BonusCell({
  imgSrc,
  currentChecks,
  requiredChecks,
  eligible,
}: {
  imgSrc: string;
  currentChecks: number;
  requiredChecks: number;
  eligible: boolean;
}) {
  const [elements, setElements] = useState<JSX.Element[]>([]);
  const [elClass, setElClass] = useState("");

  const checkboxClass = "lg:w-[30px] first-of-type:ml-1";

  useEffect(() => {
    let cs = "";
    if (currentChecks == requiredChecks) {
      cs = "bg-green-800";
    } else if (!eligible) {
      cs = "bg-red-900";
    }
    setElClass(cs);

    const els = Array.from(Array(requiredChecks)).map((_, idx) => {
      if (idx < currentChecks) {
        return (
          <input
            key={idx}
            value=""
            type="checkbox"
            className={checkboxClass}
            disabled
            checked
          />
        );
      }
      return (
        <input
          key={idx}
          value=""
          type="checkbox"
          className={checkboxClass}
          disabled
        />
      );
    });

    setElements(els);
  }, [currentChecks, eligible]);

  return (
    <div className={[elClass, "flex w-full m-1  "].join(" ")}>
      <Cell imgSrc={imgSrc} id={imgSrc} />
      {elements}
    </div>
  );
}

export default BonusBoard;
