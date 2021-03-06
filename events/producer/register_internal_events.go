/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

//All changes for handling internal events should be in this file
//Step 1 - add the event type name to the const
//Step 2 - add a case statement to getMessageType
//Step 3 - add an AddEventType call to addInternalEventTypes

package producer

import (
	pb "github.com/openblockchain/obc-peer/protos"
)

//----Event Types -----
const (
	RegisterType = "register"
	BlockType    = "block"
)

func getMessageType(e *pb.OpenchainEvent) string {
	switch e.Event.(type) {
	case *pb.OpenchainEvent_Register:
		return "register"
	case *pb.OpenchainEvent_Block:
		return "block"
	case *pb.OpenchainEvent_Generic:
		return "generic"
	default:
		return ""
	}
}

//should be called at init time to register supported internal events
func addInternalEventTypes() {
	AddEventType(BlockType)
	AddEventType(RegisterType)
}
