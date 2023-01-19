import { Cell } from "./common"
import { Bonus, getBadgeIcon } from "./common"
import { useEffect, useState, memo } from "react";


const BonusBoard = memo(
    function BonusBoard({ bonusCard }: { bonusCard?: Bonus[] }) {

        return (
            <div className="bonusboard">
                {bonusCard?.sort((a: Bonus, b: Bonus) => a.name > b.name ? 1 : -1).map(({ name, currentChecks, requiredChecks, eligible }, idx) => {
                    return <BonusCell key={idx} imgSrc={getBadgeIcon(name)} currentChecks={currentChecks} requiredChecks={requiredChecks} eligible={eligible} />
                })}
            </div>
        )
    }
)

function BonusCell({ imgSrc, currentChecks, requiredChecks, eligible }: { imgSrc: string, currentChecks: number, requiredChecks: number, eligible: boolean }) {
    const [elements, setElements] = useState<JSX.Element[]>([])
    const [elClass, setElClass] = useState("")

    useEffect(() => {
        let cs = ""
        if (currentChecks == requiredChecks) {
            cs = " completed"
        } else if (!eligible) {
            cs = " ineligible"
        }
        setElClass(cs)

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
            <Cell imgSrc={imgSrc} id="" />
            {elements}
        </div>
    )
}

export default BonusBoard