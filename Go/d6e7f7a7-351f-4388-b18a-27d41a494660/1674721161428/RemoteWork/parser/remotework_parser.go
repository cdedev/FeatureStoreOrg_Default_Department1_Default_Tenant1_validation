// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721161428/RemoteWork.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RemoteWork

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

type RemoteWorkParser struct {
	*antlr.BaseParser
}

var remoteworkParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func remoteworkParserInit() {
	staticData := &remoteworkParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "REMOTEWORK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "remotework",
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

// RemoteWorkParserInit initializes any static state used to implement RemoteWorkParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRemoteWorkParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RemoteWorkParserInit() {
	staticData := &remoteworkParserStaticData
	staticData.once.Do(remoteworkParserInit)
}

// NewRemoteWorkParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRemoteWorkParser(input antlr.TokenStream) *RemoteWorkParser {
	RemoteWorkParserInit()
	this := new(RemoteWorkParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &remoteworkParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RemoteWork.g4"

	return this
}

// RemoteWorkParser tokens.
const (
	RemoteWorkParserEOF        = antlr.TokenEOF
	RemoteWorkParserREMOTEWORK = 1
	RemoteWorkParserOWNKEY     = 2
	RemoteWorkParserSPLITKEY   = 3
	RemoteWorkParserWS         = 4
)

// RemoteWorkParser rules.
const (
	RemoteWorkParserRULE_expression = 0
	RemoteWorkParserRULE_remotework = 1
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
	p.RuleIndex = RemoteWorkParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RemoteWorkParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Remotework() IRemoteworkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemoteworkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemoteworkContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RemoteWorkParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RemoteWorkListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RemoteWorkListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RemoteWorkParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RemoteWorkParserRULE_expression)

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
		p.Remotework()
	}
	{
		p.SetState(5)
		p.Match(RemoteWorkParserEOF)
	}

	return localctx
}

// IRemoteworkContext is an interface to support dynamic dispatch.
type IRemoteworkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRemoteworkContext differentiates from other interfaces.
	IsRemoteworkContext()
}

type RemoteworkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemoteworkContext() *RemoteworkContext {
	var p = new(RemoteworkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RemoteWorkParserRULE_remotework
	return p
}

func (*RemoteworkContext) IsRemoteworkContext() {}

func NewRemoteworkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemoteworkContext {
	var p = new(RemoteworkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RemoteWorkParserRULE_remotework

	return p
}

func (s *RemoteworkContext) GetParser() antlr.Parser { return s.parser }

func (s *RemoteworkContext) REMOTEWORK() antlr.TerminalNode {
	return s.GetToken(RemoteWorkParserREMOTEWORK, 0)
}

func (s *RemoteworkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoteworkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemoteworkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RemoteWorkListener); ok {
		listenerT.EnterRemotework(s)
	}
}

func (s *RemoteworkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RemoteWorkListener); ok {
		listenerT.ExitRemotework(s)
	}
}

func (p *RemoteWorkParser) Remotework() (localctx IRemoteworkContext) {
	this := p
	_ = this

	localctx = NewRemoteworkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RemoteWorkParserRULE_remotework)

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
		p.Match(RemoteWorkParserREMOTEWORK)
	}

	return localctx
}
