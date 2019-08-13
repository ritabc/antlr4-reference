package main

import (
    cParser "Projects/src/ANTLR/text-follow-along/real-apps/go/cymbol/validate-symbols/Cymbol"
    "Projects/src/ANTLR/text-follow-along/real-apps/go/cymbol/validate-symbols/symTbl"
    "github.com/antlr/antlr4/runtime/Go/antlr"
    "github.com/sanity-io/litter"
    "os"
)

type defPhaseListener struct {
    *cParser.BaseCymbolListener

    scopes map[antlr.ParseTree]symTbl.Scope
    globals symTbl.GlobalScope
    currentScope symTbl.Scope // needs to be a pointer?
}

func (dPL *defPhaseListener) EnterFile(ctx *cParser.FileContext) {
    dPL.scopes = make(map[antlr.ParseTree]symTbl.Scope)
    dPL.globals = symTbl.GlobalScope{
        Symbols: make(map[string]symTbl.Symbol),
    }
    //dPL.currentScope = dPL.globals.(symTbl.Scope) // ??? proper use of Enclosing Scope?
    dPL.currentScope = dPL.globals
}

func (dPL *defPhaseListener) ExitFile(ctx *cParser.FileContext) {
    litter.Dump(dPL.globals)
}

func (dPL *defPhaseListener) EnterFunctionDecl(ctx *cParser.FunctionDeclContext) {
    name := ctx.ID().GetText()
    typeTokenType := ctx.CymbolType().GetStart().GetTokenType()
    symType := getType(typeTokenType)

    // Make new functionSymbol that points to enclosing scope
    var function = symTbl.FunctionSymbol{
        Symbol: symTbl.Symbol{
            Name: name,
            SymbolType: symType,
        },
        Arguments: make(map[string]symTbl.Symbol),
        EnclosingScope: dPL.currentScope,
    }
    dPL.currentScope.Define(function.Symbol) // Add (define) function to current scope. This processes function as a symbol
    dPL.saveScope(ctx, function) // Push: set function's parent to current
    dPL.currentScope = function // Current scope is now function scope
}

func (dPL *defPhaseListener) ExitFunctionDecl(ctx *cParser.FunctionDeclContext) {
    litter.Dump(dPL.currentScope)
    dPL.currentScope = dPL.currentScope.GetEnclosingScope() // pop currentScope into oblivion, move up stack
}

func getType(tokenType int) int {
   switch tokenType {
   case cParser.CymbolParserT_VOID:
       return symTbl.VOID_T
   case cParser.CymbolParserT_INT:
       return symTbl.INT_T
   case cParser.CymbolParserT_FLOAT:
       return symTbl.FLOAT_T
   default:
       return symTbl.INVALID_T
   }
}

// saveScope annotates the functionDecl rule node with the fn scope so our reference phase can pick it up later
func (dPL *defPhaseListener) saveScope(ctx antlr.ParserRuleContext, s symTbl.Scope) {
    dPL.scopes[ctx] = s
}

func (dPL *defPhaseListener) EnterBlock(ctx *cParser.BlockContext) {
    // Make new localscope that points to enclosing scope (aka parent aka currentScope)
    var local = symTbl.LocalScope{
        Symbols: make(map[string]symTbl.Symbol),
        EnclosingScope: dPL.currentScope,
    }

    //switch v := dPL.currentScope.(type) {
    //case symTbl.FunctionSymbol:
    //    local = symTbl.LocalScope{
    //        BaseScope: dPL.currentScope.(symTbl.FunctionSymbol).EnclosingScope.(symTbl.BaseScope), // is this a correct use of enclosing scope?
    //    }
    //default:
    //    fmt.Printf("dPL.currentScope is type: %t, value: %v\n", v, v) // TODO: pass block statement in as input, update this line
    //}
    dPL.currentScope = local
    dPL.saveScope(ctx, dPL.currentScope) // push new local scope
}

func (dPL *defPhaseListener) ExitBlock(ctx *cParser.BlockContext) {
    litter.Dump(dPL.currentScope)
    dPL.currentScope = dPL.currentScope.GetEnclosingScope() // pop currentScope into oblivion (overwrite it), move up stack
}

func (dPL *defPhaseListener) ExitFormalParameter(ctx *cParser.FormalParameterContext) {
    dPL.defineVar(ctx.CymbolType(), ctx.ID().GetSymbol())
}

func (dPL *defPhaseListener) ExitVarDecl(ctx *cParser.VarDeclContext) {
    dPL.defineVar(ctx.CymbolType(), ctx.ID().GetSymbol())
}

func (dPL *defPhaseListener) defineVar(typeCtx cParser.ICymbolTypeContext, nameToken antlr.Token) {
    typeTokenType := typeCtx.GetStart().GetTokenType()
    symType := getType(typeTokenType)
    var varSymbol = symTbl.VariableSymbol{
        Symbol: symTbl.Symbol{
            Name: nameToken.GetText(),
            SymbolType: symType,
        },
    }
    dPL.currentScope.Define(varSymbol.Symbol) // Define symbol in current scope
}

func main() {
    input, _ := antlr.NewFileStream(os.Args[1])
    lexer := cParser.NewCymbolLexer(input)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := cParser.NewCymbolParser(tokens)
    tree := parser.File()
    walker := antlr.NewParseTreeWalker()
    defPhL := &defPhaseListener{}
    walker.Walk(defPhL, tree)
}