//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package expression

/*
The type of Visitor is an interface with a list of methods that are
implemented in Stringer.go. The general class of methods fall under
are Arithmetic and Case Expressions, collections, concat, constant,
identifier, construction, logic, navigation, and function all of
which take an input parameter of that type and return an interface
and an error.
*/
type Visitor interface {
	/*
	   Arithmetic Expressions.
	*/
	VisitAdd(expr *Add) (interface{}, error)
	VisitDiv(expr *Div) (interface{}, error)
	VisitMod(expr *Mod) (interface{}, error)
	VisitMult(expr *Mult) (interface{}, error)
	VisitNeg(expr *Neg) (interface{}, error)
	VisitSub(expr *Sub) (interface{}, error)

	/*
	   Case Expressions. There are two types, searched case and
	   simple case expressions. Refer to N1QL specs.
	*/
	VisitSearchedCase(expr *SearchedCase) (interface{}, error)
	VisitSimpleCase(expr *SimpleCase) (interface{}, error)

	/*
	   Collections. (A collection is an ordered group of elements).
	   Refer to the N1QL specs for the list of supported
	   collections.
	*/
	VisitAny(expr *Any) (interface{}, error)
	VisitArray(expr *Array) (interface{}, error)
	VisitEvery(expr *Every) (interface{}, error)
	VisitExists(expr *Exists) (interface{}, error)
	VisitFirst(expr *First) (interface{}, error)
	VisitIn(expr *In) (interface{}, error)
	VisitWithin(expr *Within) (interface{}, error)

	/*
	   Comparison terms hwlp compare two or more expressions.
	   Refer to the N1QL specs for the list of supported
	   comparison terms.
	*/
	VisitBetween(expr *Between) (interface{}, error)
	VisitEq(expr *Eq) (interface{}, error)
	VisitLE(expr *LE) (interface{}, error)
	VisitLike(expr *Like) (interface{}, error)
	VisitLT(expr *LT) (interface{}, error)
	VisitIsMissing(expr *IsMissing) (interface{}, error)
	VisitIsNotMissing(expr *IsNotMissing) (interface{}, error)
	VisitIsNotNull(expr *IsNotNull) (interface{}, error)
	VisitIsNotValued(expr *IsNotValued) (interface{}, error)
	VisitIsNull(expr *IsNull) (interface{}, error)
	VisitIsValued(expr *IsValued) (interface{}, error)

	/*
	   Concat. Both expressions need to be strings, else
	   returns null.
	*/
	VisitConcat(expr *Concat) (interface{}, error)

	/*
	   Visit a Constant expression.
	*/
	VisitConstant(expr *Constant) (interface{}, error)

	/*
	   Identifier. They can be of two types, escaped and unescaped.
	   Refer to the N1QL specs.
	*/
	VisitIdentifier(expr *Identifier) (interface{}, error)

	/*
	   Construction. As per the N1ql specs, objects and arrays can
	   be constructed with arbitrary structure, nesting, and
	   embedded expressions.
	*/
	VisitArrayConstruct(expr *ArrayConstruct) (interface{}, error)
	VisitObjectConstruct(expr *ObjectConstruct) (interface{}, error)

	/*
	   Logical Terms use boolean logic. Standard operators.
	*/
	VisitAnd(expr *And) (interface{}, error)
	VisitNot(expr *Not) (interface{}, error)
	VisitOr(expr *Or) (interface{}, error)

	/*
	   Navigation. Used to navigate through objects and
	   slices(arrays).
	*/
	VisitElement(expr *Element) (interface{}, error)
	VisitField(expr *Field) (interface{}, error)
	VisitFieldName(expr *FieldName) (interface{}, error)
	VisitSlice(expr *Slice) (interface{}, error)

	/*
	   Function. Refer to N1QL for a list of supported functions.
	*/
	VisitFunction(expr Function) (interface{}, error)

	/*
	   Subquery. Returns the subquery string.
	*/
	VisitSubquery(expr Subquery) (interface{}, error)

	/*
	   Parameters. There are 2 types, named and positional.
	   It allows passing of parameters into the query using
	   position or name.
	*/
	VisitNamedParameter(expr NamedParameter) (interface{}, error)
	VisitPositionalParameter(expr PositionalParameter) (interface{}, error)
}
