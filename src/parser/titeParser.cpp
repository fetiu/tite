
// Generated from tite.g4 by ANTLR 4.13.2


#include "titeListener.h"

#include "titeParser.h"


using namespace antlrcpp;

using namespace antlr4;

namespace {

struct TiteParserStaticData final {
  TiteParserStaticData(std::vector<std::string> ruleNames,
                        std::vector<std::string> literalNames,
                        std::vector<std::string> symbolicNames)
      : ruleNames(std::move(ruleNames)), literalNames(std::move(literalNames)),
        symbolicNames(std::move(symbolicNames)),
        vocabulary(this->literalNames, this->symbolicNames) {}

  TiteParserStaticData(const TiteParserStaticData&) = delete;
  TiteParserStaticData(TiteParserStaticData&&) = delete;
  TiteParserStaticData& operator=(const TiteParserStaticData&) = delete;
  TiteParserStaticData& operator=(TiteParserStaticData&&) = delete;

  std::vector<antlr4::dfa::DFA> decisionToDFA;
  antlr4::atn::PredictionContextCache sharedContextCache;
  const std::vector<std::string> ruleNames;
  const std::vector<std::string> literalNames;
  const std::vector<std::string> symbolicNames;
  const antlr4::dfa::Vocabulary vocabulary;
  antlr4::atn::SerializedATNView serializedATN;
  std::unique_ptr<antlr4::atn::ATN> atn;
};

::antlr4::internal::OnceFlag titeParserOnceFlag;
#if ANTLR4_USE_THREAD_LOCAL_CACHE
static thread_local
#endif
std::unique_ptr<TiteParserStaticData> titeParserStaticData = nullptr;

void titeParserInitialize() {
#if ANTLR4_USE_THREAD_LOCAL_CACHE
  if (titeParserStaticData != nullptr) {
    return;
  }
#else
  assert(titeParserStaticData == nullptr);
#endif
  auto staticData = std::make_unique<TiteParserStaticData>(
    std::vector<std::string>{
      "program", "sequence", "declaration", "identifiers", "expression", 
      "type", "function", "condition", "otherwise", "disjunction", "conjunction", 
      "or", "xor", "and", "equality", "relation", "shift", "additive", "product", 
      "factor", "power", "postfix", "primary", "tag", "assign", "literal", 
      "block", "array", "input", "delim"
    },
    std::vector<std::string>{
      "", "':'", "','", "'*'", "'\\u003F'", "'||'", "'&&'", "'|'", "'^'", 
      "'&'", "'=='", "'!='", "'==='", "'!=='", "'<'", "'>'", "'<='", "'>='", 
      "'<<'", "'>>'", "'+'", "'-'", "'..'", "'/'", "'%'", "'//'", "'++'", 
      "'--'", "'~'", "'!'", "'**'", "'.'", "'('", "')'", "'$'", "'#'", "'='", 
      "'=>'", "'<-'", "'@'", "'*='", "'/='", "'%='", "'//='", "'**='", "'+='", 
      "'-='", "'<<='", "'>>='", "'&='", "'^='", "'|='", "'{'", "'}'", "'['", 
      "']'"
    },
    std::vector<std::string>{
      "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
      "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
      "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
      "", "", "", "", "", "IDENTIFIER", "INT", "FLOAT", "CHAR", "STR", "WS", 
      "LF"
    }
  );
  static const int32_t serializedATNSegment[] = {
  	4,1,62,415,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,
  	7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,
  	14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,
  	21,2,22,7,22,2,23,7,23,2,24,7,24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,
  	28,2,29,7,29,1,0,1,0,1,1,1,1,1,1,1,1,5,1,67,8,1,10,1,12,1,70,9,1,1,2,
  	1,2,1,2,3,2,75,8,2,1,2,1,2,1,3,1,3,1,3,5,3,82,8,3,10,3,12,3,85,9,3,1,
  	4,3,4,88,8,4,1,4,1,4,3,4,92,8,4,1,4,1,4,3,4,96,8,4,1,4,1,4,3,4,100,8,
  	4,1,5,1,5,3,5,104,8,5,1,6,1,6,3,6,108,8,6,1,7,1,7,1,7,3,7,113,8,7,1,7,
  	1,7,3,7,117,8,7,1,7,1,7,3,7,121,8,7,3,7,123,8,7,1,8,3,8,126,8,8,1,8,3,
  	8,129,8,8,1,8,1,8,3,8,133,8,8,1,8,1,8,1,9,1,9,1,9,1,9,1,9,3,9,142,8,9,
  	1,9,1,9,3,9,146,8,9,1,9,5,9,149,8,9,10,9,12,9,152,9,9,1,10,1,10,1,10,
  	1,10,1,10,3,10,159,8,10,1,10,1,10,3,10,163,8,10,1,10,5,10,166,8,10,10,
  	10,12,10,169,9,10,1,11,1,11,1,11,1,11,1,11,3,11,176,8,11,1,11,1,11,3,
  	11,180,8,11,1,11,5,11,183,8,11,10,11,12,11,186,9,11,1,12,1,12,1,12,1,
  	12,1,12,3,12,193,8,12,1,12,1,12,3,12,197,8,12,1,12,5,12,200,8,12,10,12,
  	12,12,203,9,12,1,13,1,13,1,13,1,13,1,13,3,13,210,8,13,1,13,1,13,3,13,
  	214,8,13,1,13,5,13,217,8,13,10,13,12,13,220,9,13,1,14,1,14,1,14,1,14,
  	1,14,3,14,227,8,14,1,14,1,14,3,14,231,8,14,1,14,5,14,234,8,14,10,14,12,
  	14,237,9,14,1,15,1,15,1,15,1,15,1,15,3,15,244,8,15,1,15,1,15,3,15,248,
  	8,15,1,15,5,15,251,8,15,10,15,12,15,254,9,15,1,16,1,16,1,16,1,16,1,16,
  	3,16,261,8,16,1,16,1,16,3,16,265,8,16,1,16,5,16,268,8,16,10,16,12,16,
  	271,9,16,1,17,1,17,1,17,1,17,1,17,3,17,278,8,17,1,17,1,17,3,17,282,8,
  	17,1,17,5,17,285,8,17,10,17,12,17,288,9,17,1,18,1,18,1,18,1,18,1,18,3,
  	18,295,8,18,1,18,1,18,3,18,299,8,18,1,18,5,18,302,8,18,10,18,12,18,305,
  	9,18,1,19,1,19,1,19,3,19,310,8,19,1,19,1,19,1,19,1,19,1,19,3,19,317,8,
  	19,1,20,1,20,1,20,3,20,322,8,20,1,20,1,20,3,20,326,8,20,1,20,1,20,3,20,
  	330,8,20,1,21,1,21,1,21,1,21,1,21,3,21,337,8,21,1,21,1,21,1,21,1,21,1,
  	21,1,21,3,21,345,8,21,5,21,347,8,21,10,21,12,21,350,9,21,1,22,3,22,353,
  	8,22,1,22,1,22,1,22,1,22,1,22,1,22,3,22,361,8,22,1,23,1,23,1,24,1,24,
  	1,25,1,25,1,25,1,25,1,25,1,25,3,25,373,8,25,1,26,1,26,3,26,377,8,26,1,
  	26,3,26,380,8,26,1,26,3,26,383,8,26,1,26,1,26,1,27,1,27,3,27,389,8,27,
  	1,27,3,27,392,8,27,1,27,3,27,395,8,27,1,27,1,27,1,28,1,28,3,28,401,8,
  	28,1,28,3,28,404,8,28,1,28,3,28,407,8,28,1,28,1,28,1,29,1,29,3,29,413,
  	8,29,1,29,0,11,18,20,22,24,26,28,30,32,34,36,42,30,0,2,4,6,8,10,12,14,
  	16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,52,54,56,58,0,9,
  	1,0,10,13,1,0,14,17,1,0,18,19,1,0,20,22,2,0,3,3,23,25,2,0,20,21,26,29,
  	2,0,31,31,34,35,1,0,36,51,2,0,2,2,62,62,460,0,60,1,0,0,0,2,62,1,0,0,0,
  	4,74,1,0,0,0,6,78,1,0,0,0,8,87,1,0,0,0,10,103,1,0,0,0,12,105,1,0,0,0,
  	14,122,1,0,0,0,16,125,1,0,0,0,18,136,1,0,0,0,20,153,1,0,0,0,22,170,1,
  	0,0,0,24,187,1,0,0,0,26,204,1,0,0,0,28,221,1,0,0,0,30,238,1,0,0,0,32,
  	255,1,0,0,0,34,272,1,0,0,0,36,289,1,0,0,0,38,316,1,0,0,0,40,329,1,0,0,
  	0,42,331,1,0,0,0,44,360,1,0,0,0,46,362,1,0,0,0,48,364,1,0,0,0,50,372,
  	1,0,0,0,52,374,1,0,0,0,54,386,1,0,0,0,56,398,1,0,0,0,58,410,1,0,0,0,60,
  	61,3,2,1,0,61,1,1,0,0,0,62,68,3,4,2,0,63,64,3,58,29,0,64,65,3,4,2,0,65,
  	67,1,0,0,0,66,63,1,0,0,0,67,70,1,0,0,0,68,66,1,0,0,0,68,69,1,0,0,0,69,
  	3,1,0,0,0,70,68,1,0,0,0,71,72,3,6,3,0,72,73,5,1,0,0,73,75,1,0,0,0,74,
  	71,1,0,0,0,74,75,1,0,0,0,75,76,1,0,0,0,76,77,3,8,4,0,77,5,1,0,0,0,78,
  	83,3,44,22,0,79,80,5,2,0,0,80,82,3,44,22,0,81,79,1,0,0,0,82,85,1,0,0,
  	0,83,81,1,0,0,0,83,84,1,0,0,0,84,7,1,0,0,0,85,83,1,0,0,0,86,88,5,3,0,
  	0,87,86,1,0,0,0,87,88,1,0,0,0,88,99,1,0,0,0,89,100,3,10,5,0,90,92,3,10,
  	5,0,91,90,1,0,0,0,91,92,1,0,0,0,92,93,1,0,0,0,93,95,3,48,24,0,94,96,5,
  	62,0,0,95,94,1,0,0,0,95,96,1,0,0,0,96,97,1,0,0,0,97,98,3,8,4,0,98,100,
  	1,0,0,0,99,89,1,0,0,0,99,91,1,0,0,0,100,9,1,0,0,0,101,104,3,14,7,0,102,
  	104,3,12,6,0,103,101,1,0,0,0,103,102,1,0,0,0,104,11,1,0,0,0,105,107,3,
  	56,28,0,106,108,3,10,5,0,107,106,1,0,0,0,107,108,1,0,0,0,108,13,1,0,0,
  	0,109,123,3,18,9,0,110,112,3,18,9,0,111,113,5,62,0,0,112,111,1,0,0,0,
  	112,113,1,0,0,0,113,114,1,0,0,0,114,116,5,4,0,0,115,117,5,62,0,0,116,
  	115,1,0,0,0,116,117,1,0,0,0,117,120,1,0,0,0,118,121,3,8,4,0,119,121,3,
  	16,8,0,120,118,1,0,0,0,120,119,1,0,0,0,121,123,1,0,0,0,122,109,1,0,0,
  	0,122,110,1,0,0,0,123,15,1,0,0,0,124,126,3,8,4,0,125,124,1,0,0,0,125,
  	126,1,0,0,0,126,128,1,0,0,0,127,129,5,62,0,0,128,127,1,0,0,0,128,129,
  	1,0,0,0,129,130,1,0,0,0,130,132,5,1,0,0,131,133,5,62,0,0,132,131,1,0,
  	0,0,132,133,1,0,0,0,133,134,1,0,0,0,134,135,3,14,7,0,135,17,1,0,0,0,136,
  	137,6,9,-1,0,137,138,3,20,10,0,138,150,1,0,0,0,139,141,10,1,0,0,140,142,
  	5,62,0,0,141,140,1,0,0,0,141,142,1,0,0,0,142,143,1,0,0,0,143,145,5,5,
  	0,0,144,146,5,62,0,0,145,144,1,0,0,0,145,146,1,0,0,0,146,147,1,0,0,0,
  	147,149,3,20,10,0,148,139,1,0,0,0,149,152,1,0,0,0,150,148,1,0,0,0,150,
  	151,1,0,0,0,151,19,1,0,0,0,152,150,1,0,0,0,153,154,6,10,-1,0,154,155,
  	3,22,11,0,155,167,1,0,0,0,156,158,10,1,0,0,157,159,5,62,0,0,158,157,1,
  	0,0,0,158,159,1,0,0,0,159,160,1,0,0,0,160,162,5,6,0,0,161,163,5,62,0,
  	0,162,161,1,0,0,0,162,163,1,0,0,0,163,164,1,0,0,0,164,166,3,22,11,0,165,
  	156,1,0,0,0,166,169,1,0,0,0,167,165,1,0,0,0,167,168,1,0,0,0,168,21,1,
  	0,0,0,169,167,1,0,0,0,170,171,6,11,-1,0,171,172,3,24,12,0,172,184,1,0,
  	0,0,173,175,10,1,0,0,174,176,5,62,0,0,175,174,1,0,0,0,175,176,1,0,0,0,
  	176,177,1,0,0,0,177,179,5,7,0,0,178,180,5,62,0,0,179,178,1,0,0,0,179,
  	180,1,0,0,0,180,181,1,0,0,0,181,183,3,24,12,0,182,173,1,0,0,0,183,186,
  	1,0,0,0,184,182,1,0,0,0,184,185,1,0,0,0,185,23,1,0,0,0,186,184,1,0,0,
  	0,187,188,6,12,-1,0,188,189,3,26,13,0,189,201,1,0,0,0,190,192,10,1,0,
  	0,191,193,5,62,0,0,192,191,1,0,0,0,192,193,1,0,0,0,193,194,1,0,0,0,194,
  	196,5,8,0,0,195,197,5,62,0,0,196,195,1,0,0,0,196,197,1,0,0,0,197,198,
  	1,0,0,0,198,200,3,26,13,0,199,190,1,0,0,0,200,203,1,0,0,0,201,199,1,0,
  	0,0,201,202,1,0,0,0,202,25,1,0,0,0,203,201,1,0,0,0,204,205,6,13,-1,0,
  	205,206,3,28,14,0,206,218,1,0,0,0,207,209,10,1,0,0,208,210,5,62,0,0,209,
  	208,1,0,0,0,209,210,1,0,0,0,210,211,1,0,0,0,211,213,5,9,0,0,212,214,5,
  	62,0,0,213,212,1,0,0,0,213,214,1,0,0,0,214,215,1,0,0,0,215,217,3,28,14,
  	0,216,207,1,0,0,0,217,220,1,0,0,0,218,216,1,0,0,0,218,219,1,0,0,0,219,
  	27,1,0,0,0,220,218,1,0,0,0,221,222,6,14,-1,0,222,223,3,30,15,0,223,235,
  	1,0,0,0,224,226,10,1,0,0,225,227,5,62,0,0,226,225,1,0,0,0,226,227,1,0,
  	0,0,227,228,1,0,0,0,228,230,7,0,0,0,229,231,5,62,0,0,230,229,1,0,0,0,
  	230,231,1,0,0,0,231,232,1,0,0,0,232,234,3,30,15,0,233,224,1,0,0,0,234,
  	237,1,0,0,0,235,233,1,0,0,0,235,236,1,0,0,0,236,29,1,0,0,0,237,235,1,
  	0,0,0,238,239,6,15,-1,0,239,240,3,32,16,0,240,252,1,0,0,0,241,243,10,
  	1,0,0,242,244,5,62,0,0,243,242,1,0,0,0,243,244,1,0,0,0,244,245,1,0,0,
  	0,245,247,7,1,0,0,246,248,5,62,0,0,247,246,1,0,0,0,247,248,1,0,0,0,248,
  	249,1,0,0,0,249,251,3,32,16,0,250,241,1,0,0,0,251,254,1,0,0,0,252,250,
  	1,0,0,0,252,253,1,0,0,0,253,31,1,0,0,0,254,252,1,0,0,0,255,256,6,16,-1,
  	0,256,257,3,34,17,0,257,269,1,0,0,0,258,260,10,1,0,0,259,261,5,62,0,0,
  	260,259,1,0,0,0,260,261,1,0,0,0,261,262,1,0,0,0,262,264,7,2,0,0,263,265,
  	5,62,0,0,264,263,1,0,0,0,264,265,1,0,0,0,265,266,1,0,0,0,266,268,3,34,
  	17,0,267,258,1,0,0,0,268,271,1,0,0,0,269,267,1,0,0,0,269,270,1,0,0,0,
  	270,33,1,0,0,0,271,269,1,0,0,0,272,273,6,17,-1,0,273,274,3,36,18,0,274,
  	286,1,0,0,0,275,277,10,1,0,0,276,278,5,62,0,0,277,276,1,0,0,0,277,278,
  	1,0,0,0,278,279,1,0,0,0,279,281,7,3,0,0,280,282,5,62,0,0,281,280,1,0,
  	0,0,281,282,1,0,0,0,282,283,1,0,0,0,283,285,3,36,18,0,284,275,1,0,0,0,
  	285,288,1,0,0,0,286,284,1,0,0,0,286,287,1,0,0,0,287,35,1,0,0,0,288,286,
  	1,0,0,0,289,290,6,18,-1,0,290,291,3,38,19,0,291,303,1,0,0,0,292,294,10,
  	1,0,0,293,295,5,62,0,0,294,293,1,0,0,0,294,295,1,0,0,0,295,296,1,0,0,
  	0,296,298,7,4,0,0,297,299,5,62,0,0,298,297,1,0,0,0,298,299,1,0,0,0,299,
  	300,1,0,0,0,300,302,3,38,19,0,301,292,1,0,0,0,302,305,1,0,0,0,303,301,
  	1,0,0,0,303,304,1,0,0,0,304,37,1,0,0,0,305,303,1,0,0,0,306,317,3,40,20,
  	0,307,309,7,5,0,0,308,310,5,62,0,0,309,308,1,0,0,0,309,310,1,0,0,0,310,
  	311,1,0,0,0,311,317,3,38,19,0,312,313,5,7,0,0,313,314,3,38,19,0,314,315,
  	5,7,0,0,315,317,1,0,0,0,316,306,1,0,0,0,316,307,1,0,0,0,316,312,1,0,0,
  	0,317,39,1,0,0,0,318,330,3,42,21,0,319,321,3,42,21,0,320,322,5,62,0,0,
  	321,320,1,0,0,0,321,322,1,0,0,0,322,323,1,0,0,0,323,325,5,30,0,0,324,
  	326,5,62,0,0,325,324,1,0,0,0,325,326,1,0,0,0,326,327,1,0,0,0,327,328,
  	3,38,19,0,328,330,1,0,0,0,329,318,1,0,0,0,329,319,1,0,0,0,330,41,1,0,
  	0,0,331,332,6,21,-1,0,332,333,3,44,22,0,333,348,1,0,0,0,334,336,10,1,
  	0,0,335,337,5,62,0,0,336,335,1,0,0,0,336,337,1,0,0,0,337,344,1,0,0,0,
  	338,345,5,26,0,0,339,345,5,27,0,0,340,341,5,31,0,0,341,345,5,56,0,0,342,
  	345,3,54,27,0,343,345,3,56,28,0,344,338,1,0,0,0,344,339,1,0,0,0,344,340,
  	1,0,0,0,344,342,1,0,0,0,344,343,1,0,0,0,345,347,1,0,0,0,346,334,1,0,0,
  	0,347,350,1,0,0,0,348,346,1,0,0,0,348,349,1,0,0,0,349,43,1,0,0,0,350,
  	348,1,0,0,0,351,353,3,46,23,0,352,351,1,0,0,0,352,353,1,0,0,0,353,354,
  	1,0,0,0,354,361,5,56,0,0,355,361,3,50,25,0,356,357,5,32,0,0,357,358,3,
  	8,4,0,358,359,5,33,0,0,359,361,1,0,0,0,360,352,1,0,0,0,360,355,1,0,0,
  	0,360,356,1,0,0,0,361,45,1,0,0,0,362,363,7,6,0,0,363,47,1,0,0,0,364,365,
  	7,7,0,0,365,49,1,0,0,0,366,373,5,57,0,0,367,373,5,58,0,0,368,373,5,59,
  	0,0,369,373,5,60,0,0,370,373,3,52,26,0,371,373,3,54,27,0,372,366,1,0,
  	0,0,372,367,1,0,0,0,372,368,1,0,0,0,372,369,1,0,0,0,372,370,1,0,0,0,372,
  	371,1,0,0,0,373,51,1,0,0,0,374,376,5,52,0,0,375,377,5,62,0,0,376,375,
  	1,0,0,0,376,377,1,0,0,0,377,379,1,0,0,0,378,380,3,2,1,0,379,378,1,0,0,
  	0,379,380,1,0,0,0,380,382,1,0,0,0,381,383,5,62,0,0,382,381,1,0,0,0,382,
  	383,1,0,0,0,383,384,1,0,0,0,384,385,5,53,0,0,385,53,1,0,0,0,386,388,5,
  	54,0,0,387,389,5,62,0,0,388,387,1,0,0,0,388,389,1,0,0,0,389,391,1,0,0,
  	0,390,392,3,2,1,0,391,390,1,0,0,0,391,392,1,0,0,0,392,394,1,0,0,0,393,
  	395,5,62,0,0,394,393,1,0,0,0,394,395,1,0,0,0,395,396,1,0,0,0,396,397,
  	5,55,0,0,397,55,1,0,0,0,398,400,5,32,0,0,399,401,5,62,0,0,400,399,1,0,
  	0,0,400,401,1,0,0,0,401,403,1,0,0,0,402,404,3,2,1,0,403,402,1,0,0,0,403,
  	404,1,0,0,0,404,406,1,0,0,0,405,407,5,62,0,0,406,405,1,0,0,0,406,407,
  	1,0,0,0,407,408,1,0,0,0,408,409,5,33,0,0,409,57,1,0,0,0,410,412,7,8,0,
  	0,411,413,5,62,0,0,412,411,1,0,0,0,412,413,1,0,0,0,413,59,1,0,0,0,67,
  	68,74,83,87,91,95,99,103,107,112,116,120,122,125,128,132,141,145,150,
  	158,162,167,175,179,184,192,196,201,209,213,218,226,230,235,243,247,252,
  	260,264,269,277,281,286,294,298,303,309,316,321,325,329,336,344,348,352,
  	360,372,376,379,382,388,391,394,400,403,406,412
  };
  staticData->serializedATN = antlr4::atn::SerializedATNView(serializedATNSegment, sizeof(serializedATNSegment) / sizeof(serializedATNSegment[0]));

  antlr4::atn::ATNDeserializer deserializer;
  staticData->atn = deserializer.deserialize(staticData->serializedATN);

  const size_t count = staticData->atn->getNumberOfDecisions();
  staticData->decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    staticData->decisionToDFA.emplace_back(staticData->atn->getDecisionState(i), i);
  }
  titeParserStaticData = std::move(staticData);
}

}

