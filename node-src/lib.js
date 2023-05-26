const ts = require("typescript");

function getAST(code) {
  const fileName = "csrankings-rules-extractor.ts";
  const compilerHost = {
    fileExists: () => true,
    getCanonicalFileName: (filename) => filename,
    getCurrentDirectory: () => "",
    getDefaultLibFileName: () => "lib.d.ts",
    getNewLine: () => "\n",
    getSourceFile: (filename) => ts.createSourceFile(filename, code, ts.ScriptTarget.Latest, true),
    readFile: () => null,
    useCaseSensitiveFileNames: () => true,
    writeFile: () => null,
  };
  const program = ts.createProgram([fileName], { noResolve: true, target: ts.ScriptTarget.Latest, experimentalDecorators: true, experimentalAsyncFunctions: true }, compilerHost);
  return program.getSourceFile(fileName);
}

function getList(member) {
  if (member.initializer.kind != ts.SyntaxKind.ArrayLiteralExpression) {
    return [];
  }
  return member.initializer.elements.map((element) => {
    let value = null;
    switch (element.kind) {
      case ts.SyntaxKind.TrueKeyword:
        value = true;
        break;
      case ts.SyntaxKind.FalseKeyword:
        value = false;
        break;
      case ts.SyntaxKind.StringLiteral:
        value = element.text;
        break;
      default:
        console.log(element.kind);
    }
    return value;
  });
}

function getListMap(member) {
  if (member.initializer.kind != ts.SyntaxKind.ArrayLiteralExpression) return [];
  return member.initializer.elements
    .filter((element) => element.kind == ts.SyntaxKind.ObjectLiteralExpression)
    .map((element) => {
      return element.properties
        .map((prop) => [prop.name.escapedText, prop.initializer.text])
        .reduce((prev, [key, value]) => {
          prev[key] = value;
          return prev;
        }, {});
    });
}

function getInt(member) {
  if (member.initializer && member.initializer.text !== "" && member.initializer.kind == ts.SyntaxKind.NumericLiteral) {
    return parseInt(member.initializer.text);
  }
  return -1;
}

function getMap(member) {
  var elements = member.initializer.properties;
  return elements
    .map((e) => {
      let key = e.name.text;
      let value = null;
      switch (e.initializer.kind) {
        case ts.SyntaxKind.TrueKeyword:
          value = true;
          break;
        case ts.SyntaxKind.FalseKeyword:
          value = false;
          break;
        case ts.SyntaxKind.StringLiteral:
          value = e.initializer.text;
          break;
        default:
          console.log(e.initializer.kind);
      }
      return [key, value];
    })
    .reduce((dest, [key, value]) => {
      dest[key] = value;
      return dest;
    }, {});
}

function GetMainClass(sourceFile) {
  let classDeclaration = sourceFile.statements.filter((statement) => statement.kind === ts.SyntaxKind.ClassDeclaration);
  if (classDeclaration.length === 0) return null;
  classDeclaration = classDeclaration[0];
  if (!classDeclaration.name || classDeclaration.name.escapedText != "CSRankings") return null;
  if (!classDeclaration.members || classDeclaration.members.length === 0) return null;
  return classDeclaration;
}

function GetDataFromClass(mainClass, found, processFn) {
  let member = mainClass.members.filter((member) => member.name && member.name.escapedText == found);
  if (member.length === 0) return null;
  return processFn(member[0]);
}

function GetCSRankingsConfig(code) {
  const mainClass = GetMainClass(getAST(code));
  if (mainClass === null) return "";
  return JSON.stringify({
    MinToRank: GetDataFromClass(mainClass, "minToRank", getInt),
    Regions: GetDataFromClass(mainClass, "regions", getList),
    AIAreas: GetDataFromClass(mainClass, "aiAreas", getList),
    SystemsAreas: GetDataFromClass(mainClass, "systemsAreas", getList),
    TheoryAreas: GetDataFromClass(mainClass, "theoryAreas", getList),
    InterdisciplinaryAreas: GetDataFromClass(mainClass, "interdisciplinaryAreas", getList),
    AreaMap: GetDataFromClass(mainClass, "areaMap", getListMap),
    ParentMap: GetDataFromClass(mainClass, "parentMap", getMap),
    NextTier: GetDataFromClass(mainClass, "nextTier", getMap),
    NoteMap: GetDataFromClass(mainClass, "noteMap", getMap),
  });
}

module.exports = GetCSRankingsConfig;
