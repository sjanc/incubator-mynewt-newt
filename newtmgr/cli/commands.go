/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"mynewt.apache.org/newt/newtmgr/nmutil"
	"mynewt.apache.org/newt/util"
)

var ConnProfileName string
var NewtmgrLogLevel log.Level

func Commands() *cobra.Command {
	logLevelStr := ""
	nmCmd := &cobra.Command{
		Use:   "newtmgr",
		Short: "Newtmgr helps you manage remote instances of the Mynewt OS.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			NewtmgrLogLevel, err := log.ParseLevel(logLevelStr)
			err = util.Init(NewtmgrLogLevel, "", util.VERBOSITY_DEFAULT)
			if err != nil {
				nmUsage(nil, err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	nmCmd.PersistentFlags().StringVarP(&ConnProfileName, "conn", "c", "",
		"connection profile to use.")

	nmCmd.PersistentFlags().StringVarP(&logLevelStr, "loglevel", "l", "info",
		"log level to use (default INFO.)")

	nmCmd.PersistentFlags().BoolVarP(&nmutil.TraceLogEnabled, "trace", "t",
		false, "print all bytes transmitted and received")

	nmCmd.AddCommand(connProfileCmd())
	nmCmd.AddCommand(echoCmd())
	nmCmd.AddCommand(imageCmd())
	nmCmd.AddCommand(statsCmd())
	nmCmd.AddCommand(taskStatsCmd())
	nmCmd.AddCommand(mempoolStatsCmd())
	nmCmd.AddCommand(configCmd())
	nmCmd.AddCommand(logsCmd())
	nmCmd.AddCommand(dTimeCmd())
	nmCmd.AddCommand(resetCmd())
	nmCmd.AddCommand(crashCmd())
	nmCmd.AddCommand(runtestCmd())

	return nmCmd
}