titeParser::titeParser(TokenStream *input) : titeParser(input, antlr4::atn::ParserATNSimulatorOptions()) {}

titeParser::titeParser(TokenStream *input, const antlr4::atn::ParserATNSimulatorOptions &options) : Parser(input) {
  titeParser::initialize();
  _interpreter = new atn::ParserATNSimulator(this, *titeParserStaticData->atn, titeParserStaticData->decisionToDFA, titeParserStaticData->sharedContextCache, options);
}

titeParser::~titeParser() {
  delete _interpreter;
}

const atn::ATN& titeParser::getATN() const {
  return *titeParserStaticData->atn;
}

std::string titeParser::getGrammarFileName() const {
  return "tite.g4";
}

const std::vector<std::string>& titeParser::getRuleNames() const {
  return titeParserStaticData->ruleNames;
}

const dfa::Vocabulary& titeParser::getVocabulary() const {
  return titeParserStaticData->vocabulary;
}

antlr4::atn::SerializedATNView titeParser::getSerializedATN() const {
  return titeParserStaticData->serializedATN;
}


//----------------- ProgramContext ------------------------------------------------------------------

titeParser::ProgramContext::ProgramContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::SequenceContext* titeParser::ProgramContext::sequence() {
  return getRuleContext<titeParser::SequenceContext>(0);
}


