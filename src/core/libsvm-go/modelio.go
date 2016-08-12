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
** Description: Input/output routines for the Support Vector Machine model
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

func (model *Model) Dump(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("Fail to open file %s\n", file)
	}

	defer f.Close() // close f on method return

	var output []string

	//svm_type_string := [5]string{"c_svc", "nu_svc", "one_class", "epsilon_svr", "nu_svr"}
	output = append(output, fmt.Sprintf("svm_type %s\n", svm_type_string[model.Param.SvmType]))

	output = append(output, fmt.Sprintf("kernel_type %s\n", kernel_type_string[model.Param.KernelType]))

	if model.Param.KernelType == POLY {
		output = append(output, fmt.Sprintf("degree %d\n", model.Param.Degree))
	}

	if model.Param.KernelType == POLY || model.Param.KernelType == RBF || model.Param.KernelType == SIGMOID {
		output = append(output, fmt.Sprintf("gamma %.6g\n", model.Param.Gamma))
	}

	if model.Param.KernelType == POLY || model.Param.KernelType == SIGMOID {
		output = append(output, fmt.Sprintf("coef0 %.6g\n", model.Param.Coef0))
	}

	var nrClass int = model.Nrclass
	output = append(output, fmt.Sprintf("nr_class %d\n", nrClass))

	var l int = model.L
	output = append(output, fmt.Sprintf("total_sv %d\n", l))

	output = append(output, "rho")
	total_models := nrClass * (nrClass - 1) / 2
	for i := 0; i < total_models; i++ {
		output = append(output, fmt.Sprintf(" %.6g", model.Rho[i]))
	}
	output = append(output, "\n")

	if len(model.Label) > 0 {
		output = append(output, "label")
		for i := 0; i < nrClass; i++ {
			output = append(output, fmt.Sprintf(" %d", model.Label[i]))
		}
		output = append(output, "\n")
	}

	if len(model.ProbA) > 0 {
		output = append(output, "probA")
		for i := 0; i < total_models; i++ {
			output = append(output, fmt.Sprintf(" %.8g", model.ProbA[i]))
		}
		output = append(output, "\n")
	}

	if len(model.ProbB) > 0 {
		output = append(output, "probB")
		for i := 0; i < total_models; i++ {
			output = append(output, fmt.Sprintf(" %.8g", model.ProbB[i]))
		}
		output = append(output, "\n")
	}

	if len(model.NSV) > 0 {
		output = append(output, "nr_sv")
		for i := 0; i < nrClass; i++ {
			output = append(output, fmt.Sprintf(" %d", model.NSV[i]))
		}
		output = append(output, "\n")
	}

	output = append(output, "SV\n")

	for i := 0; i < l; i++ {
		for j := 0; j < nrClass-1; j++ {
			output = append(output, fmt.Sprintf("%.16g ", model.SvCoef[j][i]))
		}

		i_idx := model.SV[i]
		if model.Param.KernelType == PRECOMPUTED {
			output = append(output, fmt.Sprintf("0:%d ", model.SvSpace[i_idx]))
		} else {
			for model.SvSpace[i_idx].Index != -1 {
				index := model.SvSpace[i_idx].Index
				value := model.SvSpace[i_idx].Value
				output = append(output, fmt.Sprintf("%d:%.8g ", index, value))
				i_idx++
			}
			output = append(output, "\n")
		}
	}

	f.WriteString(strings.Join(output, ""))

	return nil
}

