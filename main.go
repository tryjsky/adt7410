/*
Copyright 2023 Y., Ryota <tryjsky@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
    "fmt"
    "log"

    "github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
)

func main() {
    // Create new connection to i2c_1-bus on 1 line with address 0x48.
    i2c_1, err := i2c.NewI2C(0x48, 1)
    if err != nil {
        log.Fatal(err)
    }
    // Free I2C connection on exit
    defer i2c_1.Close()

	// Decrease i2c verbosity
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)

	// Read word from I2C device
    w, err := i2c_1.ReadRegS16BE(0)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%v\n", float32(w) / 8 * 0.0625)
}