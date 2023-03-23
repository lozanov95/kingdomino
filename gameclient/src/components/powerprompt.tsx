import { PlayerPower } from "./common";

export function PowerPrompt({
  power,
  handlePowerChoice,
}: {
  power: PlayerPower;
  handlePowerChoice: any;
}) {
  return (
    <div className="row-start-1 col-start-2 col-end-4 bg-neutral-900 h-fit p-3 rounded-lg m-auto fixed top-[50%] left-[50%] translate-y-[-50%] translate-x-[-50%] opacity-90">
      <div >
        <p className="text-center text-xl font-bold">
          Do you want to use the following power:
        </p>
        <p className="text-justify indent-5">{power.description}</p>
      </div>
      <div className="text-center">
        <button
          onClick={() => {
            handlePowerChoice(true);
          }}
          className="bg-blue-800 py-1 px-5 rounded-lg mx-2 mt-3"
        >
          Yes
        </button>
        <button
          onClick={() => {
            handlePowerChoice(false);
          }}
          className="bg-red-800 py-1 px-5 rounded-lg"
        >
          No
        </button>
      </div>
    </div>
  );
}
