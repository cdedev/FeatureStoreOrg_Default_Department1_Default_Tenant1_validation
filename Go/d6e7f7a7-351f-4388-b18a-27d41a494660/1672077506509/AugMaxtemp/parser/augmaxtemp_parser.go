// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077506509/AugMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AugMaxtemp

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

type AugMaxtempParser struct {
	*antlr.BaseParser
}

var augmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func augmaxtempParserInit() {
	staticData := &augmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AUGMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "augmaxtemp",
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

// AugMaxtempParserInit initializes any static state used to implement AugMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAugMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AugMaxtempParserInit() {
	staticData := &augmaxtempParserStaticData
	staticData.once.Do(augmaxtempParserInit)
}

// NewAugMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAugMaxtempParser(input antlr.TokenStream) *AugMaxtempParser {
	AugMaxtempParserInit()
	this := new(AugMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &augmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AugMaxtemp.g4"

	return this
}

// AugMaxtempParser tokens.
const (
	AugMaxtempParserEOF        = antlr.TokenEOF
	AugMaxtempParserAUGMAXTEMP = 1
	AugMaxtempParserOWNKEY     = 2
	AugMaxtempParserSPLITKEY   = 3
	AugMaxtempParserWS         = 4
)

// AugMaxtempParser rules.
const (
	AugMaxtempParserRULE_expression = 0
	AugMaxtempParserRULE_augmaxtemp = 1
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
	p.RuleIndex = AugMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AugMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Augmaxtemp() IAugmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAugmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAugmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AugMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AugMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AugMaxtempParserRULE_expression)

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
		p.Augmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(AugMaxtempParserEOF)
	}

	return localctx
}

// IAugmaxtempContext is an interface to support dynamic dispatch.
type IAugmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAugmaxtempContext differentiates from other interfaces.
	IsAugmaxtempContext()
}

type AugmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAugmaxtempContext() *AugmaxtempContext {
	var p = new(AugmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AugMaxtempParserRULE_augmaxtemp
	return p
}

func (*AugmaxtempContext) IsAugmaxtempContext() {}

func NewAugmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AugmaxtempContext {
	var p = new(AugmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AugMaxtempParserRULE_augmaxtemp

	return p
}

func (s *AugmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *AugmaxtempContext) AUGMAXTEMP() antlr.TerminalNode {
	return s.GetToken(AugMaxtempParserAUGMAXTEMP, 0)
}

func (s *AugmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AugmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AugmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMaxtempListener); ok {
		listenerT.EnterAugmaxtemp(s)
	}
}

func (s *AugmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMaxtempListener); ok {
		listenerT.ExitAugmaxtemp(s)
	}
}

func (p *AugMaxtempParser) Augmaxtemp() (localctx IAugmaxtempContext) {
	this := p
	_ = this

	localctx = NewAugmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AugMaxtempParserRULE_augmaxtemp)

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
		p.Match(AugMaxtempParserAUGMAXTEMP)
	}

	return localctx
}
