# A  complete compiler front-end
Implementation of a complete compiler front-end as described in chapter 2 of the dragon book
for a subset of the functionality in the java language as defined by the grammar below:

```
program  -> block

block    -> { decls stmts }

decls    -> decls decl
         |  ϵ

decl     -> type id

type     -> type [ num ] 
         |  basic

stmts    -> stmts stmt
         |  ϵ

stmt     -> loc = bool;
         |  if ( expr ) stmt
         |  if ( expr ) stmt else stmt
         |  while ( expr ) stmt
         |  do stmt while ( expr ) ;
         |  block

loc      -> loc[ bool ]
         |  id

bool     -> bool || join 
         | join

join     -> join && equality
         |  equality

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

term     -> term * unary
         |  term / unary
         |  unary

unary    -> ! unary
         |  - unary
         |  factor

factor   -> ( expr )
         |  num
         |  id
```

