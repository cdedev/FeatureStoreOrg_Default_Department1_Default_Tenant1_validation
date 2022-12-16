// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671189790235/Package.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Package

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

type PackageParser struct {
	*antlr.BaseParser
}

var packageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func packageParserInit() {
	staticData := &packageParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PACKAGE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "package",
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

// PackageParserInit initializes any static state used to implement PackageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPackageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PackageParserInit() {
	staticData := &packageParserStaticData
	staticData.once.Do(packageParserInit)
}

// NewPackageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPackageParser(input antlr.TokenStream) *PackageParser {
	PackageParserInit()
	this := new(PackageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &packageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Package.g4"

	return this
}

// PackageParser tokens.
const (
	PackageParserEOF      = antlr.TokenEOF
	PackageParserPACKAGE  = 1
	PackageParserOWNKEY   = 2
	PackageParserSPLITKEY = 3
	PackageParserWS       = 4
)

// PackageParser rules.
const (
	PackageParserRULE_expression = 0
	PackageParserRULE_package    = 1
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
	p.RuleIndex = PackageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PackageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Package() IPackageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PackageParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PackageListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PackageListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PackageParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PackageParserRULE_expression)

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
		p.Package()
	}
	{
		p.SetState(5)
		p.Match(PackageParserEOF)
	}

	return localctx
}

// IPackageContext is an interface to support dynamic dispatch.
type IPackageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageContext differentiates from other interfaces.
	IsPackageContext()
}

type PackageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageContext() *PackageContext {
	var p = new(PackageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PackageParserRULE_package
	return p
}

func (*PackageContext) IsPackageContext() {}

func NewPackageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageContext {
	var p = new(PackageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PackageParserRULE_package

	return p
}

func (s *PackageContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(PackageParserPACKAGE, 0)
}

func (s *PackageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PackageListener); ok {
		listenerT.EnterPackage(s)
	}
}

func (s *PackageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PackageListener); ok {
		listenerT.ExitPackage(s)
	}
}

func (p *PackageParser) Package() (localctx IPackageContext) {
	this := p
	_ = this

	localctx = NewPackageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PackageParserRULE_package)

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
		p.Match(PackageParserPACKAGE)
	}

	return localctx
}
