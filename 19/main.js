const fs = require("fs");
fs.readFile("input.txt", (err, data) => {
  if (err) throw err;
  const puzzleInput = data.toString().split("\n");
  const allRules = {};
  const rules = puzzleInput.slice(0, puzzleInput.indexOf(""));
  rules.forEach((rule) => {
    const key = rule.split(":")[0];
    const value = rule.split(":")[1].trim().replace(/\"/g, "");
    allRules[key] = value;
  });
  // console.log(allRules);
  const messages = puzzleInput.slice(puzzleInput.indexOf("") + 1);
  // const regex = new RegExp(`^${findRegex(allRules["0"])}$`);
  allRules["0"] = "8 11";
  allRules["8"] = "42 | 42 8";
  allRules["11"] = "42 31 | 42 11 31";
  const rule = new RegExp(
    `^(?<group42>(${findRegex(allRules["42"])})+)(?<group31>(${findRegex(
      allRules["31"]
    )})+)$`
  );
  let matchesCount = 0;
  messages.forEach((m) => {
    const matches = m.match(rule);
    if (matches) {
      const { groups } = matches;
      const matches42 = groups.group42.match(
        new RegExp(findRegex(allRules[42]), 'g')
      ).length;
      const matches31 = groups.group31.match(
        new RegExp(findRegex(allRules[31]), 'g')
      ).length;
      if (matches42 > matches31) {
        matchesCount++;
      }
    }
  });
  console.log({ matchesCount });
  // let matchesCount = 0;
  // messages.forEach((m) => {
  //   if (regex.test(m)) {
  //     matchesCount++;
  //   }
  // });
  function findRegex(specificRules) {
    if (specificRules === "a" || specificRules === "b") {
      return specificRules;
    }
    let message = "";
    if (specificRules.includes("|")) {
      const subRules = specificRules.split("|").map((r) => r.trim());
      message += `((${findRegex(subRules[0])})|(${findRegex(subRules[1])}))`;
    } else {
      specificRules.split(/\s/).forEach((r) => {
        message += findRegex(allRules[r]);
      });
    }
    return message;
  }
});