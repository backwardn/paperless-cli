// Copyright © 2019 Steve Garf <stgarf@gmail.com>
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

package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var documentsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"li", "l"},
	Short:   "List documents from Paperless",
	Run: func(cmd *cobra.Command, args []string) {
		docs, err := PaperInst.GetDocuments()
		if err != nil {
			log.Errorf("%s", err)
		}
		fmt.Printf("%v results found:\n", len(docs))
		for i, doc := range docs {
			fmt.Printf("%d. %v - %v %+q\n", i+1, doc.Correspondent, doc.Title, doc.Tags)
		}
	},
}

func init() {
	documentsCmd.AddCommand(documentsListCmd)
}
