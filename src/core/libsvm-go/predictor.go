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
** Description: Prediciton related APIs
** @author: Ed Walker
 */
package libSvm

/**
*  This function gives decision values on a test vector x given a
   model, and return the predicted label (classification) or
   the function value (regression).

   For a classification model with nrClass classes, this function
   gives nrClass*(nrClass-1)/2 decision values in the slice
   decisionValues.  The order is label[0] vs. label[1], ...,
   label[0] vs. label[nr_class-1], label[1] vs. label[2], ...,
   label[nrClass-2] vs. label[nrClass-1]. The returned returnValue is
   the predicted class for x. Note that when nrClass = 1, this
   function does not give any decision value.

   For a regression model, decisionValues[0] and the returned returnValue are
   both the function value of x calculated using the model. For a
   one-class model, decisionValues[0] is the decision value of x, while
   the returned returnValue is +1/-1.

*/
func (model Model) PredictValues(x map[int]float64) (returnValue float64, decisionValues []float64) {
	returnValue = 0

	px := MapToSnode(x)

	switch model.Param.SvmType {
	case ONE_CLASS, EPSILON_SVR, NU_SVR:
		var svCoef []float64 = model.SvCoef[0]

		var sum float64 = 0
		for i := 0; i < model.L; i++ {
			var idx_y int = model.SV[i]
			py := model.SvSpace[idx_y:]
			sum += svCoef[i] * computeKernelValue(px, py, model.Param)
		}
		sum -= model.Rho[0]

		decisionValues = append(decisionValues, sum)

		if model.Param.SvmType == ONE_CLASS {
			if sum > 0 {
				returnValue = 1 // return
			} else {
				returnValue = -1 // return
			}
			return // returnValue, decisionValues
		} else {
			returnValue = sum
			return // returnValue, decisionValues
		}

	case C_SVC, NU_SVC:
		var nrClass int = model.Nrclass
		var l int = model.L

		kvalue := make([]float64, l)
		for i := 0; i < l; i++ {
			var idx_y int = model.SV[i]
			py := model.SvSpace[idx_y:]
			kvalue[i] = computeKernelValue(px, py, model.Param)
		}

		start := make([]int, nrClass)
		start[0] = 0
		for i := 1; i < nrClass; i++ {
			start[i] = start[i-1] + model.NSV[i-1]
		}

		vote := make([]int, nrClass)
		for i := 0; i < nrClass; i++ {
			vote[i] = 0
		}

		var p int = 0
		for i := 0; i < nrClass; i++ {
			for j := i + 1; j < nrClass; j++ {
				var sum float64 = 0

				var si int = start[i]
				var sj int = start[j]

				var ci int = model.NSV[i]
				var cj int = model.NSV[j]

				coef1 := model.SvCoef[j-1]
				coef2 := model.SvCoef[i]
				for k := 0; k < ci; k++ {
					sum += coef1[si+k] * kvalue[si+k]
				}
				for k := 0; k < cj; k++ {
					sum += coef2[sj+k] * kvalue[sj+k]
				}
				sum -= model.Rho[p]
				decisionValues = append(decisionValues, sum)
				if sum > 0 {
					vote[i]++
				} else {
					vote[j]++
				}
				p++
			}
		}

		var maxIdx int = 0
		for i := 1; i < nrClass; i++ {
			if vote[i] > vote[maxIdx] {
				maxIdx = i
			}
		}

		returnValue = float64(model.Label[maxIdx])
		return // returnValue, decisionValues
	}

	return
}

/**
* This function does classification or regression on a test vector x
   given a model.

   For a classification model, the predicted class for x is returned.
   For a regression model, the function value of x calculated using
   the model is returned. For an one-class model, +1 or -1 is
   returned.

*/
func (model Model) Predict(x map[int]float64) float64 {

	predict, _ := model.PredictValues(x)

	return predict
}
