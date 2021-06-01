package tree

import "fmt"

// IdentifierName is referenced in the expression
type IdentifierName interface {
	Expr
}

//sql indentifier
type Identifier string

//
type UnrestrictedIdentifier string

//the list of identifiers.
type IdentifierList []Identifier

type ColumnItem struct {
	IdentifierName

	//the name of the column
	ColumnName Identifier
}

//the unresolved qualified name like column name.
type UnresolvedName struct {
	exprImpl
	//the number of name parts specified, including the star. Always 1 or greater.
	NumParts int

	//the name ends with a star. then the first element is empty in the Parts
	Star bool

	// Parts are the name components (at most 4: column, table, schema, catalog/db.), in reverse order.
	Parts NameParts
}

//the path in an UnresolvedName.
type NameParts = [4]string

func NewUnresolvedName(parts ...string)(*UnresolvedName,error){
	l:=len(parts)
	if l < 1 || l > 4{
		return nil,fmt.Errorf("the count of name parts among [1,4]")
	}
	u:= &UnresolvedName{
		NumParts: len(parts),
		Star:     false,
	}
	for i:=0 ; i < len(parts);i++{
		u.Parts[i] = parts[l - 1 - i]
	}
	return u,nil
}

func NewUnresolvedNameWithStar(parts ...string)(*UnresolvedName,error){
	l:=len(parts)
	if l < 1 || l > 3{
		return nil,fmt.Errorf("the count of name parts among [1,3]")
	}
	u:= &UnresolvedName{
		NumParts: 1+len(parts),
		Star:     true,
	}
	u.Parts[0] = ""
	for i:=0 ; i < len(parts);i++{
		u.Parts[i+1] = parts[l - 1 - i]
	}
	return u,nil
}
//variable in the scalar expression
type VarName interface {
	Expr
}

var _ VarName = &UnresolvedName{}
var _ VarName = UnqualifiedStar{}

//'*' in the scalar expression
type UnqualifiedStar struct{
	VarName
}

var starName VarName = UnqualifiedStar{}

func StarExpr()VarName{
	return starName
}