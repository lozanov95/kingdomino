import { PlayerPower } from "./common";

export function PowerPrompt({
  power,
  handlePowerChoice,
}: {
  power: PlayerPower;
  handlePowerChoice: any;
}) {
  return (
    <div>
      <span>
        Do you want to use the following power: <p>{power.description}</p>
      </span>
      <button
        onClick={() => {
          handlePowerChoice(true);
        }}
      >
        Yes
      </button>
      <button
        onClick={() => {
          handlePowerChoice(false);
        }}
      >
        No
      </button>
    </div>
  );
}
