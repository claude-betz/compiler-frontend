## Recursive descent parser for subset of java to three address code

To implement this we have to eliminate some of the left-recursive productions in the grammar

```
program  -> block

block    -> { decls stmts }

decls    -> decl decls
         |  ϵ

decl     -> type id

type     -> type [ num ] **
         |  basic

stmts    -> stmt stmts
         |  ϵ

stmt     -> loc = bool;
         |  if ( bool ) stmt
         |  if ( bool ) stmt else stmt **
         |  while ( bool ) stmt
         |  do stmt while ( bool ) ;
         |  block

loc      -> id restLoc
restLoc     -> [ bool ] restLoc
         |  ϵ

bool     -> join restBool
restBool -> || join restBool 
         |  ϵ

join     -> equality restJoin
restJoin -> && equality restJoin
         |  ϵ

equality -> rel restEquality
restEquality -> == rel restEquality
         |  != rel restEquality
         |  ϵ

rel      -> expr < expr 
         |  expr <= expr
         |  expr >= expr 
         |  expr > expr
         |  expr 

expr     -> term restExpr
restExpr -> + term restExpr
         |  - term restExpr
         |  ϵ

term     -> unary restTerm
restTerm -> * unary restTerm
         |  / unary restTerm
         |  ϵ

unary    -> ! unary
         |  - unary
         |  factor

factor   -> ( expr )
         |  num
         |  loc
         |  id
```