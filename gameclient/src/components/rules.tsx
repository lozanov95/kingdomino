export function RulesSection() {
  const paragraphs = [
    {
      header: "Object of the game",
      text: `Choose your dice and combine them to make “dominoes” which are then added to your kingdom to score points. 
            Kingdoms are made up of domains (groups of identical coats of arms, connected either horizontally or vertically) which score points 
            at the end of the game equal to the number of high dignitaries (crosses) in that 
            domain multiplied by the number of its coats of arms.`,
    },
    {
      header: "Connection rules",
      text: `When adding a new domino to your kingdom, you must
            follow at least one of the following two rules:
            Connect one of the symbols
            on your domino to the
            central “castle” space
            (regardless of the coat of
            arms being drawn). Connect at least one of the symbols
            on your domino to a matching
            coat of arms already
            in your kingdom
            (drawn on a
            previous turn).
            Dominoes must be
            connected orthogonally
            (up, down, left, or right).
            Diagonal connections
            do not count. When two coats of arms are side by side, shade the lign
            separating the two squares. This will allow you to better
            visualize your domain.
            If you cannot connect your domino by either of these rules,
            then you draw nothing for this turn.`,
    },
  ];

  return (
    <div>
      {paragraphs.map((paragraph, idx) => {
        return (
          <Paragraph
            key={idx}
            header={paragraph.header}
            text={paragraph.text}
          />
        );
      })}
    </div>
  );
}

function Paragraph({ header, text }: { header: string; text: string }) {
  return (
    <div>
      <p>{header}</p>
      <span>{text}</span>
    </div>
  );
}
