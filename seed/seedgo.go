/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Seed struct {

  LotNumber string `json:"lotNumber"`
  owner   string `json:"owner"`
  crop    string `json:"crop"`
  variety string `json:"variety"`
  SourceTagNo string `json:"sourceTagNo"`
  SourceClass string `json:"sourceClass"`
  DestinationClass string `json:"destinationClass"`
  SourceQuantity string `json:"SourceQuantity"`
  SourceDateOfIssue string `json:"sourceDateOfIssue"`
  growerName string `json:"growerName"`
  spaName string `json:"spaName"`
  dateOfIssue string `json:"dateOfIssue"`
  sourceStorehouse  string `json:"sourceStorehouse"`
  destiStorehouse string `json:"destiStorehouse"`
  sgName string `json:"sgName"`
  sgId  string `json:"sgId"`
  FinYear string `json:"finYear"`
  Season  string `json:"Season"`
  landRecordsKhatano  string `json:"landRecordsKhatano"`
  landRecordsPlotno string `json:"landRecordsPlotno"`
  landRecordsArea string `json:"landRecordsArea"`
  cropRegCode string `json:"cropRegCode"`
  SppName string `json:"sppName"`
  SppID string `json:"sppID"`
  TotalQuantityProcessed string `json:"TotalQuantityProcessed"`
  processingDate  string `json:"processingDate"`
  verificationDate  string `json:"verificationDate"`
  SampleSecretCode string `json:"sampleSecretCode"`
  SamplePassed  string `json:"samplePassed"`
  SampleTestDate  string `json:"SampleTestDate"`
  CertificateNo string `json:"certificateNo"`
  CerticateDate string `json:"certificateDate"`
  TagSeris  string `json:"tagSeris"`
  TagIssuedRangeFrom  string `json:"tagIssuedRangeTo"`
  TagIssuedRangeTo  string `json:"tagIssuedRangeTo"`
  NoOfTagsIssued  string `json:"noOfTagsIssued"`
  CertificateValidityInMonth  string `json:"certificateValidityInMonth"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Seed
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	Seeds := Seed{
      LotNumber : "APR19-33-028-117(2)",
      owner : "OSSC",
      crop : "Paddy",
      variety : "LALATA",
      SourceTagNo : "F33162200068",
      SourceClass : "Foundation-1",
      DestinationClass : "Certified-1",
      SourceQuantity : "20KG",
      SourceDateOfIssue : "24/02/2020",
      growerName : "Ram Babu",
      spaName : "OSSC",
      dateOfIssue : "Paddy",
      sourceStorehouse : "KHURDHA, PIN 753011",
      destiStorehouse : "PURI, PIN 7534511",
      sgName : "Aanad singh",
      sgId : "AA-2345",
      FinYear : "2020",
      Season : "monsoon",
      landRecordsKhatano : "00031",
      landRecordsPlotno : "1306",
      landRecordsArea : "1306",
      cropRegCode : "1.535",
      SppName : "Z20-7-9",
      SppID : "33-001",
      TotalQuantityProcessed : "55.55",
      processingDate : "10/10/2020:12:30",
      verificationDate : "15/10/2020:10:30".
      SampleSecretCode : "seechain01",
      SamplePassed : "Yes",
      SampleTestDate : "25/10/2020:12:30",
      CertificateNo : "123456789",
      CerticateDate :"30/10/2020:12:30",
      TagSeris : "C7577",
      TagIssuedRangeFrom :"C7577-4499115",
      TagIssuedRangeTo :  "C7577-4499534",
      NoOfTagsIssued :"420",
      CertificateValidityInMonth : "9"

	}

	for i, seed := range Seeds {
		seedAsBytes, _ := json.Marshal(seed)
		err := ctx.GetStub().PutState("SEED"+strconv.Itoa(i), seedAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateCar adds a new car to the world state with given details
func (s *SmartContract)
CreateSeed(ctx contractapi.TransactionContextInterface, lotNumber string,owner  string,crop string,variety string,sourceTagNo string,sourceClass string,destinationClass string,sourceQuantity string,sourceDateOfIssue string,growerName string,spaName string,dateOfIssue string,sourceStorehouse string,destiStorehouse string,sgName string,sgId string,finYear string,season string,landRecordsKhatano string,landRecordsPlotno string,landRecordsArea string,cropRegCode string,sppName string, sppID string) error {
	seed := Seed{
    LotNumber : lotNumber ,
    owner : owner,
    crop : crop,
    variety : variety,
    SourceTagNo : sourceTagNo,
    SourceClass : sourceClass,
    DestinationClass : destinationClass,
    SourceQuantity : sourceQuantity,
    SourceDateOfIssue : sourceDateOfIssue,
    growerName : growerName,
    spaName : spaName,
    dateOfIssue : dateOfIssue,
    sourceStorehouse : sourceStorehouse,
    destiStorehouse :  destiStorehouse,
    sgName : sgName,
    sgId : sgId,
    FinYear : finYear,
    Season : season,
    landRecordsKhatano : landRecordsKhatano,
    landRecordsPlotno : landRecordsPlotno,
    landRecordsArea : landRecordsArea,
    cropRegCode : cropRegCode,
    SppName : sppName,
    SppID : SppID,

}

	seedAsBytes, _ := json.Marshal(seed)

	return ctx.GetStub().PutState(lotNumber, seedAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QuerySeedByLot(ctx contractapi.TransactionContextInterface, lotNumber string) (*Seed, error) {
	seedAsBytes, err := ctx.GetStub().GetState(lotNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if seedAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", lotNumber)
	}

	seed := new(Seed)
	_ = json.Unmarshal(seedAsBytes, seed)

	return seed, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllSeeds(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		seed := new(Seed)
		_ = json.Unmarshal(queryResponse.Value, seed)

		queryResult := QueryResult{Key: queryResponse.Key, Record: seed}
		results = append(results, queryResult)
	}

	return results, nil
}

// CertificateUpdate updates certains fields of seed with given lotNumber in world state
func (s *SmartContract) CertificateUpdate(ctx contractapi.TransactionContextInterface, lotNumber string, newOwner string) error {
	seed, err := s.QueryCar(ctx, lotNumber)

	if err != nil {
		return err
	}

  seed.TotalQuantityProcessed = totalQuantityProcessed
  seed.processingDate =processingDate
  seed.verificationDate =verificationDate
  seed.SampleSecretCode =SampleSecretCode
  seed.SamplePassed = SamplePassed
  seed.SampleTestDate=SampleTestDate
  seed.CertificateNo = CertificateNo
  seed.CerticateDate = CerticateDate
  seed.TagSeris = TagSeris
  seed.TagIssuedRangeFrom =TagIssuedRangeFrom
  seed.TagIssuedRangeTo = TagIssuedRangeTo
  seed.NoOfTagsIssued = NoOfTagsIssued
  seed.CertificateValidityInMonth = CertificateValidityInMonth

	seedAsBytes, _ := json.Marshal(seed)

	return ctx.GetStub().PutState(lotNumber, seedAsBytes)
}

func (s *SmartContract) testUpdate(ctx contractapi.TransactionContextInterface, sampleSecretCode string, samplePassed string , SampleTestDate string) error {
	seed, err := s.QueryCar(ctx, sampleSecretCode)

	if err != nil {
		return err
	}

  seed.SamplePassed = samplePassed
  seed.SampleTestDate = SampleTestDate
	seedAsBytes, _ := json.Marshal(seed)

	return ctx.GetStub().PutState(seed.lotNumber, seedAsBytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create seed chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting seed chaincode: %s", err.Error())
	}
}
