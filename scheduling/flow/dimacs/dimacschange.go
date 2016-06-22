// Copyright 2016 The ksched Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The dimacs_change.h class equivalent
// This is meant to be inherited by its children clases such as dimacs_change_arc/node.h
// So we implement this as an interface to get the effect of polymorphism in the child classes

// DIMACS documentation:
// DIMACS is a standard format of representing a graph used
// in order to facitilate solving problems such as max flow - min cost
// See http://lpsolve.sourceforge.net/5.5/DIMACS_mcf.htm
// Essentially it represent the graph in the following format:
//
// 1. Comment Lines: Comment lines give human-readable information about the file and are ignored by programs
//
// 2. Problem Line: p min NODES ARCS
// Where 'p' indicates it is a problem line, min specifies the problem type as min cost
// NODES specifies the number of nodes and ARCS specifies the number of arcs in the graph
//
// 3. Node Descriptors: n ID FLOW
// Where 'n' indicates this a node descriptor line, ID is the node ID between 1 and n
// and FLOW indicates the supply of flow at this node
//
// 4. Arc Descriptors: a SRC DST LOW CAP COST
// Where 'a' indicates this is an Arc descriptor line, SRC and DST specify the source and destination
// vertex IDs between 1 and n, LOW specifies the lower bound on the flow
// and CAP and COST specify the capacity and cost of the arc respectively

package dimacs

type Change interface {
	// Every struct implementing this interface must have this field
	// comment string

	// Getter Setter for comment
	Comment() string
	SetComment() string
	// This had an implementation in the parent class
	// So it should be implemented as is given in the original file
	GenerateChangeDescription() string
	GenerateChange() string
}
