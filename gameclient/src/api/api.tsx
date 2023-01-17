import { Badge, Domino } from "../components/common"

export function getBoard() {
    let board: [Domino, Domino, Domino, Domino, Domino, Domino, Domino][] = [
        [
            { name: Badge.EMPTY, nobles: 0, }, { name: Badge.DOT, nobles: 0, }, { name: Badge.LINE, nobles: 0, }, { name: Badge.DOUBLEDOT, nobles: 0, },
            { name: Badge.DOUBLELINE, nobles: 0, }, { name: Badge.FILLED, nobles: 0, }, { name: Badge.CHECKED, nobles: 0 }
        ],
        [
            { name: Badge.EMPTY, nobles: 2, }, { name: Badge.DOT, nobles: 2, }, { name: Badge.LINE, nobles: 0, }, { name: Badge.DOUBLEDOT, nobles: 0, },
            { name: Badge.DOUBLELINE, nobles: 0, }, { name: Badge.FILLED, nobles: 0, }, { name: Badge.CHECKED, nobles: 0 }
        ],
        [
            { name: Badge.EMPTY, nobles: 0, }, { name: Badge.DOT, nobles: 0, }, { name: Badge.LINE, nobles: 0, }, { name: Badge.CASTLE, nobles: 0, },
            { name: Badge.DOUBLELINE, nobles: 0, }, { name: Badge.FILLED, nobles: 1, }, { name: Badge.CHECKED, nobles: 0 }
        ],
        [
            { name: Badge.EMPTY, nobles: 1, }, { name: Badge.DOT, nobles: 0, }, { name: Badge.LINE, nobles: 0, }, { name: Badge.DOUBLEDOT, nobles: 0, },
            { name: Badge.DOUBLELINE, nobles: 0, }, { name: Badge.FILLED, nobles: 2, }, { name: Badge.CHECKED, nobles: 0 }
        ],
        [
            { name: Badge.EMPTY, nobles: 0, }, { name: Badge.DOT, nobles: 0, }, { name: Badge.LINE, nobles: 0, }, { name: Badge.DOUBLEDOT, nobles: 0, },
            { name: Badge.DOUBLELINE, nobles: 0, }, { name: Badge.FILLED, nobles: 1, }, { name: Badge.CHECKED, nobles: 0 }
        ],
    ]
    return board
}

export function getDices() {
    const dices = [
        { name: Badge.DOT, nobles: 1 },
        { name: Badge.QUESTIONMARK, nobles: 0 },
        { name: Badge.DOUBLELINE, nobles: 0 },
        { name: Badge.CHECKED, nobles: 2 }
    ]

    return dices
}

export function getBonus() {
    const bonus = [
        {
            name: Badge.DOT, requiredChecks: 5, currentChecks: 4, eligible: true
        },
        {
            name: Badge.LINE, requiredChecks: 5, currentChecks: 0, eligible: true
        },
        {
            name: Badge.DOUBLEDOT, requiredChecks: 4, currentChecks: 0, eligible: true
        },
        {
            name: Badge.DOUBLELINE, requiredChecks: 4, currentChecks: 2, eligible: false
        },
        {
            name: Badge.CHECKED, requiredChecks: 3, currentChecks: 3, eligible: false
        },
        {
            name: Badge.FILLED, requiredChecks: 3, currentChecks: 0, eligible: true
        },
    ]

    return bonus
}