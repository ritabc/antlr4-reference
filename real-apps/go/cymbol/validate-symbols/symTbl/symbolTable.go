package symTbl

import (
    "errors"
    "fmt"
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
    resolve(name string) (Symbol, error)
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
    gs.Symbols[sym.Name] = sym
   // track the scope in each symbol
   sym.Scope = gs
}

func (gs GlobalScope) resolve(name string) (Symbol, error) {
    var sym Symbol
    sym, ok := gs.Symbols[name]
    if !ok {
        // if not here, check any enclosing scope
        if gs.EnclosingScope != nil {
            enclosing, err := gs.EnclosingScope.resolve(name)
            if err != nil {
                return enclosing, nil
            }
        }
        // if not found & no enclosing scope,
        return sym, errors.New(fmt.Sprintf("No symbol found of name %v", name))
    }
    return sym, nil
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
    ls.Symbols[sym.Name] = sym
    // track the scope in each symbol
    sym.Scope = ls
}

func (ls LocalScope) resolve(name string) (Symbol, error) {
    var sym Symbol
    sym, ok := ls.Symbols[name]
    if !ok {
        // if not here, check any enclosing scope
        if ls.EnclosingScope != nil {
            enclosing, err := ls.EnclosingScope.resolve(name)
            if err != nil {
                return enclosing, nil
            }
        }
        // if not found & no enclosing scope,
        return sym, errors.New(fmt.Sprintf("No symbol found of name %v", name))
    }
    return sym, nil
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
    fs.Arguments[sym.Name] = sym
    sym.Scope = fs
}

func (fs FunctionSymbol) resolve(name string) (Symbol, error) {
    var sym Symbol
    sym, ok := fs.Arguments[name]
    if !ok {
        // if not here, check any enclosing scope
        if fs.EnclosingScope != nil {
            enclosing, err := fs.EnclosingScope.resolve(name)
            if err != nil {
                return enclosing, nil
            }
        }
        // if not found & no enclosing scope,
        return sym, errors.New(fmt.Sprintf("No symbol found of name %v", name))
    }
    return sym, nil
}

func (fs FunctionSymbol) GetEnclosingScope() Scope {
    return fs.EnclosingScope
}

type VariableSymbol struct {
    Symbol Symbol
}

