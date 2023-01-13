import { Badge, Domino } from "../components/common"

export function getBoard() {
    let board: [Domino, Domino, Domino, Domino, Domino, Domino, Domino][] = [
        [
            { badge: Badge.EMPTY, nobles: 0, }, { badge: Badge.DOT, nobles: 0, }, { badge: Badge.LINE, nobles: 0, }, { badge: Badge.DOUBLEDOT, nobles: 0, },
            { badge: Badge.DOUBLELINE, nobles: 0, }, { badge: Badge.FILLED, nobles: 0, }, { badge: Badge.CHECKED, nobles: 0 }
        ],
        [
            { badge: Badge.EMPTY, nobles: 2, }, { badge: Badge.DOT, nobles: 2, }, { badge: Badge.LINE, nobles: 0, }, { badge: Badge.DOUBLEDOT, nobles: 0, },
            { badge: Badge.DOUBLELINE, nobles: 0, }, { badge: Badge.FILLED, nobles: 0, }, { badge: Badge.CHECKED, nobles: 0 }
        ],
        [
            { badge: Badge.EMPTY, nobles: 0, }, { badge: Badge.DOT, nobles: 0, }, { badge: Badge.LINE, nobles: 0, }, { badge: Badge.CASTLE, nobles: 0, },
            { badge: Badge.DOUBLELINE, nobles: 0, }, { badge: Badge.FILLED, nobles: 1, }, { badge: Badge.CHECKED, nobles: 0 }
        ],
        [
            { badge: Badge.EMPTY, nobles: 1, }, { badge: Badge.DOT, nobles: 0, }, { badge: Badge.LINE, nobles: 0, }, { badge: Badge.DOUBLEDOT, nobles: 0, },
            { badge: Badge.DOUBLELINE, nobles: 0, }, { badge: Badge.FILLED, nobles: 2, }, { badge: Badge.CHECKED, nobles: 0 }
        ],
        [
            { badge: Badge.EMPTY, nobles: 0, }, { badge: Badge.DOT, nobles: 0, }, { badge: Badge.LINE, nobles: 0, }, { badge: Badge.DOUBLEDOT, nobles: 0, },
            { badge: Badge.DOUBLELINE, nobles: 0, }, { badge: Badge.FILLED, nobles: 1, }, { badge: Badge.CHECKED, nobles: 0 }
        ],
    ]
    return board
}

export function getDices() {
    const dices = [
        { badge: Badge.DOT, nobles: 1 },
        { badge: Badge.QUESTIONMARK, nobles: 0 },
        { badge: Badge.DOUBLELINE, nobles: 0 },
        { badge: Badge.CHECKED, nobles: 2 }
    ]

    return dices
}

export function getBonus() {
    const bonus = [
        {
            badge: Badge.DOT, requiredChecks: 5, currentChecks: 4, eligible: true
        },
        {
            badge: Badge.LINE, requiredChecks: 5, currentChecks: 0, eligible: true
        },
        {
            badge: Badge.DOUBLEDOT, requiredChecks: 4, currentChecks: 0, eligible: true
        },
        {
            badge: Badge.DOUBLELINE, requiredChecks: 4, currentChecks: 2, eligible: false
        },
        {
            badge: Badge.CHECKED, requiredChecks: 3, currentChecks: 3, eligible: false
        },
        {
            badge: Badge.FILLED, requiredChecks: 3, currentChecks: 0, eligible: true
        },
    ]

    return bonus
}