import { Cell } from "./common";
import { Bonus, getBadgeIcon } from "./common";
import { useEffect, useState, memo } from "react";

const BonusBoard = memo(function BonusBoard({
  bonusCard,
}: {
  bonusCard: Bonus[] | null;
}) {


  return (
    <>
      {bonusCard !== null ?
        <div className="col-start-5 w-fit m-auto">
          <h2>Bonuses</h2>
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
        : ""}
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

  const checkboxClass = "lg:w-[30px]";

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
    <div
      className={[
        elClass,
        "flex w-full pr-2 border-2 border-b-0 border-solid last-of-type:border-b-2",
      ].join(" ")}
    >
      <Cell imgSrc={imgSrc} id="" className="border-r-2" />
      {elements}
    </div>
  );
}

export default BonusBoard;
