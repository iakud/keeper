// Define a grammar called kds
grammar kds;

kds
	: packageStatement (importStatement | topLevelDef | emptyStatement_)* EOF
	;

// Package

packageStatement
	: PACKAGE fullIdent SEMI
	;

// Import Statement

importStatement
	: IMPORT STR_LIT SEMI
	;

// Normal Field

field
	: fieldLabel? type_ fieldName EQ fieldNumber (LB fieldOptions RB)? SEMI
	;

fieldLabel
	: REPEATED
	;

fieldOptions
	: fieldOption (COMMA fieldOption)*
	;

fieldOption
	: fullIdent
	;

fieldNumber
	: intLit
	;

// Map field

mapField
	: MAP LT keyType COMMA type_ GT mapName EQ fieldNumber (LB fieldOptions RB)? SEMI
	;

keyType
	: INT32
	| INT64
	| UINT32
	| UINT64
	| SINT32
	| SINT64
	| FIXED32
	| FIXED64
	| SFIXED32
	| SFIXED64
	| BOOL
	| STRING
	;

// field types

type_
	: DOUBLE
	| FLOAT
	| INT32
	| INT64
	| UINT32
	| UINT64
	| SINT32
	| SINT64
	| FIXED32
	| FIXED64
	| SFIXED32
	| SFIXED64
	| BOOL
	| STRING
	| BYTES
	| TIMESTAMP
	| DURATION
	| messageType
	| enumType
	;

// Top Level definitions

topLevelDef
	: enumDef
	| entityDef
	| componentDef
	;

// enum

enumDef
	: ENUM enumName enumBody
	;

enumBody
	: LC enumElement* RC
	;

enumElement
	: enumField
	;

enumField
	: ident EQ (MINUS)? intLit SEMI
	;

// entity

entityDef
	: ENTITY entityName entityBody
	;

entityName
	: ident
	;

entityBody
	: LC entityElement* RC
	;

entityElement
	: field
	| mapField
	| emptyStatement_
	;

// component

componentDef
	: COMPONENT componentName componentBody
	;

componentName
	: ident
	;

componentBody
	: LC componentElement* RC
	;

componentElement
	: field
	| mapField
	| emptyStatement_
	;

// lexical

emptyStatement_
	: SEMI
	;

// Lexical elements

ident
	: IDENTIFIER
	| keywords
	;

fullIdent
	: ident (DOT ident)*
	;

fieldName
	: ident
	;

messageName
	: ident
	;

enumName
	: ident
	;

mapName
	: ident
	;

messageType
	: (DOT)? (ident DOT)* messageName
	;

enumType
	: (DOT)? (ident DOT)* enumName
	;

intLit
	: INT_LIT
	;

// keywords

IMPORT
	: 'import'
	;

PACKAGE
	: 'package'
	;

REPEATED
	: 'repeated'
	;

MAP
	: 'map'
	;

INT32
	: 'int32'
	;

INT64
	: 'int64'
	;

UINT32
	: 'uint32'
	;

UINT64
	: 'uint64'
	;

SINT32
	: 'sint32'
	;

SINT64
	: 'sint64'
	;

FIXED32
	: 'fixed32'
	;

FIXED64
	: 'fixed64'
	;

SFIXED32
	: 'sfixed32'
	;

SFIXED64
	: 'sfixed64'
	;

BOOL
	: 'bool'
	;

STRING
	: 'string'
	;

DOUBLE
	: 'double'
	;

FLOAT
	: 'float'
	;

BYTES
	: 'bytes'
	;

TIMESTAMP
	: 'timestamp'
	;

DURATION
	: 'duration'
	;

ENUM
	: 'enum'
	;

ENTITY
	: 'entity'
	;

COMPONENT
	: 'component'
	;

// symbols

SEMI
	: ';'
	;

EQ
	: '='
	;

LP
	: '('
	;

RP
	: ')'
	;

LB
	: '['
	;

RB
	: ']'
	;

LC
	: '{'
	;

RC
	: '}'
	;

LT
	: '<'
	;

GT
	: '>'
	;

DOT
	: '.'
	;

COMMA
	: ','
	;

COLON
	: ':'
	;

PLUS
	: '+'
	;

MINUS
	: '-'
	;

STR_LIT
	: ('\'' ( CHAR_VALUE)*? '\'')
	| ( '"' ( CHAR_VALUE)*? '"')
	;

fragment CHAR_VALUE
	: HEX_ESCAPE
	| OCT_ESCAPE
	| CHAR_ESCAPE
	| ~[\u0000\n\\]
	;

fragment HEX_ESCAPE
	: '\\' ('x' | 'X') HEX_DIGIT HEX_DIGIT
	;

fragment OCT_ESCAPE
	: '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
	;

fragment CHAR_ESCAPE
	: '\\' ('a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '\\' | '\'' | '"')
	;

BOOL_LIT
	: 'true'
	| 'false'
	;

INT_LIT
	: DECIMAL_LIT
	| OCTAL_LIT
	| HEX_LIT
	;

fragment DECIMAL_LIT
	: ([1-9]) DECIMAL_DIGIT*
	;

fragment OCTAL_LIT
	: '0' OCTAL_DIGIT*
	;

fragment HEX_LIT
	: '0' ('x' | 'X') HEX_DIGIT+
	;

IDENTIFIER
	: LETTER (LETTER | DECIMAL_DIGIT)*
	;

fragment LETTER
	: [A-Za-z_]
	;

fragment DECIMAL_DIGIT
	: [0-9]
	;

fragment OCTAL_DIGIT
	: [0-7]
	;

fragment HEX_DIGIT
	: [0-9A-Fa-f]
	;

// comments
WS
	: [ \t\r\n\u000C]+ -> skip
	;

LINE_COMMENT
	: '//' ~[\r\n]* -> channel(HIDDEN)
	;

COMMENT
	: '/*' .*? '*/' -> channel(HIDDEN)
	;

keywords
	: IMPORT
	| PACKAGE
	| MAP
	| INT32
	| INT64
	| UINT32
	| UINT64
	| SINT32
	| SINT64
	| FIXED32
	| FIXED64
	| SFIXED32
	| SFIXED64
	| BOOL
	| STRING
	| DOUBLE
	| FLOAT
	| BYTES
	| ENUM
	| ENTITY
	| COMPONENT
	| BOOL_LIT
	;