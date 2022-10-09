# A  complete compiler front-end
Implementation of a complete compiler front-end as described in chapter 2 of the dragon book
for a subset of the functionality in the java language as defined by the grammar below:

```
program  -> block

block    -> { decls stmts }

decls    -> decls decl
         |  ϵ

stmts    -> stmts stmt
         |  ϵ

stmt     -> expr ;
         |  if ( expr ) stmt
         |  if ( expr ) stmt else stmt
         |  while ( expr ) stmt
         |  do stmt while ( expr ) ;
         |  block

expr     -> loc = bool

loc      -> loc[ bool ]
         |  id

expr     -> loc = bool
         |  bool

bool     -> bool || join 
         | join

join     -> join && equality
         | equality

equality -> equality == rel
         |  equality != rel
         |  rel

rel      -> expr < expr 
         |  expr <= expr
         |  expr >= expr 
         |  expr > expr
         |  expr 

expr     -> expr + term 
         |  expr - term
         |  term

term     -> term * factor
         |  term / factor
         |  factor

factor   -> ( expr )
         |  num
         |  id
```

