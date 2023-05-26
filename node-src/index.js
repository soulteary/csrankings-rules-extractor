const { readFileSync } = require("fs");
const GetCSRankingsConfig = require("./lib");
const code = readFileSync("../data/emeryberger/CSrankings/csrankings.ts", "utf-8");

const config = GetCSRankingsConfig(code);
console.log(config);
