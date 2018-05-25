package main
 
 import (
         "bytes"
         "encoding/json"
         "fmt"
         "strconv"
 
         "github.com/hyperledger/fabric/core/chaincode/shim"
         sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 type SmartContract struct {
 }
 
 
 type Doc struct {
         Docid string `json:"docid"`
         Timestamp string `json:"timestamp"`
         Dochash  string `json:"dochash"`
         Owner  string `json:"owner"`
 }

 type User struct {
	 UserHash   string `json:"userhash"`
	 MetaHash    string `json:"metahash"`
	 UserClass  string `json:"userclass"`
	 IsOrg      bool   `json:"isorg"`
 }
 
 
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
         return shim.Success(nil)
 }

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
         // Retrieve the requested Smart Contract function and arguments
         function, args := APIstub.GetFunctionAndParameters()
         // Route to the appropriate handler function to interact with the ledger
         if function == "queryDoc" {
                 return s.queryDoc(APIstub, args)
         } else if function == "initLedger" {
                 return s.initLedger(APIstub)
         } else if function == "recordDoc" {
                 return s.recordDoc(APIstub, args)
         } else if function == "queryAllDoc" {
                 return s.queryAllDoc(APIstub)
         } else if function == "changeDocOwner" {
                 return s.changeDocOwner(APIstub, args)
         } else if function == "createUser" {
                 return s.createUser(APIstub, args)
         } else if function == "queryAllUser" {
                 return s.queryAllUser(APIstub)
         } else if function == "getUser" {
                 return s.getUser(APIstub, args)
         } 
 
         return shim.Error("Invalid Smart Contract function name.")
 }

func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 1 {
                 return shim.Error("Incorrect number of arguments. Expecting 1")
         }
 
         userAsBytes, _ := APIstub.GetState(args[0])
         if userAsBytes == nil {
                 return shim.Error("Could not find User")
         }
         return shim.Success(userAsBytes)
 } 

func (s *SmartContract) queryDoc(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 1 {
                 return shim.Error("Incorrect number of arguments. Expecting 1")
         }
 
         docAsBytes, _ := APIstub.GetState(args[0])
         if docAsBytes == nil {
                 return shim.Error("Could not locate document")
         }
         return shim.Success(docAsBytes)
 }

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
         user := []User{
                 User{UserHash: "aaaaaaaaaaaa", MetaHash: "xxxxxxxxxxxxxx", UserClass: "ashish", IsOrg: "Yes"},
                 User{UserHash: "bbbbbbbbbbbb", MetaHash: "yyyyyyyyyyyyyy", UserClass: "vasa", IsOrg: "No"},
                 User{UserHash: "cccccccccccc", MetaHash: "zzzzzzzzzzzzzz", UserClass: "hamza", IsOrg: "Yes"},

         }
 
         i := 0
         for i < len(user) {
                 fmt.Println("i is ", i)
                 userAsBytes, _ := json.Marshal(user[i])
                 APIstub.PutState(strconv.Itoa(i+1), userAsBytes)
                 fmt.Println("Added", user[i])
                 i = i + 1
         }
 
         return shim.Success(nil)
 }

/*func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
         doc := []Doc{
                 Doc{Docid: "0001", Timestamp: "1504054225", Dochash: "agsbsgwwgegwgges", Owner: "Ashish"},
                 Doc{Docid: "0002", Timestamp: "1504054227", Dochash: "nedbweiehwjwjeje", Owner: "Vasa"},
                 Doc{Docid: "0003", Timestamp: "1504054230", Dochash: "jjehdhh2iwejwjwj", Owner: "Hamza"},

         }
 
         i := 0
         for i < len(doc) {
                 fmt.Println("i is ", i)
                 docAsBytes, _ := json.Marshal(doc[i])
                 APIstub.PutState(strconv.Itoa(i+1), docAsBytes)
                 fmt.Println("Added", doc[i])
                 i = i + 1
         }
 
         return shim.Success(nil)
 }*/

func (s *SmartContract) createUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 5 {
                 return shim.Error("Incorrect number of arguments. Expecting 5")
         }
 
         var user = User{ UserHash: args[1], MetaHash: args[2], UserClass: args[3], IsOrg: args[4] }
 
         userAsBytes, _ := json.Marshal(user)
         err := APIstub.PutState(args[0], userAsBytes)
         if err != nil {
                 return shim.Error(fmt.Sprintf("Failed to create user: %s", args[0]))
         }
 
         return shim.Success(nil)
 }


func (s *SmartContract) recordDoc(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 5 {
                 return shim.Error("Incorrect number of arguments. Expecting 5")
         }
 
         var doc = Doc{ Docid: args[1], Timestamp: args[2], Dochash: args[3], Owner: args[4] }
 
         docAsBytes, _ := json.Marshal(doc)
         err := APIstub.PutState(args[0], docAsBytes)
         if err != nil {
                 return shim.Error(fmt.Sprintf("Failed to record doc: %s", args[0]))
         }
 
         return shim.Success(nil)
 }

func (s *SmartContract) queryAllDoc(APIstub shim.ChaincodeStubInterface) sc.Response {
 
         startKey := "0"
         endKey := "999999"
 
         resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
         if err != nil {
                 return shim.Error(err.Error())
         }
         defer resultsIterator.Close()
 
         var buffer bytes.Buffer
         buffer.WriteString("[")
 
         bArrayMemberAlreadyWritten := false
         for resultsIterator.HasNext() {
                 queryResponse, err := resultsIterator.Next()
                 if err != nil {
                         return shim.Error(err.Error())
                 }

                 if bArrayMemberAlreadyWritten == true {
                         buffer.WriteString(",")
                 }
                 buffer.WriteString("{\"Key\":")
                 buffer.WriteString("\"")
                 buffer.WriteString(queryResponse.Key)
                 buffer.WriteString("\"")
 
                 buffer.WriteString(", \"Record\":")

                 buffer.WriteString(string(queryResponse.Value))
                 buffer.WriteString("}")
                 bArrayMemberAlreadyWritten = true
         }
         buffer.WriteString("]")
        fmt.Printf("- queryAllDoc:\n%s\n", buffer.String())
 
         return shim.Success(buffer.Bytes())
 }

 func (s *SmartContract) changeDocOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 2 {
                 return shim.Error("Incorrect number of arguments. Expecting 2")
         }
 
         docAsBytes, _ := APIstub.GetState(args[0])
         if docAsBytes == nil {
                 return shim.Error("Could not locate doc")
         }
         doc := Doc{}
 
         json.Unmarshal(docAsBytes, &doc)

         doc.Owner = args[1]
 
         docAsBytes, _ = json.Marshal(doc)
         err := APIstub.PutState(args[0], docAsBytes)
         if err != nil {
                 return shim.Error(fmt.Sprintf("Failed to change doc Owner: %s", args[0]))
         }
 
         return shim.Success(nil)
 }

 func main() {
 
         // Create a new Smart Contract
         err := shim.Start(new(SmartContract))
         if err != nil {
                 fmt.Printf("Error creating new Smart Contract: %s", err)
         }
 }


