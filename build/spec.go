/*
 * Copyright 1999-2019 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"
	"os"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"

	"github.com/chaosblade-io/chaosblade-exec-cri/exec"
)

// main creates a yaml file of the experiments in the project
func main() {
	if len(os.Args) < 2 {
		log.Panicln("less yaml file path")
	}
	if len(os.Args) == 3 {
		exec.JvmSpecFileForYaml = os.Args[2]
	}
	err := util.CreateYamlFile(getModels(), os.Args[1])
	if err != nil {
		log.Panicf("create yaml file error, %v", err)
	}
}

// getModels returns the supported experiment specs
func getModels() *spec.Models {
	models := make([]*spec.Models, 0)
	dockerModelSpec := exec.NewCriExpModelSpec()
	for _, modelSpec := range dockerModelSpec.ExpModels() {
		model := util.ConvertSpecToModels(modelSpec, spec.ExpPrepareModel{}, dockerModelSpec.Scope())
		models = append(models, model)
	}
	return util.MergeModels(models...)
}
