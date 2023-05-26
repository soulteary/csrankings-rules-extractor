const { readFileSync } = require("fs");
const code = readFileSync("../data/emeryberger/CSrankings/csrankings.ts", "utf-8");
const GetCSRankingsConfig = require("./lib");

const config = GetCSRankingsConfig(code);
console.log(config);