size_t titeParser::ProgramContext::getRuleIndex() const {
  return titeParser::RuleProgram;
}

void titeParser::ProgramContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProgram(this);
}

void titeParser::ProgramContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProgram(this);
}

titeParser::ProgramContext* titeParser::program() {
  ProgramContext *_localctx = _tracker.createInstance<ProgramContext>(_ctx, getState());
  enterRule(_localctx, 0, titeParser::RuleProgram);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(60);
    sequence();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SequenceContext ------------------------------------------------------------------

titeParser::SequenceContext::SequenceContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<titeParser::DeclarationContext *> titeParser::SequenceContext::declaration() {
  return getRuleContexts<titeParser::DeclarationContext>();
}

titeParser::DeclarationContext* titeParser::SequenceContext::declaration(size_t i) {
  return getRuleContext<titeParser::DeclarationContext>(i);
}

std::vector<titeParser::DelimContext *> titeParser::SequenceContext::delim() {
  return getRuleContexts<titeParser::DelimContext>();
}

titeParser::DelimContext* titeParser::SequenceContext::delim(size_t i) {
  return getRuleContext<titeParser::DelimContext>(i);
}


size_t titeParser::SequenceContext::getRuleIndex() const {
  return titeParser::RuleSequence;
}

void titeParser::SequenceContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterSequence(this);
}

