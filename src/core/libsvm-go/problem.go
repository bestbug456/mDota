/*
** Copyright 2014 Edward Walker
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http ://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
**
** Description: Describes problem, i.e. label/vector set
** @author: Ed Walker
 */
package libSvm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Snode struct {
	Index int     // dimension (-1 indicates end of SV)
	Value float64 // coeff
}

type Problem struct {
	L      int       // #SVs
	Y      []float64 // labels
	X      []int     // starting indices in xSpace defining SVs
	XSpace []Snode   // SV coeffs
	I      int       // counter for iterator
}

func NewProblem(file string, param *Parameter) (*Problem, error) {
	prob := &Problem{L: 0, I: 0}
	err := prob.Read(file, param)
	return prob, err
}

func (problem *Problem) Read(file string, param *Parameter) error { // reads the problem from the specified file
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Fail to open file %s\n", file)
	}

	defer f.Close() // close f on method return

	problem.Y = nil
	problem.X = nil
	problem.XSpace = nil

	reader := bufio.NewReader(f)
	var max_idx int = 0
	var l int = 0

	for {
		line, err := readline(reader)
		if err != nil {
			break
		}
		problem.X = append(problem.X, len(problem.XSpace))

		lineSansComments := strings.Split(line, "#") // remove any comments

		tokens := strings.Fields(lineSansComments[0]) // get all the word tokens (seperated by white spaces)
		if label, err := strconv.ParseFloat(tokens[0], 64); err == nil {
			problem.Y = append(problem.Y, label)
		} else {
			return fmt.Errorf("Fail to parse label\n")
		}

		space := tokens[1:]
		for _, w := range space {
			if len(w) > 0 {
				node := strings.Split(w, ":")
				if len(node) > 1 {
					var index int
					var value float64
					if index, err = strconv.Atoi(node[0]); err != nil {
						return fmt.Errorf("Fail to parse index from token %v\n", w)
					}
					if value, err = strconv.ParseFloat(node[1], 64); err != nil {
						return fmt.Errorf("Fail to parse value from token %v\n", w)
					}
					problem.XSpace = append(problem.XSpace, Snode{Index: index, Value: value})
					if index > max_idx {
						max_idx = index
					}

				}
			}
		}

		problem.XSpace = append(problem.XSpace, Snode{Index: -1})
		l++
	}
	problem.L = l

	if param.Gamma == 0 && max_idx > 0 {
		param.Gamma = 1.0 / float64(max_idx)
	}

	return nil
}

/**
 * Initialize the start of iterating through the labels and vectors in the problem set
 */
func (problem *Problem) Begin() {
	problem.I = 0
}

/**
 * Finished iterating through all the labels and vectors in the problem set
 */
func (problem *Problem) Done() bool {
	if problem.I >= problem.L {
		return true
	}
	return false
}

/**
 * Move to the next label and vector in the problem set
 */
func (problem *Problem) Next() {
	problem.I++
	return
}

/**
 * Return one label and vector from the problem set
 * @return y label
 * @return x vector (map of dimension/value)
 */
func (problem *Problem) GetLine() (y float64, x map[int]float64) {
	y = problem.Y[problem.I]
	idx := problem.X[problem.I]
	x = SnodeToMap(problem.XSpace[idx:])
	return // y, x
}

/**
 * Returns number of label and vectors in the problem set
 * @return problem set size
 */
func (problem *Problem) ProblemSize() int {
	return problem.L
}
