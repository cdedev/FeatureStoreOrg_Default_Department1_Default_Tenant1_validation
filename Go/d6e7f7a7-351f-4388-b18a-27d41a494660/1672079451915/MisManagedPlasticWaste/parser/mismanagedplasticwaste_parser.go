// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672079451915/MisManagedPlasticWaste.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MisManagedPlasticWaste

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MisManagedPlasticWasteParser struct {
	*antlr.BaseParser
}

var mismanagedplasticwasteParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mismanagedplasticwasteParserInit() {
	staticData := &mismanagedplasticwasteParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MISMANAGEDPLASTICWASTE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mismanagedplasticwaste",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 10, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 7, 0, 4, 1, 0, 0, 0, 2, 7, 1, 0, 0, 0, 4, 5, 3, 2,
		1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 5, 1, 0, 0, 8, 3, 1, 0,
		0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MisManagedPlasticWasteParserInit initializes any static state used to implement MisManagedPlasticWasteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMisManagedPlasticWasteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MisManagedPlasticWasteParserInit() {
	staticData := &mismanagedplasticwasteParserStaticData
	staticData.once.Do(mismanagedplasticwasteParserInit)
}

// NewMisManagedPlasticWasteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMisManagedPlasticWasteParser(input antlr.TokenStream) *MisManagedPlasticWasteParser {
	MisManagedPlasticWasteParserInit()
	this := new(MisManagedPlasticWasteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mismanagedplasticwasteParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MisManagedPlasticWaste.g4"

	return this
}

// MisManagedPlasticWasteParser tokens.
const (
	MisManagedPlasticWasteParserEOF                    = antlr.TokenEOF
	MisManagedPlasticWasteParserMISMANAGEDPLASTICWASTE = 1
	MisManagedPlasticWasteParserOWNKEY                 = 2
	MisManagedPlasticWasteParserSPLITKEY               = 3
	MisManagedPlasticWasteParserWS                     = 4
)

// MisManagedPlasticWasteParser rules.
const (
	MisManagedPlasticWasteParserRULE_expression             = 0
	MisManagedPlasticWasteParserRULE_mismanagedplasticwaste = 1
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MisManagedPlasticWasteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MisManagedPlasticWasteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mismanagedplasticwaste() IMismanagedplasticwasteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMismanagedplasticwasteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMismanagedplasticwasteContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MisManagedPlasticWasteParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MisManagedPlasticWasteListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MisManagedPlasticWasteListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MisManagedPlasticWasteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MisManagedPlasticWasteParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Mismanagedplasticwaste()
	}
	{
		p.SetState(5)
		p.Match(MisManagedPlasticWasteParserEOF)
	}

	return localctx
}

// IMismanagedplasticwasteContext is an interface to support dynamic dispatch.
type IMismanagedplasticwasteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMismanagedplasticwasteContext differentiates from other interfaces.
	IsMismanagedplasticwasteContext()
}

type MismanagedplasticwasteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMismanagedplasticwasteContext() *MismanagedplasticwasteContext {
	var p = new(MismanagedplasticwasteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MisManagedPlasticWasteParserRULE_mismanagedplasticwaste
	return p
}

func (*MismanagedplasticwasteContext) IsMismanagedplasticwasteContext() {}

func NewMismanagedplasticwasteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MismanagedplasticwasteContext {
	var p = new(MismanagedplasticwasteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MisManagedPlasticWasteParserRULE_mismanagedplasticwaste

	return p
}

func (s *MismanagedplasticwasteContext) GetParser() antlr.Parser { return s.parser }

func (s *MismanagedplasticwasteContext) MISMANAGEDPLASTICWASTE() antlr.TerminalNode {
	return s.GetToken(MisManagedPlasticWasteParserMISMANAGEDPLASTICWASTE, 0)
}

func (s *MismanagedplasticwasteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MismanagedplasticwasteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MismanagedplasticwasteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MisManagedPlasticWasteListener); ok {
		listenerT.EnterMismanagedplasticwaste(s)
	}
}

func (s *MismanagedplasticwasteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MisManagedPlasticWasteListener); ok {
		listenerT.ExitMismanagedplasticwaste(s)
	}
}

func (p *MisManagedPlasticWasteParser) Mismanagedplasticwaste() (localctx IMismanagedplasticwasteContext) {
	this := p
	_ = this

	localctx = NewMismanagedplasticwasteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MisManagedPlasticWasteParserRULE_mismanagedplasticwaste)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)
		p.Match(MisManagedPlasticWasteParserMISMANAGEDPLASTICWASTE)
	}

	return localctx
}