void titeParser::SequenceContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitSequence(this);
}

titeParser::SequenceContext* titeParser::sequence() {
  SequenceContext *_localctx = _tracker.createInstance<SequenceContext>(_ctx, getState());
  enterRule(_localctx, 2, titeParser::RuleSequence);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(62);
    declaration();
    setState(68);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 0, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        setState(63);
        delim();
        setState(64);
        declaration(); 
      }
      setState(70);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 0, _ctx);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DeclarationContext ------------------------------------------------------------------

titeParser::DeclarationContext::DeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::ExpressionContext* titeParser::DeclarationContext::expression() {
  return getRuleContext<titeParser::ExpressionContext>(0);
}

titeParser::IdentifiersContext* titeParser::DeclarationContext::identifiers() {
  return getRuleContext<titeParser::IdentifiersContext>(0);
}


size_t titeParser::DeclarationContext::getRuleIndex() const {
  return titeParser::RuleDeclaration;
}

void titeParser::DeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterDeclaration(this);
}

void titeParser::DeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitDeclaration(this);
}

titeParser::DeclarationContext* titeParser::declaration() {
  DeclarationContext *_localctx = _tracker.createInstance<DeclarationContext>(_ctx, getState());
  enterRule(_localctx, 4, titeParser::RuleDeclaration);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(74);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 1, _ctx)) {
    case 1: {
      setState(71);
      identifiers();
      setState(72);
      match(titeParser::T__0);
      break;
    }

    default:
      break;
    }
    setState(76);
    expression();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- IdentifiersContext ------------------------------------------------------------------

titeParser::IdentifiersContext::IdentifiersContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<titeParser::PrimaryContext *> titeParser::IdentifiersContext::primary() {
  return getRuleContexts<titeParser::PrimaryContext>();
}

titeParser::PrimaryContext* titeParser::IdentifiersContext::primary(size_t i) {
  return getRuleContext<titeParser::PrimaryContext>(i);
}


size_t titeParser::IdentifiersContext::getRuleIndex() const {
  return titeParser::RuleIdentifiers;
}

void titeParser::IdentifiersContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterIdentifiers(this);
}

void titeParser::IdentifiersContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitIdentifiers(this);
}

titeParser::IdentifiersContext* titeParser::identifiers() {
  IdentifiersContext *_localctx = _tracker.createInstance<IdentifiersContext>(_ctx, getState());
  enterRule(_localctx, 6, titeParser::RuleIdentifiers);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(78);
    primary();
    setState(83);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while (_la == titeParser::T__1) {
      setState(79);
      match(titeParser::T__1);
      setState(80);
      primary();
      setState(85);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExpressionContext ------------------------------------------------------------------

titeParser::ExpressionContext::ExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::TypeContext* titeParser::ExpressionContext::type() {
  return getRuleContext<titeParser::TypeContext>(0);
}

titeParser::AssignContext* titeParser::ExpressionContext::assign() {
  return getRuleContext<titeParser::AssignContext>(0);
}

titeParser::ExpressionContext* titeParser::ExpressionContext::expression() {
  return getRuleContext<titeParser::ExpressionContext>(0);
}

tree::TerminalNode* titeParser::ExpressionContext::LF() {
  return getToken(titeParser::LF, 0);
}


size_t titeParser::ExpressionContext::getRuleIndex() const {
  return titeParser::RuleExpression;
}

void titeParser::ExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExpression(this);
}

void titeParser::ExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExpression(this);
}

titeParser::ExpressionContext* titeParser::expression() {
  ExpressionContext *_localctx = _tracker.createInstance<ExpressionContext>(_ctx, getState());
  enterRule(_localctx, 8, titeParser::RuleExpression);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(87);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::T__2) {
      setState(86);
      match(titeParser::T__2);
    }
    setState(99);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 6, _ctx)) {
    case 1: {
      setState(89);
      type();
      break;
    }

    case 2: {
      setState(91);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if ((((_la & ~ 0x3fULL) == 0) &&
        ((1ULL << _la) & 2256303472304455808) != 0)) {
        setState(90);
        type();
      }
      setState(93);
      assign();
      setState(95);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == titeParser::LF) {
        setState(94);
        match(titeParser::LF);
      }
      setState(97);
      expression();
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TypeContext ------------------------------------------------------------------

titeParser::TypeContext::TypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::ConditionContext* titeParser::TypeContext::condition() {
  return getRuleContext<titeParser::ConditionContext>(0);
}

titeParser::FunctionContext* titeParser::TypeContext::function() {
  return getRuleContext<titeParser::FunctionContext>(0);
}


size_t titeParser::TypeContext::getRuleIndex() const {
  return titeParser::RuleType;
}

void titeParser::TypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterType(this);
}

void titeParser::TypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitType(this);
}

titeParser::TypeContext* titeParser::type() {
  TypeContext *_localctx = _tracker.createInstance<TypeContext>(_ctx, getState());
  enterRule(_localctx, 10, titeParser::RuleType);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(103);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 7, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(101);
      condition();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(102);
      function();
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FunctionContext ------------------------------------------------------------------

titeParser::FunctionContext::FunctionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::InputContext* titeParser::FunctionContext::input() {
  return getRuleContext<titeParser::InputContext>(0);
}

titeParser::TypeContext* titeParser::FunctionContext::type() {
  return getRuleContext<titeParser::TypeContext>(0);
}


size_t titeParser::FunctionContext::getRuleIndex() const {
  return titeParser::RuleFunction;
}

void titeParser::FunctionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFunction(this);
}

void titeParser::FunctionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFunction(this);
}

titeParser::FunctionContext* titeParser::function() {
  FunctionContext *_localctx = _tracker.createInstance<FunctionContext>(_ctx, getState());
  enterRule(_localctx, 12, titeParser::RuleFunction);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(105);
    input();
    setState(107);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 2256303472304455808) != 0)) {
      setState(106);
      type();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ConditionContext ------------------------------------------------------------------

titeParser::ConditionContext::ConditionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::DisjunctionContext* titeParser::ConditionContext::disjunction() {
  return getRuleContext<titeParser::DisjunctionContext>(0);
}

titeParser::ExpressionContext* titeParser::ConditionContext::expression() {
  return getRuleContext<titeParser::ExpressionContext>(0);
}

titeParser::OtherwiseContext* titeParser::ConditionContext::otherwise() {
  return getRuleContext<titeParser::OtherwiseContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::ConditionContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::ConditionContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::ConditionContext::getRuleIndex() const {
  return titeParser::RuleCondition;
}

void titeParser::ConditionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterCondition(this);
}

void titeParser::ConditionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitCondition(this);
}

