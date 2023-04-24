export type Dice = {
  name: Badge;
  nobles: number;
};

export type BoardPosition = {
  row: number;
  cell: number;
};

export enum GameTurn {
  Disconnected = 0,
  PickDice,
  PlaceDice,
  HandlePlayerPower,
}

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
  dice: Dice;
  isPicked: boolean;
  playerId: number;
  isPlaced: boolean;
};

export type GameState = {
  id: number;
  gameTurn: GameTurn;
  bonusCard: Bonus[];
  message: string;
  board: Dice[][];
  dices: DiceResult[];
  playerPower: PlayerPower;
  scoreboards: Scoreboard[];
};

export type ServerPayload = {
  name?: string;
  boardPosition?: BoardPosition;
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
