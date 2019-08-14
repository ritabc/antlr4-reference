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

func (s Symbol) toString() string {
    if s.SymbolType == INVALID_T {
        return s.Name
    }
    return "<" + s.Name + ":" + strconv.Itoa(s.SymbolType) + ">"
}

type Scope interface {
    //getScopeName() string
    //getEnclopsingScope() scope
    Define(sym Symbol)
    Resolve(name string) (Symbol, string)
    GetEnclosingScope() Scope
}

//type BaseScope struct { // implements scope
//    EnclosingScope Scope // nil if global (outermost) scope
//    Symbols map[string]Symbol
//}

//func (bs BaseScope) Define(sym Symbol) {
//    bs.Symbols[sym.Name] = sym
//    // track the scope in each symbol
//    sym.Scope = bs
//}
//
//func (bs BaseScope) resolve(name string) (Symbol, error) {
//    var sym Symbol
//    sym, ok := bs.Symbols[name]
//    if !ok {
//        // if not here, check any enclosing scope
//        if bs.EnclosingScope != nil {
//            enclosing, err := bs.EnclosingScope.resolve(name)
//            if err != nil {
//                return enclosing, nil
//            }
//        }
//        // if not found & no enclosing scope,
//        return sym, errors.New(fmt.Sprintf("No symbol found of name %v", name))
//    }
//    return sym, nil
//}
//
//func (bs BaseScope) GetEnclosingScope() Scope {
//    return bs.EnclosingScope
//}

type GlobalScope struct {
    Symbols map[string]Symbol
    EnclosingScope Scope // this one is nil, since it's GlobalScope. TODO: set to nil upon initialization?
}

func (gs GlobalScope) Define(sym Symbol) {
   // track the scope in each symbol
   sym.Scope = gs
    gs.Symbols[sym.Name] = sym
}

func (gs GlobalScope) Resolve(name string) (Symbol, string) {
    var sym Symbol
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
    Symbols map[string]Symbol
    EnclosingScope Scope
}

func (ls LocalScope) Define(sym Symbol) {
    sym.Scope = ls
    ls.Symbols[sym.Name] = sym
    // track the scope in each symbol
}

func (ls LocalScope) Resolve(name string) (Symbol, string) {
    var sym Symbol
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
type FunctionSymbol struct { // implements Scope
    Symbol
    Arguments map[string]Symbol
    EnclosingScope Scope
}

func (fs FunctionSymbol) Define(sym Symbol) {
    // track the scope in each symbol. TODO: should sym come in as a pointer?
    sym.Scope = fs
    fs.Arguments[sym.Name] = sym
}

func (fs FunctionSymbol) Resolve(name string) (Symbol, string) {
    var sym Symbol
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

type VariableSymbol struct {
    Symbol Symbol
}