titeParser::ConditionContext* titeParser::condition() {
  ConditionContext *_localctx = _tracker.createInstance<ConditionContext>(_ctx, getState());
  enterRule(_localctx, 14, titeParser::RuleCondition);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(122);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 12, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(109);
      disjunction(0);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(110);
      disjunction(0);
      setState(112);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == titeParser::LF) {
        setState(111);
        match(titeParser::LF);
      }
      setState(114);
      match(titeParser::T__3);
      setState(116);
      _errHandler->sync(this);

      switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 10, _ctx)) {
      case 1: {
        setState(115);
        match(titeParser::LF);
        break;
      }

      default:
        break;
      }
      setState(120);
      _errHandler->sync(this);
      switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx)) {
      case 1: {
        setState(118);
        expression();
        break;
      }

      case 2: {
        setState(119);
        otherwise();
        break;
      }

      default:
        break;
      }
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- OtherwiseContext ------------------------------------------------------------------

titeParser::OtherwiseContext::OtherwiseContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::ConditionContext* titeParser::OtherwiseContext::condition() {
  return getRuleContext<titeParser::ConditionContext>(0);
}

titeParser::ExpressionContext* titeParser::OtherwiseContext::expression() {
  return getRuleContext<titeParser::ExpressionContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::OtherwiseContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::OtherwiseContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::OtherwiseContext::getRuleIndex() const {
  return titeParser::RuleOtherwise;
}

void titeParser::OtherwiseContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterOtherwise(this);
}

void titeParser::OtherwiseContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitOtherwise(this);
}

titeParser::OtherwiseContext* titeParser::otherwise() {
  OtherwiseContext *_localctx = _tracker.createInstance<OtherwiseContext>(_ctx, getState());
  enterRule(_localctx, 16, titeParser::RuleOtherwise);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(125);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 2260807003212349576) != 0)) {
      setState(124);
      expression();
    }
    setState(128);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::LF) {
      setState(127);
      match(titeParser::LF);
    }
    setState(130);
    match(titeParser::T__0);
    setState(132);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::LF) {
      setState(131);
      match(titeParser::LF);
    }
    setState(134);
    condition();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DisjunctionContext ------------------------------------------------------------------

titeParser::DisjunctionContext::DisjunctionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::ConjunctionContext* titeParser::DisjunctionContext::conjunction() {
  return getRuleContext<titeParser::ConjunctionContext>(0);
}

titeParser::DisjunctionContext* titeParser::DisjunctionContext::disjunction() {
  return getRuleContext<titeParser::DisjunctionContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::DisjunctionContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::DisjunctionContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::DisjunctionContext::getRuleIndex() const {
  return titeParser::RuleDisjunction;
}

void titeParser::DisjunctionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterDisjunction(this);
}

void titeParser::DisjunctionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitDisjunction(this);
}


titeParser::DisjunctionContext* titeParser::disjunction() {
   return disjunction(0);
}

titeParser::DisjunctionContext* titeParser::disjunction(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::DisjunctionContext *_localctx = _tracker.createInstance<DisjunctionContext>(_ctx, parentState);
  titeParser::DisjunctionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 18;
  enterRecursionRule(_localctx, 18, titeParser::RuleDisjunction, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(137);
    conjunction(0);
    _ctx->stop = _input->LT(-1);
    setState(150);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 18, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<DisjunctionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleDisjunction);
        setState(139);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(141);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(140);
          match(titeParser::LF);
        }
        setState(143);
        match(titeParser::T__4);
        setState(145);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(144);
          match(titeParser::LF);
        }
        setState(147);
        conjunction(0); 
      }
      setState(152);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 18, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ConjunctionContext ------------------------------------------------------------------

titeParser::ConjunctionContext::ConjunctionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::OrContext* titeParser::ConjunctionContext::or_() {
  return getRuleContext<titeParser::OrContext>(0);
}

titeParser::ConjunctionContext* titeParser::ConjunctionContext::conjunction() {
  return getRuleContext<titeParser::ConjunctionContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::ConjunctionContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::ConjunctionContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::ConjunctionContext::getRuleIndex() const {
  return titeParser::RuleConjunction;
}

void titeParser::ConjunctionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterConjunction(this);
}

void titeParser::ConjunctionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitConjunction(this);
}


titeParser::ConjunctionContext* titeParser::conjunction() {
   return conjunction(0);
}

titeParser::ConjunctionContext* titeParser::conjunction(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::ConjunctionContext *_localctx = _tracker.createInstance<ConjunctionContext>(_ctx, parentState);
  titeParser::ConjunctionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 20;
  enterRecursionRule(_localctx, 20, titeParser::RuleConjunction, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(154);
    or_(0);
    _ctx->stop = _input->LT(-1);
    setState(167);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 21, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ConjunctionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleConjunction);
        setState(156);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(158);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(157);
          match(titeParser::LF);
        }
        setState(160);
        match(titeParser::T__5);
        setState(162);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(161);
          match(titeParser::LF);
        }
        setState(164);
        or_(0); 
      }
      setState(169);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 21, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- OrContext ------------------------------------------------------------------

titeParser::OrContext::OrContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::XorContext* titeParser::OrContext::xor_() {
  return getRuleContext<titeParser::XorContext>(0);
}

titeParser::OrContext* titeParser::OrContext::or_() {
  return getRuleContext<titeParser::OrContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::OrContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::OrContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::OrContext::getRuleIndex() const {
  return titeParser::RuleOr;
}

void titeParser::OrContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterOr(this);
}

void titeParser::OrContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitOr(this);
}


titeParser::OrContext* titeParser::or_() {
   return or_(0);
}

titeParser::OrContext* titeParser::or_(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::OrContext *_localctx = _tracker.createInstance<OrContext>(_ctx, parentState);
  titeParser::OrContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 22;
  enterRecursionRule(_localctx, 22, titeParser::RuleOr, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(171);
    xor_(0);
    _ctx->stop = _input->LT(-1);
    setState(184);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 24, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<OrContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleOr);
        setState(173);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(175);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(174);
          match(titeParser::LF);
        }
        setState(177);
        match(titeParser::T__6);
        setState(179);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(178);
          match(titeParser::LF);
        }
        setState(181);
        xor_(0); 
      }
      setState(186);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 24, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- XorContext ------------------------------------------------------------------

titeParser::XorContext::XorContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::AndContext* titeParser::XorContext::and_() {
  return getRuleContext<titeParser::AndContext>(0);
}

titeParser::XorContext* titeParser::XorContext::xor_() {
  return getRuleContext<titeParser::XorContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::XorContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::XorContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::XorContext::getRuleIndex() const {
  return titeParser::RuleXor;
}

void titeParser::XorContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterXor(this);
}

void titeParser::XorContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitXor(this);
}


titeParser::XorContext* titeParser::xor_() {
   return xor_(0);
}

titeParser::XorContext* titeParser::xor_(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::XorContext *_localctx = _tracker.createInstance<XorContext>(_ctx, parentState);
  titeParser::XorContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 24;
  enterRecursionRule(_localctx, 24, titeParser::RuleXor, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(188);
    and_(0);
    _ctx->stop = _input->LT(-1);
    setState(201);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 27, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<XorContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleXor);
        setState(190);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(192);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(191);
          match(titeParser::LF);
        }
        setState(194);
        match(titeParser::T__7);
        setState(196);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(195);
          match(titeParser::LF);
        }
        setState(198);
        and_(0); 
      }
      setState(203);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 27, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- AndContext ------------------------------------------------------------------

titeParser::AndContext::AndContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::EqualityContext* titeParser::AndContext::equality() {
  return getRuleContext<titeParser::EqualityContext>(0);
}

titeParser::AndContext* titeParser::AndContext::and_() {
  return getRuleContext<titeParser::AndContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::AndContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::AndContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::AndContext::getRuleIndex() const {
  return titeParser::RuleAnd;
}

void titeParser::AndContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAnd(this);
}

void titeParser::AndContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAnd(this);
}


titeParser::AndContext* titeParser::and_() {
   return and_(0);
}

titeParser::AndContext* titeParser::and_(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::AndContext *_localctx = _tracker.createInstance<AndContext>(_ctx, parentState);
  titeParser::AndContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 26;
  enterRecursionRule(_localctx, 26, titeParser::RuleAnd, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(205);
    equality(0);
    _ctx->stop = _input->LT(-1);
    setState(218);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 30, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<AndContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleAnd);
        setState(207);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(209);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(208);
          match(titeParser::LF);
        }
        setState(211);
        match(titeParser::T__8);
        setState(213);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(212);
          match(titeParser::LF);
        }
        setState(215);
        equality(0); 
      }
      setState(220);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 30, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- EqualityContext ------------------------------------------------------------------