func (model *Model) readHeader(reader *bufio.Reader) error {

	for {
		var i int = 0
		var err error
		var line string

		line, err = readline(reader)
		if err != nil { // We should not encounter an EOF.  If we do, it is an error.
			return err
		}

		tokens := strings.Fields(line)

		switch tokens[0] {
		case "svm_type":

			for i = 0; i < len(svm_type_string); i++ {
				if svm_type_string[i] == tokens[1] {
					model.Param.SvmType = i
					break
				}
			}

			if i == len(svm_type_string) {
				return fmt.Errorf("fail to parse svm model %s\n", tokens[1])
			}

		case "kernel_type":

			for i = 0; i < len(kernel_type_string); i++ {
				if kernel_type_string[i] == tokens[1] {
					model.Param.KernelType = i
					break
				}
			}

			if i == len(kernel_type_string) {
				return fmt.Errorf("fail to parse kernel type %s\n", tokens[1])
			}

		case "degree":

			if model.Param.Degree, err = strconv.Atoi(tokens[1]); err != nil {
				return err
			}

		case "gamma":

			if model.Param.Gamma, err = strconv.ParseFloat(tokens[1], 64); err != nil {
				return err
			}

		case "coef0":

			if model.Param.Coef0, err = strconv.ParseFloat(tokens[1], 64); err != nil {
				return err
			}

		case "nr_class":

			if model.Nrclass, err = strconv.Atoi(tokens[1]); err != nil {
				return err
			}

		case "total_sv":

			if model.L, err = strconv.Atoi(tokens[1]); err != nil {
				return err
			}

		case "rho":

			total_class_comparisons := model.Nrclass * (model.Nrclass - 1) / 2
			if total_class_comparisons != len(tokens)-1 {
				return fmt.Errorf("Number of rhos %d does not mactch the required number %d\n", len(tokens)-1, total_class_comparisons)
			}

			model.Rho = make([]float64, total_class_comparisons)
			for i = 0; i < total_class_comparisons; i++ {
				if model.Rho[i], err = strconv.ParseFloat(tokens[i+1], 64); err != nil {
					return err
				}
			}

		case "label":

			if model.Nrclass != len(tokens)-1 {
				return fmt.Errorf("Number of labels %d does not appear in the file\n", model.Nrclass)
			}

			model.Label = make([]int, model.Nrclass)
			for i = 0; i < model.Nrclass; i++ {
				if model.Label[i], err = strconv.Atoi(tokens[i+1]); err != nil {
					return err
				}
			}

		case "probA":

			total_class_comparisons := model.Nrclass * (model.Nrclass - 1) / 2
			if total_class_comparisons != len(tokens)-1 {
				return fmt.Errorf("Number of probA %d does not mactch the required number %d\n", len(tokens)-1, total_class_comparisons)
			}

			model.ProbA = make([]float64, total_class_comparisons)
			for i = 0; i < total_class_comparisons; i++ {
				if model.ProbA[i], err = strconv.ParseFloat(tokens[i+1], 64); err != nil {
					return err
				}
			}

		case "probB":

			total_class_comparisons := model.Nrclass * (model.Nrclass - 1) / 2
			if total_class_comparisons != len(tokens)-1 {
				return fmt.Errorf("Number of probB %d does not mactch the required number %d\n", len(tokens)-1, total_class_comparisons)
			}

			model.ProbB = make([]float64, total_class_comparisons)
			for i = 0; i < total_class_comparisons; i++ {
				if model.ProbB[i], err = strconv.ParseFloat(tokens[i+1], 64); err != nil {
					return err
				}
			}

		case "nr_sv":

			if model.Nrclass != len(tokens)-1 {
				return fmt.Errorf("Number of nSV %d does not appear in the file %v\n", model.Nrclass, tokens)
			}

			model.NSV = make([]int, model.Nrclass)
			for i = 0; i < model.Nrclass; i++ {
				if model.NSV[i], err = strconv.Atoi(tokens[i+1]); err != nil {
					return err
				}
			}

		case "SV":
			return nil // done reading the header!
		default:
			return fmt.Errorf("unknown text in model file: [%s]\n", tokens[0])

		}
	}

	return fmt.Errorf("Fail to completely read header")
}

func (model *Model) ReadModel(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Fail to open file %s\n", file)
	}

	defer f.Close() // close f on method return

	reader := bufio.NewReader(f)

	if err := model.readHeader(reader); err != nil {
		return err
	}

	var l int = model.L           // read l from header
	var m int = model.Nrclass - 1 // read nrClass from header
	model.SvCoef = make([][]float64, m)
	for i := 0; i < m; i++ {
		model.SvCoef[i] = make([]float64, l)
	}

	model.SV = make([]int, l)
	var i int = 0
	for {
		line, err := readline(reader) // read a line
		if err != nil {
			break
		}

		tokens := strings.Fields(line) // get all the word tokens (seperated by white spaces)
		if len(tokens) < 2 {           // there should be at least 2 fields -- label + SV
			continue
		}
		if i >= l {
			return fmt.Errorf("Error in reading support vectors.  i=%d and l=%d\n", i, l)
		}

		model.SV[i] = len(model.SvSpace) // starting index into svSpace for this SV

		var k int = 0
		for _, token := range tokens {
			if k < m {
				model.SvCoef[k][i], err = strconv.ParseFloat(token, 64)
				k++
			} else {
				node := strings.Split(token, ":")
				if len(node) < 2 {
					return fmt.Errorf("Fail to parse svSpace from token %v\n", token)
				}
				var index int
				var value float64
				if index, err = strconv.Atoi(node[0]); err != nil {
					return fmt.Errorf("Fail to parse index from token %v\n", token)
				}
				if value, err = strconv.ParseFloat(node[1], 64); err != nil {
					return fmt.Errorf("Fail to parse value from token %v\n", token)
				}
				model.SvSpace = append(model.SvSpace, Snode{Index: index, Value: value})
			}
		}
		model.SvSpace = append(model.SvSpace, Snode{Index: -1})
		i++
	}

	return nil
}
