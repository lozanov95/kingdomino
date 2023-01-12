import { Cell } from "./game"
import { getBadgeIcon } from "./common"
import { getBonus } from "../api/api"
import { useEffect, useState } from "react";

function BonusBoard() {
    const bonus = getBonus();

    return (
        <div className="bonusboard">
            {bonus.map(({ badge, currentChecks, requiredChecks, eligible }, idx) => {
                return <BonusCell key={idx} imgSrc={getBadgeIcon(badge)} currentChecks={currentChecks} requiredChecks={requiredChecks} eligible={eligible} />
            })}
        </div>
    )
}

function BonusCell({ imgSrc, currentChecks, requiredChecks, eligible }: { imgSrc: string, currentChecks: number, requiredChecks: number, eligible: boolean }) {
    const [elements, setElements] = useState<JSX.Element[]>([])
    const [elClass, setElClass] = useState("")

    useEffect(() => {
        setElClass(eligible ? "" : " ineligible")

        const els = Array.from(Array(requiredChecks)).map((_, idx) => {

            if (idx < currentChecks) {
                return <input key={idx} className="bonus-checkbox" type="checkbox" disabled checked />
            }
            return <input key={idx} className="bonus-checkbox" type="checkbox" disabled />
        })

        setElements(els)
    }, [currentChecks, eligible])

    return (
        <div className={"bonus-cell" + elClass}>
            <Cell imgSrc={imgSrc} />
            {elements}
        </div>
    )
}

export default BonusBoard