titeParser::EqualityContext::EqualityContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::RelationContext* titeParser::EqualityContext::relation() {
  return getRuleContext<titeParser::RelationContext>(0);
}

titeParser::EqualityContext* titeParser::EqualityContext::equality() {
  return getRuleContext<titeParser::EqualityContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::EqualityContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::EqualityContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::EqualityContext::getRuleIndex() const {
  return titeParser::RuleEquality;
}

void titeParser::EqualityContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterEquality(this);
}

void titeParser::EqualityContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitEquality(this);
}


titeParser::EqualityContext* titeParser::equality() {
   return equality(0);
}

titeParser::EqualityContext* titeParser::equality(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::EqualityContext *_localctx = _tracker.createInstance<EqualityContext>(_ctx, parentState);
  titeParser::EqualityContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 28;
  enterRecursionRule(_localctx, 28, titeParser::RuleEquality, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(222);
    relation(0);
    _ctx->stop = _input->LT(-1);
    setState(235);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 33, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<EqualityContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleEquality);
        setState(224);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(226);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(225);
          match(titeParser::LF);
        }
        setState(228);
        _la = _input->LA(1);
        if (!((((_la & ~ 0x3fULL) == 0) &&
          ((1ULL << _la) & 15360) != 0))) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(230);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(229);
          match(titeParser::LF);
        }
        setState(232);
        relation(0); 
      }
      setState(237);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 33, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- RelationContext ------------------------------------------------------------------

titeParser::RelationContext::RelationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::ShiftContext* titeParser::RelationContext::shift() {
  return getRuleContext<titeParser::ShiftContext>(0);
}

titeParser::RelationContext* titeParser::RelationContext::relation() {
  return getRuleContext<titeParser::RelationContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::RelationContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::RelationContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::RelationContext::getRuleIndex() const {
  return titeParser::RuleRelation;
}

void titeParser::RelationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRelation(this);
}

void titeParser::RelationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRelation(this);
}


titeParser::RelationContext* titeParser::relation() {
   return relation(0);
}

titeParser::RelationContext* titeParser::relation(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::RelationContext *_localctx = _tracker.createInstance<RelationContext>(_ctx, parentState);
  titeParser::RelationContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 30;
  enterRecursionRule(_localctx, 30, titeParser::RuleRelation, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(239);
    shift(0);
    _ctx->stop = _input->LT(-1);
    setState(252);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 36, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<RelationContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleRelation);
        setState(241);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(243);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(242);
          match(titeParser::LF);
        }
        setState(245);
        _la = _input->LA(1);
        if (!((((_la & ~ 0x3fULL) == 0) &&
          ((1ULL << _la) & 245760) != 0))) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(247);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(246);
          match(titeParser::LF);
        }
        setState(249);
        shift(0); 
      }
      setState(254);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 36, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ShiftContext ------------------------------------------------------------------

titeParser::ShiftContext::ShiftContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::AdditiveContext* titeParser::ShiftContext::additive() {
  return getRuleContext<titeParser::AdditiveContext>(0);
}

titeParser::ShiftContext* titeParser::ShiftContext::shift() {
  return getRuleContext<titeParser::ShiftContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::ShiftContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::ShiftContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::ShiftContext::getRuleIndex() const {
  return titeParser::RuleShift;
}

void titeParser::ShiftContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterShift(this);
}

void titeParser::ShiftContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitShift(this);
}


titeParser::ShiftContext* titeParser::shift() {
   return shift(0);
}

titeParser::ShiftContext* titeParser::shift(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::ShiftContext *_localctx = _tracker.createInstance<ShiftContext>(_ctx, parentState);
  titeParser::ShiftContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 32;
  enterRecursionRule(_localctx, 32, titeParser::RuleShift, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(256);
    additive(0);
    _ctx->stop = _input->LT(-1);
    setState(269);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 39, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ShiftContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleShift);
        setState(258);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(260);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(259);
          match(titeParser::LF);
        }
        setState(262);
        _la = _input->LA(1);
        if (!(_la == titeParser::T__17

        || _la == titeParser::T__18)) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(264);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(263);
          match(titeParser::LF);
        }
        setState(266);
        additive(0); 
      }
      setState(271);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 39, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- AdditiveContext ------------------------------------------------------------------

titeParser::AdditiveContext::AdditiveContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::ProductContext* titeParser::AdditiveContext::product() {
  return getRuleContext<titeParser::ProductContext>(0);
}

titeParser::AdditiveContext* titeParser::AdditiveContext::additive() {
  return getRuleContext<titeParser::AdditiveContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::AdditiveContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::AdditiveContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::AdditiveContext::getRuleIndex() const {
  return titeParser::RuleAdditive;
}

void titeParser::AdditiveContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAdditive(this);
}

void titeParser::AdditiveContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAdditive(this);
}


titeParser::AdditiveContext* titeParser::additive() {
   return additive(0);
}

titeParser::AdditiveContext* titeParser::additive(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::AdditiveContext *_localctx = _tracker.createInstance<AdditiveContext>(_ctx, parentState);
  titeParser::AdditiveContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 34;
  enterRecursionRule(_localctx, 34, titeParser::RuleAdditive, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(273);
    product(0);
    _ctx->stop = _input->LT(-1);
    setState(286);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 42, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<AdditiveContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleAdditive);
        setState(275);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(277);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(276);
          match(titeParser::LF);
        }
        setState(279);
        _la = _input->LA(1);
        if (!((((_la & ~ 0x3fULL) == 0) &&
          ((1ULL << _la) & 7340032) != 0))) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(281);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(280);
          match(titeParser::LF);
        }
        setState(283);
        product(0); 
      }
      setState(288);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 42, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ProductContext ------------------------------------------------------------------

titeParser::ProductContext::ProductContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::FactorContext* titeParser::ProductContext::factor() {
  return getRuleContext<titeParser::FactorContext>(0);
}

titeParser::ProductContext* titeParser::ProductContext::product() {
  return getRuleContext<titeParser::ProductContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::ProductContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::ProductContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::ProductContext::getRuleIndex() const {
  return titeParser::RuleProduct;
}

void titeParser::ProductContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProduct(this);
}

void titeParser::ProductContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProduct(this);
}


titeParser::ProductContext* titeParser::product() {
   return product(0);
}

titeParser::ProductContext* titeParser::product(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::ProductContext *_localctx = _tracker.createInstance<ProductContext>(_ctx, parentState);
  titeParser::ProductContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 36;
  enterRecursionRule(_localctx, 36, titeParser::RuleProduct, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(290);
    factor();
    _ctx->stop = _input->LT(-1);
    setState(303);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 45, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<ProductContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleProduct);
        setState(292);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(294);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(293);
          match(titeParser::LF);
        }
        setState(296);
        _la = _input->LA(1);
        if (!((((_la & ~ 0x3fULL) == 0) &&
          ((1ULL << _la) & 58720264) != 0))) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(298);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(297);
          match(titeParser::LF);
        }
        setState(300);
        factor(); 
      }
      setState(305);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 45, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- FactorContext ------------------------------------------------------------------

titeParser::FactorContext::FactorContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::PowerContext* titeParser::FactorContext::power() {
  return getRuleContext<titeParser::PowerContext>(0);
}

titeParser::FactorContext* titeParser::FactorContext::factor() {
  return getRuleContext<titeParser::FactorContext>(0);
}

tree::TerminalNode* titeParser::FactorContext::LF() {
  return getToken(titeParser::LF, 0);
}


size_t titeParser::FactorContext::getRuleIndex() const {
  return titeParser::RuleFactor;
}

void titeParser::FactorContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFactor(this);
}

void titeParser::FactorContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFactor(this);
}

