package symTbl

import (
    "fmt"
    "github.com/sanity-io/litter"
    "strconv"
)

const (
    INVALID_T = iota
    VOID_T
    INT_T
    FLOAT_T
)

type Symbol struct {
    Name string
    Scope Scope // All symbols know which scope contains them
    SymbolType int
}

//func (s Symbol) toString() string {
//    if s.SymbolType == INVALID_T {
//        return s.Name
//    }
//    return "<" + s.Name + ":" + strconv.Itoa(s.SymbolType) + ">"
//}

type Scope interface {
    Define(sym ISymbol)
    Resolve(name string) (ISymbol, string)
    GetEnclosingScope() Scope
}

type ISymbol interface {
    ToString() string
}

type GlobalScope struct {
    Symbols map[string]ISymbol
    EnclosingScope Scope // this one is nil, since it's GlobalScope. TODO: set to nil upon initialization?
}

func (gs GlobalScope) Define(sym ISymbol) {
    switch sym.(type) {
    case FunctionSymbol:
        dupFnSym := sym.(FunctionSymbol)
        dupFnSym.Symbol.Scope = gs
        gs.Symbols[dupFnSym.Symbol.Name] = dupFnSym
    case VariableSymbol:
        dupVarSym := sym.(VariableSymbol)
        dupVarSym.Symbol.Scope = gs
        gs.Symbols[dupVarSym.Symbol.Name] = dupVarSym
    default:
        litter.Dump("SOMETHING WENT WRONG")
    }
}

func (gs GlobalScope) Resolve(name string) (ISymbol, string) {
    var sym ISymbol
    sym, ok := gs.Symbols[name]
    if !ok {
        // if not here, check any enclosing scope
        if gs.EnclosingScope != nil {
            // Should never reach since gs is Outermost
            litter.Dump("SOMETHING WENT WRONG")
        }
        // if not found & no enclosing scope,
        return sym, fmt.Sprintf("No symbol found of name %v", name)
    }
    return sym, ""
}

func (gs GlobalScope) GetEnclosingScope() Scope {
    return gs.EnclosingScope
}

func (gs GlobalScope) getScopeName() string {
    return "globals"
}

type LocalScope struct {
    Symbols map[string]ISymbol
    EnclosingScope Scope
}

func (ls LocalScope) Define(sym ISymbol) {
    switch sym.(type) {
    case FunctionSymbol:
        dupFnSym := sym.(FunctionSymbol)
        dupFnSym.Symbol.Scope = ls
        ls.Symbols[dupFnSym.Symbol.Name] = dupFnSym
    case VariableSymbol:
        dupVarSym := sym.(VariableSymbol)
        dupVarSym.Symbol.Scope = ls
        ls.Symbols[dupVarSym.Symbol.Name] = dupVarSym
    default:
        litter.Dump("SOMETHING WENT WRONG")
    }
}

func (ls LocalScope) Resolve(name string) (ISymbol, string) {
    var sym ISymbol
    sym, ok := ls.Symbols[name]
    if !ok {
        // if not here, check any enclosing scope
        if ls.EnclosingScope != nil {
            symInEnclosingScope, err := ls.EnclosingScope.Resolve(name)
            if err == "" {
                return symInEnclosingScope, ""
            }
        }
        // if not found & no enclosing scope,
        return sym, fmt.Sprintf("No symbol found of name %v", name)
    }
    return sym, ""
}

func (ls LocalScope) GetEnclosingScope() Scope {
    return ls.EnclosingScope
}

func (ls LocalScope) getScopeName() string {
    return "locals"
}

// FunctionSymbol plays the role of a symbol and the scope which holds the arguments
type FunctionSymbol struct { // implements Scope & ISymbol
    Symbol Symbol
    Arguments map[string]ISymbol
    EnclosingScope Scope
}

func (fs FunctionSymbol) Define(sym ISymbol) {
    switch sym.(type) {
    case FunctionSymbol:
        dupFnSym := sym.(FunctionSymbol)
        dupFnSym.Symbol.Scope = fs
        fs.Arguments[dupFnSym.Symbol.Name] = dupFnSym
    case VariableSymbol:
        dupVarSym := sym.(VariableSymbol)
        dupVarSym.Symbol.Scope = fs
        fs.Arguments[dupVarSym.Symbol.Name] = dupVarSym
    default:
        litter.Dump("SOMETHING WENT WRONG")
    }
}

func (fs FunctionSymbol) Resolve(name string) (ISymbol, string) {
    var sym ISymbol
    sym, ok := fs.Arguments[name]
    if !ok {
        // if not here, check any enclosing scope
        if fs.EnclosingScope != nil {
            symInEnclosingScope, err := fs.EnclosingScope.Resolve(name)
            if err == "" {
                return symInEnclosingScope, ""
            }
        }
        // if not found & no enclosing scope,
        return sym, fmt.Sprintf("No symbol found of name %v", name)
    }
    return sym, ""
}

func (fs FunctionSymbol) GetEnclosingScope() Scope {
    return fs.EnclosingScope
}

func (fs FunctionSymbol) ToString() string {
    if fs.Symbol.SymbolType == INVALID_T {
        return fs.Symbol.Name
    }
    return "<" + fs.Symbol.Name + ":" + strconv.Itoa(fs.Symbol.SymbolType) + ">"


}

type VariableSymbol struct { // implements ISymbol
    Symbol Symbol
}

func (vs VariableSymbol) ToString() string {
    if vs.Symbol.SymbolType == INVALID_T {
        return vs.Symbol.Name
    }
    return "<" + vs.Symbol.Name + ":" + strconv.Itoa(vs.Symbol.SymbolType) + ">"
}





