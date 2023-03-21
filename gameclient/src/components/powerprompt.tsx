import { PlayerPower } from "./common";

export function PowerPrompt({
  power,
  handlePowerChoice,
}: {
  power: PlayerPower;
  handlePowerChoice: any;
}) {
  return (
    <div className="status">
      <span>
        Do you want to use the following power: <p>{power.description}</p>
      </span>
      <button
        className="font-bold text-xl px-5 mr-5"
        onClick={() => {
          handlePowerChoice(true);
        }}
      >
        Yes
      </button>
      <button
        className="font-bold text-xl px-5"
        onClick={() => {
          handlePowerChoice(false);
        }}
      >
        No
      </button>
    </div>
  );
}
