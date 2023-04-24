import { PlayerPower } from "./common";
import Modal from "./modal";

export function PowerPrompt({
  power,
  handlePowerChoice,
}: {
  power: PlayerPower;
  handlePowerChoice: (choice: boolean) => void;
}) {
  return <Modal onClick={handlePowerChoice} prompt={power.description} />;
}
