// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START bigquery_generated_bigquery_Dataset_Create]

package main

import (
	"context"

	"cloud.google.com/go/bigquery"
)

func main() {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	ds := client.Dataset("my_dataset")
	if err := ds.Create(ctx, &bigquery.DatasetMetadata{Location: "EU"}); err != nil {
		// TODO: Handle error.
	}
}

// [END bigquery_generated_bigquery_Dataset_Create]
