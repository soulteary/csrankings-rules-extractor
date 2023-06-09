var ast = JSON.parse(data);
ast = ast.Module.body.filter((item) => item.Assign).map((item) => item.Assign);

var ret = [];
var allowList = [
  "pageCountThreshold",
  "pageCounterNormal",
  "pageCounterColon",
  "TECSCounterColon",
  "ISMBpageCounter",
  "areadict",
  "EMSOFT_TECS",
  "EMSOFT_TECS_PaperNumbers",
  "EMSOFT_TCAD",
  "EMSOFT_TCAD_PaperStart",
  "DAC_TooShortPapers",
  "ISMB_Bioinformatics",
  "TOG_SIGGRAPH_Volume",
  "TOG_SIGGRAPH_Asia_Volume",
  "CGF_EUROGRAPHICS_Volume",
  "TVCG_Vis_Volume",
  "TVCG_VR_Volume",
  "ICSE_ShortPaperStart",
  "SIGMOD_NonResearchPaperStart",
  "SIGMOD_NonResearchPapersRange",
  "ASE_LongPaperThreshold",
  "startyear",
  "endyear",
];
ast.filter((item) => allowList.includes(item.targets[0].Name.id)).forEach((item) => ret.push(item));

function GetInt(key) {
  return ret.filter((item) => item.targets[0].Name.id === key)[0].value.Num.n;
}

function GetRegexp(key) {
  const value = ret.filter((item) => item.targets[0].Name.id === key)[0].value;
  if (value.Call.func.Attribute.value.Name.id !== "re") return "";
  return value.Call.args[0].Str.s;
}

function GetDictTuple(key) {
  const value = ret.filter((item) => item.targets[0].Name.id === key)[0].value;
  if (!value.Dict.keys.length) return {};
  return value.Dict.keys.reduce((prev, item, idx) => {
    const [value1, value2] = value.Dict.values[idx].Tuple.elts;
    if (value2.Str) prev[item.Num.n] = [value1.Num.n, value2.Str.s];
    if (value2.Num) prev[item.Num.n] = [value1.Num.n, value2.Num.n];
    return prev;
  }, {});
}

function GetObjectArray(key) {
  const value = ret.filter((item) => item.targets[0].Name.id === key)[0].value;
  if (!value.Dict.keys.length) return {};
  return value.Dict.keys.reduce((prev, item, idx) => {
    const values = value.Dict.values[idx].Num.n;
    prev[item.Num.n] = values;
    return prev;
  }, {});
}

function GetDictArray(key) {
  const value = ret.filter((item) => item.targets[0].Name.id === key)[0].value;
  if (!value.Dict.keys.length) return {};
  return value.Dict.keys.reduce((prev, item, idx) => {
    const values = value.Dict.values[idx]["Set"].elts.map((item) => item.Num.n);
    prev[item.Num.n] = values;
    return prev;
  }, {});
}

function GetDictTupleArray(key) {
  const value = ret.filter((item) => item.targets[0].Name.id === key)[0].value;
  if (!value.Dict.keys.length) return {};
  return value.Dict.keys.reduce((prev, item, idx) => {
    const values = value.Dict.values[idx].List.elts.map((item) => item.Tuple.elts);
    prev[item.Num.n] = values.map((item) => {
      let [value1, value2] = item;
      if (value1.Str) value1 = value1.Str.s;
      if (value1.Num) value1 = value1.Num.n;
      if (value2.Str) value2 = value2.Str.s;
      if (value2.Num) value2 = value2.Num.n;
      return [value1, value2];
    });
    return prev;
  }, {});
}

function ConvertToArray(data) {
  return Object.keys(data).map(function (key) {
    return { key: Number(key), value: data[key] };
  });
}

function FixISMB(datasets) {
  return datasets.map((item) => {
    item.value = item.value.map((value) => {
      if (typeof value == "string") {
        let v = value.toLowerCase();
        if (v == "supplement") {
          value = -1;
        } else if (v == "supplement-1") {
          value = -2;
        } else {
          throw new Error("ISMB_Bioinformatics value error");
        }
      }
      return value;
    });
    return item;
  });
}

function FixTECS(datasets) {
  return datasets.map((item) => {
    item.value = item.value.map((value) => {
      if (typeof value == "string" && value == "5s") {
        value = 5;
      }
      return value;
    });
    return item;
  });
}

JSON.stringify({
  pageCountThreshold: GetInt("pageCountThreshold"),
  pageCounterNormal: GetRegexp("pageCounterNormal"),
  pageCounterColon: GetRegexp("pageCounterColon"),
  TECSCounterColon: GetRegexp("TECSCounterColon"),
  ISMBpageCounter: GetRegexp("ISMBpageCounter"),
  EMSOFT_TECS: FixTECS(ConvertToArray(GetDictTuple("EMSOFT_TECS"))),
  EMSOFT_TECS_PaperNumbers: ConvertToArray(GetDictTuple("EMSOFT_TECS_PaperNumbers")),
  EMSOFT_TCAD: ConvertToArray(GetDictTuple("EMSOFT_TCAD")),
  EMSOFT_TCAD_PaperStart: ConvertToArray(GetDictArray("EMSOFT_TCAD_PaperStart")),
  DAC_TooShortPapers: ConvertToArray(GetDictArray("DAC_TooShortPapers")),
  ISMB_Bioinformatics: FixISMB(ConvertToArray(GetDictTuple("ISMB_Bioinformatics"))),
  TOG_SIGGRAPH_Volume: ConvertToArray(GetDictTuple("TOG_SIGGRAPH_Volume")),
  TOG_SIGGRAPH_Asia_Volume: ConvertToArray(GetDictTuple("TOG_SIGGRAPH_Asia_Volume")),
  CGF_EUROGRAPHICS_Volume: ConvertToArray(GetDictTuple("CGF_EUROGRAPHICS_Volume")),
  TVCG_Vis_Volume: ConvertToArray(GetDictTuple("TVCG_Vis_Volume")),
  TVCG_VR_Volume: ConvertToArray(GetDictTuple("TVCG_VR_Volume")),
  ICSE_ShortPaperStart: ConvertToArray(GetObjectArray("ICSE_ShortPaperStart")),
  SIGMOD_NonResearchPaperStart: ConvertToArray(GetObjectArray("SIGMOD_NonResearchPaperStart")),
  SIGMOD_NonResearchPapersRange: ConvertToArray(GetDictTupleArray("SIGMOD_NonResearchPapersRange")),
  ASE_LongPaperThreshold: GetInt("ASE_LongPaperThreshold"),
  startyear: GetInt("startyear"),
  endyear: GetInt("endyear"),
});