titeParser::FactorContext* titeParser::factor() {
  FactorContext *_localctx = _tracker.createInstance<FactorContext>(_ctx, getState());
  enterRule(_localctx, 38, titeParser::RuleFactor);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(316);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case titeParser::T__30:
      case titeParser::T__31:
      case titeParser::T__33:
      case titeParser::T__34:
      case titeParser::T__51:
      case titeParser::T__53:
      case titeParser::IDENTIFIER:
      case titeParser::INT:
      case titeParser::FLOAT:
      case titeParser::CHAR:
      case titeParser::STR: {
        enterOuterAlt(_localctx, 1);
        setState(306);
        power();
        break;
      }

      case titeParser::T__19:
      case titeParser::T__20:
      case titeParser::T__25:
      case titeParser::T__26:
      case titeParser::T__27:
      case titeParser::T__28: {
        enterOuterAlt(_localctx, 2);
        setState(307);
        _la = _input->LA(1);
        if (!((((_la & ~ 0x3fULL) == 0) &&
          ((1ULL << _la) & 1009778688) != 0))) {
        _errHandler->recoverInline(this);
        }
        else {
          _errHandler->reportMatch(this);
          consume();
        }
        setState(309);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(308);
          match(titeParser::LF);
        }
        setState(311);
        factor();
        break;
      }

      case titeParser::T__6: {
        enterOuterAlt(_localctx, 3);
        setState(312);
        match(titeParser::T__6);
        setState(313);
        factor();
        setState(314);
        match(titeParser::T__6);
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PowerContext ------------------------------------------------------------------

titeParser::PowerContext::PowerContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::PostfixContext* titeParser::PowerContext::postfix() {
  return getRuleContext<titeParser::PostfixContext>(0);
}

titeParser::FactorContext* titeParser::PowerContext::factor() {
  return getRuleContext<titeParser::FactorContext>(0);
}

std::vector<tree::TerminalNode *> titeParser::PowerContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::PowerContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::PowerContext::getRuleIndex() const {
  return titeParser::RulePower;
}

void titeParser::PowerContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPower(this);
}

void titeParser::PowerContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPower(this);
}

titeParser::PowerContext* titeParser::power() {
  PowerContext *_localctx = _tracker.createInstance<PowerContext>(_ctx, getState());
  enterRule(_localctx, 40, titeParser::RulePower);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(329);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 50, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(318);
      postfix(0);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(319);
      postfix(0);
      setState(321);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == titeParser::LF) {
        setState(320);
        match(titeParser::LF);
      }
      setState(323);
      match(titeParser::T__29);
      setState(325);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == titeParser::LF) {
        setState(324);
        match(titeParser::LF);
      }
      setState(327);
      factor();
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PostfixContext ------------------------------------------------------------------

titeParser::PostfixContext::PostfixContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

titeParser::PrimaryContext* titeParser::PostfixContext::primary() {
  return getRuleContext<titeParser::PrimaryContext>(0);
}

titeParser::PostfixContext* titeParser::PostfixContext::postfix() {
  return getRuleContext<titeParser::PostfixContext>(0);
}

tree::TerminalNode* titeParser::PostfixContext::IDENTIFIER() {
  return getToken(titeParser::IDENTIFIER, 0);
}

titeParser::ArrayContext* titeParser::PostfixContext::array() {
  return getRuleContext<titeParser::ArrayContext>(0);
}

titeParser::InputContext* titeParser::PostfixContext::input() {
  return getRuleContext<titeParser::InputContext>(0);
}

tree::TerminalNode* titeParser::PostfixContext::LF() {
  return getToken(titeParser::LF, 0);
}


size_t titeParser::PostfixContext::getRuleIndex() const {
  return titeParser::RulePostfix;
}

void titeParser::PostfixContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPostfix(this);
}

void titeParser::PostfixContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPostfix(this);
}


titeParser::PostfixContext* titeParser::postfix() {
   return postfix(0);
}

titeParser::PostfixContext* titeParser::postfix(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  titeParser::PostfixContext *_localctx = _tracker.createInstance<PostfixContext>(_ctx, parentState);
  titeParser::PostfixContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 42;
  enterRecursionRule(_localctx, 42, titeParser::RulePostfix, precedence);

    size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(332);
    primary();
    _ctx->stop = _input->LT(-1);
    setState(348);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 53, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<PostfixContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RulePostfix);
        setState(334);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(336);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == titeParser::LF) {
          setState(335);
          match(titeParser::LF);
        }
        setState(344);
        _errHandler->sync(this);
        switch (_input->LA(1)) {
          case titeParser::T__25: {
            setState(338);
            match(titeParser::T__25);
            break;
          }

          case titeParser::T__26: {
            setState(339);
            match(titeParser::T__26);
            break;
          }

          case titeParser::T__30: {
            setState(340);
            match(titeParser::T__30);
            setState(341);
            match(titeParser::IDENTIFIER);
            break;
          }

          case titeParser::T__53: {
            setState(342);
            array();
            break;
          }

          case titeParser::T__31: {
            setState(343);
            input();
            break;
          }

        default:
          throw NoViableAltException(this);
        } 
      }
      setState(350);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 53, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- PrimaryContext ------------------------------------------------------------------

titeParser::PrimaryContext::PrimaryContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* titeParser::PrimaryContext::IDENTIFIER() {
  return getToken(titeParser::IDENTIFIER, 0);
}

titeParser::TagContext* titeParser::PrimaryContext::tag() {
  return getRuleContext<titeParser::TagContext>(0);
}

titeParser::LiteralContext* titeParser::PrimaryContext::literal() {
  return getRuleContext<titeParser::LiteralContext>(0);
}

titeParser::ExpressionContext* titeParser::PrimaryContext::expression() {
  return getRuleContext<titeParser::ExpressionContext>(0);
}


size_t titeParser::PrimaryContext::getRuleIndex() const {
  return titeParser::RulePrimary;
}

void titeParser::PrimaryContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterPrimary(this);
}

void titeParser::PrimaryContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitPrimary(this);
}

titeParser::PrimaryContext* titeParser::primary() {
  PrimaryContext *_localctx = _tracker.createInstance<PrimaryContext>(_ctx, getState());
  enterRule(_localctx, 44, titeParser::RulePrimary);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(360);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case titeParser::T__30:
      case titeParser::T__33:
      case titeParser::T__34:
      case titeParser::IDENTIFIER: {
        enterOuterAlt(_localctx, 1);
        setState(352);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if ((((_la & ~ 0x3fULL) == 0) &&
          ((1ULL << _la) & 53687091200) != 0)) {
          setState(351);
          tag();
        }
        setState(354);
        match(titeParser::IDENTIFIER);
        break;
      }

      case titeParser::T__51:
      case titeParser::T__53:
      case titeParser::INT:
      case titeParser::FLOAT:
      case titeParser::CHAR:
      case titeParser::STR: {
        enterOuterAlt(_localctx, 2);
        setState(355);
        literal();
        break;
      }

      case titeParser::T__31: {
        enterOuterAlt(_localctx, 3);
        setState(356);
        match(titeParser::T__31);
        setState(357);
        expression();
        setState(358);
        match(titeParser::T__32);
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TagContext ------------------------------------------------------------------

titeParser::TagContext::TagContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}


size_t titeParser::TagContext::getRuleIndex() const {
  return titeParser::RuleTag;
}

void titeParser::TagContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTag(this);
}

void titeParser::TagContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTag(this);
}

titeParser::TagContext* titeParser::tag() {
  TagContext *_localctx = _tracker.createInstance<TagContext>(_ctx, getState());
  enterRule(_localctx, 46, titeParser::RuleTag);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(362);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 53687091200) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- AssignContext ------------------------------------------------------------------

titeParser::AssignContext::AssignContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}


size_t titeParser::AssignContext::getRuleIndex() const {
  return titeParser::RuleAssign;
}

void titeParser::AssignContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAssign(this);
}

void titeParser::AssignContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAssign(this);
}

titeParser::AssignContext* titeParser::assign() {
  AssignContext *_localctx = _tracker.createInstance<AssignContext>(_ctx, getState());
  enterRule(_localctx, 48, titeParser::RuleAssign);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(364);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 4503530907893760) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- LiteralContext ------------------------------------------------------------------

titeParser::LiteralContext::LiteralContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* titeParser::LiteralContext::INT() {
  return getToken(titeParser::INT, 0);
}

tree::TerminalNode* titeParser::LiteralContext::FLOAT() {
  return getToken(titeParser::FLOAT, 0);
}

