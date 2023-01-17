import { Cell } from "./game"
import { Bonus, getBadgeIcon } from "./common"
import { getBonus } from "../api/api"
import { useEffect, useState } from "react";

function BonusBoard({ bonusCard }: { bonusCard?: Bonus[] }) {

    return (
        <div className="bonusboard">
            {bonusCard?.sort((a: Bonus, b: Bonus) => a.name > b.name ? 1 : -1).map(({ name, CurrentChecks, RequiredChecks, Eligible }, idx) => {
                return <BonusCell key={idx} imgSrc={getBadgeIcon(name)} CurrentChecks={CurrentChecks} RequiredChecks={RequiredChecks} Eligible={Eligible} />
            })}
        </div>
    )
}

function BonusCell({ imgSrc, CurrentChecks, RequiredChecks, Eligible }: { imgSrc: string, CurrentChecks: number, RequiredChecks: number, Eligible: boolean }) {
    const [elements, setElements] = useState<JSX.Element[]>([])
    const [elClass, setElClass] = useState("")

    useEffect(() => {
        let cs = ""
        if (CurrentChecks == RequiredChecks) {
            cs = " completed"
        } else if (!Eligible) {
            cs = " ineligible"
        }
        setElClass(cs)

        const els = Array.from(Array(RequiredChecks)).map((_, idx) => {

            if (idx < CurrentChecks) {
                return <input key={idx} className="bonus-checkbox" type="checkbox" disabled checked />
            }
            return <input key={idx} className="bonus-checkbox" type="checkbox" disabled />
        })

        setElements(els)
    }, [CurrentChecks, Eligible])

    return (
        <div className={"bonus-cell" + elClass}>
            <Cell imgSrc={imgSrc} />
            {elements}
        </div>
    )
}

export default BonusBoard