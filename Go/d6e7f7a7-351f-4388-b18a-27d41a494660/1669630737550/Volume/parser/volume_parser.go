// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669630737550/Volume.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Volume

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

type VolumeParser struct {
	*antlr.BaseParser
}

var volumeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func volumeParserInit() {
	staticData := &volumeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VOLUME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "volume",
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

// VolumeParserInit initializes any static state used to implement VolumeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVolumeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VolumeParserInit() {
	staticData := &volumeParserStaticData
	staticData.once.Do(volumeParserInit)
}

// NewVolumeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVolumeParser(input antlr.TokenStream) *VolumeParser {
	VolumeParserInit()
	this := new(VolumeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &volumeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Volume.g4"

	return this
}

// VolumeParser tokens.
const (
	VolumeParserEOF      = antlr.TokenEOF
	VolumeParserVOLUME   = 1
	VolumeParserOWNKEY   = 2
	VolumeParserSPLITKEY = 3
	VolumeParserWS       = 4
)

// VolumeParser rules.
const (
	VolumeParserRULE_expression = 0
	VolumeParserRULE_volume     = 1
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
	p.RuleIndex = VolumeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VolumeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Volume() IVolumeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVolumeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVolumeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VolumeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VolumeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VolumeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VolumeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VolumeParserRULE_expression)

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
		p.Volume()
	}
	{
		p.SetState(5)
		p.Match(VolumeParserEOF)
	}

	return localctx
}

// IVolumeContext is an interface to support dynamic dispatch.
type IVolumeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVolumeContext differentiates from other interfaces.
	IsVolumeContext()
}

type VolumeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVolumeContext() *VolumeContext {
	var p = new(VolumeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VolumeParserRULE_volume
	return p
}

func (*VolumeContext) IsVolumeContext() {}

func NewVolumeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeContext {
	var p = new(VolumeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VolumeParserRULE_volume

	return p
}

func (s *VolumeContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeContext) VOLUME() antlr.TerminalNode {
	return s.GetToken(VolumeParserVOLUME, 0)
}

func (s *VolumeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VolumeListener); ok {
		listenerT.EnterVolume(s)
	}
}

func (s *VolumeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VolumeListener); ok {
		listenerT.ExitVolume(s)
	}
}

func (p *VolumeParser) Volume() (localctx IVolumeContext) {
	this := p
	_ = this

	localctx = NewVolumeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VolumeParserRULE_volume)

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
		p.Match(VolumeParserVOLUME)
	}

	return localctx
}