tree::TerminalNode* titeParser::LiteralContext::CHAR() {
  return getToken(titeParser::CHAR, 0);
}

tree::TerminalNode* titeParser::LiteralContext::STR() {
  return getToken(titeParser::STR, 0);
}

titeParser::BlockContext* titeParser::LiteralContext::block() {
  return getRuleContext<titeParser::BlockContext>(0);
}

titeParser::ArrayContext* titeParser::LiteralContext::array() {
  return getRuleContext<titeParser::ArrayContext>(0);
}


size_t titeParser::LiteralContext::getRuleIndex() const {
  return titeParser::RuleLiteral;
}

void titeParser::LiteralContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterLiteral(this);
}

void titeParser::LiteralContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitLiteral(this);
}

titeParser::LiteralContext* titeParser::literal() {
  LiteralContext *_localctx = _tracker.createInstance<LiteralContext>(_ctx, getState());
  enterRule(_localctx, 50, titeParser::RuleLiteral);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(372);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case titeParser::INT: {
        enterOuterAlt(_localctx, 1);
        setState(366);
        match(titeParser::INT);
        break;
      }

      case titeParser::FLOAT: {
        enterOuterAlt(_localctx, 2);
        setState(367);
        match(titeParser::FLOAT);
        break;
      }

      case titeParser::CHAR: {
        enterOuterAlt(_localctx, 3);
        setState(368);
        match(titeParser::CHAR);
        break;
      }

      case titeParser::STR: {
        enterOuterAlt(_localctx, 4);
        setState(369);
        match(titeParser::STR);
        break;
      }

      case titeParser::T__51: {
        enterOuterAlt(_localctx, 5);
        setState(370);
        block();
        break;
      }

      case titeParser::T__53: {
        enterOuterAlt(_localctx, 6);
        setState(371);
        array();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- BlockContext ------------------------------------------------------------------

titeParser::BlockContext::BlockContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> titeParser::BlockContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::BlockContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}

titeParser::SequenceContext* titeParser::BlockContext::sequence() {
  return getRuleContext<titeParser::SequenceContext>(0);
}


size_t titeParser::BlockContext::getRuleIndex() const {
  return titeParser::RuleBlock;
}

void titeParser::BlockContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterBlock(this);
}

void titeParser::BlockContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitBlock(this);
}

titeParser::BlockContext* titeParser::block() {
  BlockContext *_localctx = _tracker.createInstance<BlockContext>(_ctx, getState());
  enterRule(_localctx, 52, titeParser::RuleBlock);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(374);
    match(titeParser::T__51);
    setState(376);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 57, _ctx)) {
    case 1: {
      setState(375);
      match(titeParser::LF);
      break;
    }

    default:
      break;
    }
    setState(379);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 2260807003212349576) != 0)) {
      setState(378);
      sequence();
    }
    setState(382);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::LF) {
      setState(381);
      match(titeParser::LF);
    }
    setState(384);
    match(titeParser::T__52);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ArrayContext ------------------------------------------------------------------

titeParser::ArrayContext::ArrayContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> titeParser::ArrayContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::ArrayContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}

titeParser::SequenceContext* titeParser::ArrayContext::sequence() {
  return getRuleContext<titeParser::SequenceContext>(0);
}


size_t titeParser::ArrayContext::getRuleIndex() const {
  return titeParser::RuleArray;
}

void titeParser::ArrayContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterArray(this);
}

void titeParser::ArrayContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitArray(this);
}

titeParser::ArrayContext* titeParser::array() {
  ArrayContext *_localctx = _tracker.createInstance<ArrayContext>(_ctx, getState());
  enterRule(_localctx, 54, titeParser::RuleArray);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(386);
    match(titeParser::T__53);
    setState(388);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 60, _ctx)) {
    case 1: {
      setState(387);
      match(titeParser::LF);
      break;
    }

    default:
      break;
    }
    setState(391);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 2260807003212349576) != 0)) {
      setState(390);
      sequence();
    }
    setState(394);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::LF) {
      setState(393);
      match(titeParser::LF);
    }
    setState(396);
    match(titeParser::T__54);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- InputContext ------------------------------------------------------------------

titeParser::InputContext::InputContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> titeParser::InputContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::InputContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}

titeParser::SequenceContext* titeParser::InputContext::sequence() {
  return getRuleContext<titeParser::SequenceContext>(0);
}


size_t titeParser::InputContext::getRuleIndex() const {
  return titeParser::RuleInput;
}

void titeParser::InputContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterInput(this);
}

void titeParser::InputContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitInput(this);
}

titeParser::InputContext* titeParser::input() {
  InputContext *_localctx = _tracker.createInstance<InputContext>(_ctx, getState());
  enterRule(_localctx, 56, titeParser::RuleInput);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(398);
    match(titeParser::T__31);
    setState(400);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 63, _ctx)) {
    case 1: {
      setState(399);
      match(titeParser::LF);
      break;
    }

    default:
      break;
    }
    setState(403);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 2260807003212349576) != 0)) {
      setState(402);
      sequence();
    }
    setState(406);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::LF) {
      setState(405);
      match(titeParser::LF);
    }
    setState(408);
    match(titeParser::T__32);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DelimContext ------------------------------------------------------------------

titeParser::DelimContext::DelimContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> titeParser::DelimContext::LF() {
  return getTokens(titeParser::LF);
}

tree::TerminalNode* titeParser::DelimContext::LF(size_t i) {
  return getToken(titeParser::LF, i);
}


size_t titeParser::DelimContext::getRuleIndex() const {
  return titeParser::RuleDelim;
}

void titeParser::DelimContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterDelim(this);
}

void titeParser::DelimContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<titeListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitDelim(this);
}

titeParser::DelimContext* titeParser::delim() {
  DelimContext *_localctx = _tracker.createInstance<DelimContext>(_ctx, getState());
  enterRule(_localctx, 58, titeParser::RuleDelim);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(410);
    _la = _input->LA(1);
    if (!(_la == titeParser::T__1

    || _la == titeParser::LF)) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
    setState(412);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == titeParser::LF) {
      setState(411);
      match(titeParser::LF);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

bool titeParser::sempred(RuleContext *context, size_t ruleIndex, size_t predicateIndex) {
  switch (ruleIndex) {
    case 9: return disjunctionSempred(antlrcpp::downCast<DisjunctionContext *>(context), predicateIndex);
    case 10: return conjunctionSempred(antlrcpp::downCast<ConjunctionContext *>(context), predicateIndex);
    case 11: return orSempred(antlrcpp::downCast<OrContext *>(context), predicateIndex);
    case 12: return xorSempred(antlrcpp::downCast<XorContext *>(context), predicateIndex);
    case 13: return andSempred(antlrcpp::downCast<AndContext *>(context), predicateIndex);
    case 14: return equalitySempred(antlrcpp::downCast<EqualityContext *>(context), predicateIndex);
    case 15: return relationSempred(antlrcpp::downCast<RelationContext *>(context), predicateIndex);
    case 16: return shiftSempred(antlrcpp::downCast<ShiftContext *>(context), predicateIndex);
    case 17: return additiveSempred(antlrcpp::downCast<AdditiveContext *>(context), predicateIndex);
    case 18: return productSempred(antlrcpp::downCast<ProductContext *>(context), predicateIndex);
    case 21: return postfixSempred(antlrcpp::downCast<PostfixContext *>(context), predicateIndex);

  default:
    break;
  }
  return true;
}

bool titeParser::disjunctionSempred(DisjunctionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 0: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::conjunctionSempred(ConjunctionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 1: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::orSempred(OrContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 2: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::xorSempred(XorContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 3: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::andSempred(AndContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 4: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::equalitySempred(EqualityContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 5: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::relationSempred(RelationContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 6: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::shiftSempred(ShiftContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 7: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::additiveSempred(AdditiveContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 8: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::productSempred(ProductContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 9: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool titeParser::postfixSempred(PostfixContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 10: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

void titeParser::initialize() {
#if ANTLR4_USE_THREAD_LOCAL_CACHE
  titeParserInitialize();
#else
  ::antlr4::internal::call_once(titeParserOnceFlag, titeParserInitialize);
#endif
}
