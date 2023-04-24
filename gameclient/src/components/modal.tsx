export function ModalPrompt({
  prompt,
  onClick,
}: {
  prompt: string;
  onClick: (selection: boolean) => void;
}) {
  return (
    <div className="absolute w-full h-full bg-gray-800 opacity-90 top-0 left-0">
      <div className="bg-black h-fit p-3 rounded-lg m-auto fixed top-[50%] left-[50%] translate-y-[-50%] translate-x-[-50%]">
        <p className="text-center text-xl font-bold">{prompt}</p>
        <div className="text-center">
          <button
            onClick={() => {
              onClick(true);
            }}
            className="bg-blue-800 py-1 px-5 rounded-lg mx-2 mt-3"
          >
            Yes
          </button>
          <button
            onClick={() => {
              onClick(false);
            }}
            className="bg-red-800 py-1 px-5 rounded-lg"
          >
            No
          </button>
        </div>
      </div>
    </div>
  );
}

export default ModalPrompt;
