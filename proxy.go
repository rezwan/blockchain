package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ProxyChaincode example simple Chaincode implementation
type ProxyChaincode struct {
}

type Vote struct {
	ObjectType            string  `json:"docType"`
	VoteID                string  `json:"vote_id"`
	CampaignId            string  `json:"campaign_id"`
	CusipId               string  `json:"cusip"`
	ControlNumber         string  `json:"control_number"`
	VoteDate              string  `json:"vote_date"`
	NumberOfVoteRepresent float64 `json:"number_of_vote_represent"`
	Shares                float64 `json:"shares"`
	Iteration             string  `json:"iteration"`
	NetworkIndicator      string  `json:"network_indicator"`
	Source                int     `json:"source"`
	PP01                  string  `json:"PP01"`
	PP02                  string  `json:"PP02"`
	PP03                  string  `json:"PP03"`
	PP04                  string  `json:"PP04"`
	PP05                  string  `json:"PP05"`
	PP06                  string  `json:"PP06"`
	PP07                  string  `json:"PP07"`
	PP08                  string  `json:"PP08"`
	PP09                  string  `json:"PP09"`
	PP10                  string  `json:"PP10"`

	PP11 string `json:"PP11"`
	PP12 string `json:"PP12"`
	PP13 string `json:"PP13"`
	PP14 string `json:"PP14"`
	PP15 string `json:"PP15"`
	PP16 string `json:"PP16"`
	PP17 string `json:"PP17"`
	PP18 string `json:"PP18"`
	PP19 string `json:"PP19"`
	PP20 string `json:"PP20"`

	PP21 string `json:"PP21"`
	PP22 string `json:"PP22"`
	PP23 string `json:"PP23"`
	PP24 string `json:"PP24"`
	PP25 string `json:"PP25"`
	PP26 string `json:"PP26"`
	PP27 string `json:"PP27"`
	PP28 string `json:"PP28"`
	PP29 string `json:"PP29"`
	PP30 string `json:"PP30"`

	PP31 string `json:"PP31"`
	PP32 string `json:"PP32"`
	PP33 string `json:"PP33"`
	PP34 string `json:"PP34"`
	PP35 string `json:"PP35"`
	PP36 string `json:"PP36"`
	PP37 string `json:"PP37"`
	PP38 string `json:"PP38"`
	PP39 string `json:"PP39"`
	PP40 string `json:"PP40"`

	PP41 string `json:"PP41"`
	PP42 string `json:"PP42"`
	PP43 string `json:"PP43"`
	PP44 string `json:"PP44"`
	PP45 string `json:"PP45"`
	PP46 string `json:"PP46"`
	PP47 string `json:"PP47"`
	PP48 string `json:"PP48"`
	PP49 string `json:"PP49"`
	PP50 string `json:"PP50"`

	DR01 string `json:"DR01"`
	DR02 string `json:"DR02"`
	DR03 string `json:"DR03"`
	DR04 string `json:"DR04"`
	DR05 string `json:"DR05"`
	DR06 string `json:"DR06"`
	DR07 string `json:"DR07"`
	DR08 string `json:"DR08"`
	DR09 string `json:"DR09"`
	DR10 string `json:"DR10"`

	DR11         string `json:"DR11"`
	DR12         string `json:"DR12"`
	DR13         string `json:"DR13"`
	DR14         string `json:"DR14"`
	DR15         string `json:"DR15"`
	DR16         string `json:"DR16"`
	DR17         string `json:"DR17"`
	DR18         string `json:"DR18"`
	DR19         string `json:"DR19"`
	DR20         string `json:"DR20"`
	CreatedBy    string `json:"created_by"`
	CreatedDate  string `json:"created_date"`
	ModifiedBy   string `json:"modified_by"`
	ModifiedDate string `json:"modified_date"`
}
type BaseResult struct {
	Error string `json:"error"`
}

type Votes struct {
	BaseResult
	Votes []Vote
}

type Proposals struct {
	BaseResult
	Proposals []Proposal
}

type CusipClass struct {
	ObjectType     string `json:"docType"`
	CusipClassId   string `json:"cusip_class_id"`
	IssuerId       string `json:"issuer_id"`
	CusipClassName string `json:"cusip_class_name"`
	CreatedBy      string `json:"created_by"`
	CreatedDate    string `json:"created_date"`
	ModifiedBy     string `json:"modified_by"`
	ModifiedDate   string `json:"modified_date"`
}

type Proposal struct {
	ObjectType           string `json:"docType"`
	ProposalId           string `json:"proposal_id"`
	CampaignId           string `json:"campaign_id"`
	ProposalText         string `json:"proposal_text"`
	Options              string `json:"options"`
	ProposalType         string `json:"proposal_type"`
	DefaultVoteDirection string `json:"default_vote_direction"`
	Tag                  string `json:"tag"`
	CreatedBy            string `json:"created_by"`
	CreatedDate          string `json:"created_date"`
	ModifiedBy           string `json:"modified_by"`
	ModifiedDate         string `json:"modified_date"`
}

type ProposalVote struct {
	ProposalId   string  `json:"proposal_id"`
	ProposalText string  `json:"proposal_text"`
	For          float64 `json:"for"`
	Against      float64 `json:"against"`
	Abstain      float64 `json:"abstain"`
}

type ProposalSet struct {
	ObjectType      string `json:"docType"`
	ProposalSetId   string `json:"proposal_set_id"`
	ProposalSetName string `json:"proposal_set_name"`
	CreatedBy       string `json:"created_by"`
	CreatedDate     string `json:"created_date"`
	ModifiedBy      string `json:"modified_by"`
	ModifiedDate    string `json:"modified_date"`
}

type Issuer struct {
	ObjectType        string `json:"docType"`
	IssuerId          string `json:"issuer_id"`
	IssuerName        string `json:"issuer_name"`
	IssuerAddress     string `json:"issuer_address"`
	IssuerStockSymbol string `json:"issuer_stock_symbol"`
	IssuerWebAddress  string `json:"issuer_web_address"`
	ShortName         string `json:"short_name"`
	CreatedBy         string `json:"created_by"`
	CreatedDate       string `json:"created_date"`
	ModifiedBy        string `json:"modified_by"`
	ModifiedDate      string `json:"modified_date"`
	CusipIdList       string `json:"cusip_id_list"`
}

type Campaign struct {
	ObjectType        string `json:"docType"`
	CampaignId        string `json:"campaign_id"`
	IssuerId          string `json:"issuer_id"`
	CampaignTitle     string `json:"campaign_title"`
	CampaignStartDate string `json:"campaign_start_date"`
	CampaignEndDate   string `json:"campaign_end_date"`
	CampaignRunDate   string `json:"campaign_run_date"`
	CampaignCode      string `json:"campaign_code"`
	CreatedBy         string `json:"created_by"`
	CreatedDate       string `json:"created_date"`
	ModifiedBy        string `json:"modified_by"`
	ModifiedDate      string `json:"modified_date"`
	CusipIdList       string `json:"cusip_id_list"`
}

type ProposalMapping struct {
	ObjectType        string `json:"docType"`
	ProposalMappingId string `json:"proposal_mapping_id"`
	CampaignId        string `json:"campaign_id"`
	CusipId           string `json:"cusip_id"`
	ProposalId        string `json:"proposal_id"`
	Sequence          string `json:"sequence"`
	CreatedBy         string `json:"created_by"`
	CreatedDate       string `json:"created_date"`
	ModifiedBy        string `json:"modified_by"`
	ModifiedDate      string `json:"modified_date"`
}

type Fund struct {
	ObjectType   string `json:"docType"`
	FundId       string `json:"fund_id"`
	IssuerId     string `json:"issuer_id"`
	FundName     string `json:"fund_name"`
	CreatedBy    string `json:"created_by"`
	CreatedDate  string `json:"created_date"`
	ModifiedBy   string `json:"modified_by"`
	ModifiedDate string `json:"modified_date"`
	CusipIdList  string `json:"cusip_id_list"`
}
type Shareholder struct {
	ObjectType            string  `json:"docType"`
	ShareHolderId         string  `json:"share_holder_id"`
	CampaignId            string  `json:"campaign_id"`
	Cusip                 string  `json:"cusip"`
	ControlNumber         string  `json:"control_number"`
	NumberOfVoteRepresent float64 `json:"number_of_vote_represent"`
	Shares                float64 `json:"shares"`
	AccountNumber         string  `json:"account_number"`
	SharePercent          float64 `json:"share_percent"`
	CreatedBy             string  `json:"created_by"`
	CreatedDate           string  `json:"created_date"`
	ModifiedBy            string  `json:"modified_by"`
	ModifiedDate          string  `json:"modified_date"`
	Tape                  int     `json:"tape"`
	Seq                   float64 `json:"seq"`
}

type Cusip struct {
	ObjectType   string `json:"docType"`
	CusipId      string `json:"cusip_id"`
	IssuerId     string `json:"issuer_id"`
	FundId       string `json:"fund_id"`
	Tag          string `json:"tag"`
	CusipTitle   string `json:"cusip_title"`
	CusipClassId string `json:"cusip_class_id"`
	CusipTrustId string `json:"cusip_trust_id"`
	TaFund       string `json:"ta_fund"`
	CreatedBy    string `json:"created_by"`
	CreatedDate  string `json:"created_date"`
	ModifiedBy   string `json:"modified_by"`
	ModifiedDate string `json:"modified_date"`
}

type CusipTrust struct {
	ObjectType     string `json:"docType"`
	CusipTrustId   string `json:"cusip_trust_id"`
	IssuerId       string `json:"issuer_id"`
	CusipTrustName string `json:"cusip_trust_name"`
	CreatedBy      string `json:"created_by"`
	CreatedDate    string `json:"created_date"`
	ModifiedBy     string `json:"modified_by"`
	ModifiedDate   string `json:"modified_date"`
}
type CusipProposalMapping struct {
	ObjectType             string `json:"docType"`
	CusipProposalMappingId string `json:"cusip_proposal_mapping_id"`
	CampaignId             string `json:"campaign_id"`
	ProposalId             string `json:"proposal_id"`
	CusipId                string `json:"cusip_id"`
	Sequence               string `json:"sequence"`
	CreatedBy              string `json:"created_by"`
	CreatedDate            string `json:"created_date"`
	ModifiedBy             string `json:"modified_by"`
	ModifiedDate           string `json:"modified_date"`
	IsDeleted              string `json:"is_deleted"`
}
type CampaignCusipMapping struct {
	ObjectType        string `json:"docType"`
	CampaignMappingId string `json:"campaign_mapping_id"`
	CampaignId        string `json:"campaign_id"`
	CusipId           string `json:"cusip_id"`
	CreatedBy         string `json:"created_by"`
	CreatedDate       string `json:"created_date"`
	ModifiedBy        string `json:"modified_by"`
	ModifiedDate      string `json:"modified_date"`
}

type BenificialAggregate struct {
	ObjectType           string  `json:"docType"`
	Source               string  `json:"source"`
	SequenceNumber       int     `json:"sequence_number"`
	IssuerId             int     `json:"issuer_id"`
	CampaignId           int     `json:"campaign_id"`
	FundName             string  `json:"fund_name"`
	CusipId              string  `json:"cusip"`
	ProposalNo           int     `json:"proposal_no"`
	For                  float64 `json:"for_vote"`
	Against              float64 `json:"against_or_withhold_vote"`
	Abstain              float64 `json:"abstain_vote"`
	IsFound              bool    `json:"-"`
	BatchId              string  `json:"batch_Id"`
	BlockchainInsertTime string  `json:"blockchain_insert_time"`
}
type BenificialAggregates struct {
	BaseResult           `json:"-"`
	BenificialAggregates []BenificialAggregate `json:"data"`
	BatchId              string                `json:"batch_Id"`
}

type Issuers struct {
	BaseResult
	Issuers []Issuer
}

type ProposalSets struct {
	BaseResult
	ProposalSets []ProposalSet
}

type Campaigns struct {
	BaseResult
	Campaigns []Campaign
}
type Funds struct {
	BaseResult
	Funds []Fund
}
type Cusips struct {
	BaseResult
	Cusips []Cusip
}
type Shareholders struct {
	BaseResult
	Shareholders []Shareholder
}
type CusipClasses struct {
	BaseResult
	CusipClasses []CusipClass
}
type CusipTrusts struct {
	BaseResult
	CusipTrusts []CusipTrust
}
type CusipProposalMappings struct {
	BaseResult
	CusipProposalMappings []CusipProposalMapping
}

type CampaignCusipMappings struct {
	BaseResult
	CampaignCusipMappings []CampaignCusipMapping
}

type CampaignMaster struct {
	CampaignId                 string           `json:"campaignId"`
	CampaignTitle              string           `json:"Campaign_title"`
	CampaignStartDate          string           `json:"Campaign_start_date"`
	CampaignEndDate            string           `json:"campaign_end_date"`
	TotalSharesVoted           float64          `json:"totalSharesVoted"`
	TotalNumberOfVoteRepresent float64          `json:"total_number_of_vote_represent"`
	TotalVoteCastPercent       float64          `json:"total_vote_cast_percent"`
	CampaignDetails            []CampaignDetail `json:"data"`
}

type CampaignDetail struct {
	CampaignId         string  `json:"campaignId"`
	Cusip              string  `json:"cusip"`
	ProposalDetail     string  `json:"proposalDetail"`
	Sequence           string  `json:"-"`
	ForVoteCount       float64 `json:"forVoteCount"`
	ForVotePercent     float64 `json:"forVotePercent"`
	AgainstVoteCount   float64 `json:"againstVoteCount"`
	AgainstVotePercent float64 `json:"againstVotePercent"`
	AbstainVoteCount   float64 `json:"abstainVoteCount"`
	AbstainVotePercent float64 `json:"abstainVotePercent"`
	TotalVoted         float64 `json:"totalvoted"`
	FundId             float64 `json:"fund_id"`
	FundName           string  `json:"fund_name"`
}

func (t *ProxyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *ProxyChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call")
}

// Transaction makes payment of X units from A to B
func (t *ProxyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "CreateVote" {
		return t.CreateVote(stub, args)
	} else if function == "CreateProposal" {
		return t.CreateProposal(stub, args)
	} else if function == "CreateIssuer" {
		return t.CreateIssuer(stub, args)
	} else if function == "CreateCampaign" {
		return t.CreateCampaign(stub, args)
	} else if function == "CreateFund" {
		return t.CreateFund(stub, args)
	} else if function == "CreateCusip" {
		return t.CreateCusip(stub, args)
	} else if function == "CreateCusipClass" {
		return t.CreateCusipClass(stub, args)
	} else if function == "CreateCusipTrust" {
		return t.CreateCusipTrust(stub, args)
	} else if function == "CreateShareHolder" {
		return t.CreateShareHolder(stub, args)
	} else if function == "CreateCusipProposalMapping" {
		return t.CreateCusipProposalMapping(stub, args)
	} else if function == "CreateCampaignCusipMapping" {
		return t.CreateCampaignCusipMapping(stub, args)
	} else if function == "GetVoteById" {
		return t.GetVoteById(stub, args)
	} else if function == "GetProposalById" {
		return t.GetProposalById(stub, args)
	} else if function == "GetIssuerById" {
		return t.GetIssuerById(stub, args)
	} else if function == "GetCampaignById" {
		return t.GetCampaignById(stub, args)
	} else if function == "GetFundById" {
		return t.GetFundById(stub, args)
	} else if function == "TopNShareHolder" {
		return t.TopNShareHolder(stub, args)
	} else if function == "VotingReport" {
		return t.VotingReport(stub, args)
	} else if function == "BenificialAggregateReport" {
		return t.BenificialAggregateReport(stub, args)
	} else if function == "EndBenificialAggregateReport" {
		return t.EndBenificialAggregateReport(stub, args)
	} else if function == "GetCusipById" {
		return t.GetCusipById(stub, args)
	} else if function == "GetProposalMapping" {
		return t.GetProposalMapping(stub, args)
	} else if function == "GetVotes" {
		return t.GetVotes(stub, args)
	} else if function == "getHistory" {
		return t.GetHistory(stub, args)
	} else if function == "getAllKeyValuesFromPrefix" {
		return t.GetAllKeyValuesFromPrefix(stub, args)
	} else if function == "getRecordsByRange" {
		return t.GetRecordsByRange(stub, args)
	} else if function == "richQuery" {
		return t.RichQuery(stub, args)
	} else if function == "getQuery" {
		return t.GetQuery(stub, args)
	} else if function == "SetCustomValue" {
		return t.SetCustomValue(stub, args)
	} else if function == "GetCustomValue" {
		return t.GetCustomValue(stub, args)
	}
	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

//============================= Vote =====================================

func (t *ProxyChaincode) SetCustomValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	err = stub.PutState("custom", []byte("{arg:\""+args[0]+"\"}"))

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *ProxyChaincode) GetCustomValue(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	customValueByte, err := stub.GetState("custom")
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(customValueByte)
}

func (t *ProxyChaincode) CreateVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Vote")

	arrayResult := Votes{}
	arrayBody := []byte(args[0])

	if err = json.Unmarshal(arrayBody, &arrayResult.Votes); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	m := make(map[string]interface{})

	for _, value := range arrayResult.Votes {
		vote := value
		key := "vote_" + vote.CampaignId + "_" + vote.CusipId + "_" + vote.ControlNumber + "_" + vote.VoteID
		var reply interface{} = value
		t := reflect.TypeOf(reply)
		v := reflect.ValueOf(reply)

		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Type.Name() != "string" || (t.Field(i).Type.Name() == "string" && v.Field(i).Interface().(string) != "") {
				m[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
			}
		}
		m["ObjectType"] = "vote"
		j, _ := json.Marshal(m)
		fmt.Println(string(j))

		err = stub.PutState(key, j)

		if err != nil {
			return shim.Error(err.Error())
		}

	}
	return shim.Success(nil)
}

func (t *ProxyChaincode) BenificialAggregateReport(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Proposal")
	var arrayResult []BenificialAggregates
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult); err != nil {
		arrayResult[0].BaseResult.Error = string(arrayBody)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	issuerId := strconv.Itoa(arrayResult[0].BenificialAggregates[0].IssuerId)
	campaignId := strconv.Itoa(arrayResult[0].BenificialAggregates[0].CampaignId)

	for _, value := range arrayResult[0].BenificialAggregates {
		objBenificial := value
		objBenificial.ObjectType = "benificial_aggregate"

		proposalNo := strconv.Itoa(objBenificial.ProposalNo)
		key := "benificial_aggregate_" + issuerId + "_" + campaignId + "_" + proposalNo + "_" + objBenificial.CusipId + "_" + objBenificial.Source
		objBenificial.BatchId = arrayResult[0].BatchId
		benificialJSONasBytes, err := json.Marshal(objBenificial)
		err = stub.PutState(key, benificialJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (t *ProxyChaincode) EndBenificialAggregateReport(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Proposal")
	var arrayResult []BenificialAggregates
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult); err != nil {
		arrayResult[0].BaseResult.Error = string(arrayBody)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	var functionArgs []string

	issuerId := strconv.Itoa(arrayResult[0].BenificialAggregates[0].IssuerId)
	campaignId := strconv.Itoa(arrayResult[0].BenificialAggregates[0].CampaignId)

	functionArgs = append(functionArgs, "benificial_aggregate_"+issuerId+"_"+campaignId+"_"+"0")
	functionArgs = append(functionArgs, "benificial_aggregate_"+issuerId+"_"+campaignId+"_"+"z")
	query, _ := GetRecordsByRangeByte(stub, functionArgs)

	prevList := BenificialAggregates{}
	arrayBody1 := []byte(query)

	if err = json.Unmarshal(arrayBody1, &prevList.BenificialAggregates); err != nil {
		prevList.BaseResult.Error = string(arrayBody1)
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	for _, value := range prevList.BenificialAggregates {

		if value.BatchId != arrayResult[0].BatchId {
			proposalNo := strconv.Itoa(value.ProposalNo)
			key := "benificial_aggregate_" + issuerId + "_" + campaignId + "_" + proposalNo + "_" + value.CusipId + "_" + value.Source
			value.For = 0
			value.Abstain = 0
			value.Against = 0
			value.BatchId = arrayResult[0].BatchId

			benificialJSONasBytes, err := json.Marshal(value)
			err = stub.PutState(key, benificialJSONasBytes)

			if err != nil {
				return shim.Error(err.Error())
			}

		}

	}

	return shim.Success(nil)
}

func (t *ProxyChaincode) CreateCusipClass(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Proposal")

	arrayResult := CusipClasses{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.CusipClasses); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.CusipClasses {
		cusipClass := value
		cusipClass.ObjectType = "cusipClass"
		key := "cusipClass_" + cusipClass.IssuerId + "_" + cusipClass.CusipClassId

		cusipClassJSONasBytes, err := json.Marshal(cusipClass)
		err = stub.PutState(key, cusipClassJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (t *ProxyChaincode) CreateCusipTrust(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Proposal")

	arrayResult := CusipTrusts{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.CusipTrusts); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.CusipTrusts {
		cusipTrust := value
		cusipTrust.ObjectType = "cusipTrust"
		key := "cusipTrust_" + cusipTrust.IssuerId + "_" + cusipTrust.CusipTrustId

		cusipTrustJSONasBytes, err := json.Marshal(cusipTrust)
		err = stub.PutState(key, cusipTrustJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (t *ProxyChaincode) CreateShareHolder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Proposal")

	arrayResult := Shareholders{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.Shareholders); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.Shareholders {
		shareholder := value
		shareholder.ObjectType = "shareholder"
		key := "shareholder_" + shareholder.CampaignId + "_" + shareholder.ControlNumber + "_" + shareholder.Cusip + "_" + shareholder.ShareHolderId
		fmt.Println(key)
		shareholderJSONasBytes, err := json.Marshal(shareholder)
		err = stub.PutState(key, shareholderJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (t *ProxyChaincode) CreateProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Proposal")

	arrayResult := Proposals{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.Proposals); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.Proposals {
		proposal := value
		proposal.ObjectType = "proposal"
		key := "proposal_" + proposal.CampaignId + "_" + proposal.ProposalId
		fmt.Println(key)
		ProposalJSONasBytes, err := json.Marshal(proposal)
		err = stub.PutState(key, ProposalJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
	//
}

func (t *ProxyChaincode) CreateIssuer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Issuer")

	arrayResult := Issuers{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.Issuers); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.Issuers {
		issuer := value
		issuer.ObjectType = "issuer"
		key := "issuer_" + issuer.IssuerId
		fmt.Println(key)
		IssuerJSONasBytes, err := json.Marshal(issuer)
		err = stub.PutState(key, IssuerJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
	//
}

func (t *ProxyChaincode) CreateCampaign(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Issuer")

	arrayResult := Campaigns{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.Campaigns); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.Campaigns {
		campaign := value
		campaign.ObjectType = "campaign"
		key := "campaign_" + campaign.IssuerId + "_" + campaign.CampaignId
		fmt.Println(key)
		CampaignJSONasBytes, err := json.Marshal(campaign)
		err = stub.PutState(key, CampaignJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
	//
}

// need to delete previous cusip then insert

func (t *ProxyChaincode) GetVotes(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var functionArgs []string
	var err error
	functionArgs = append(functionArgs, "vote_"+args[0]+"_"+args[1]+"_0")
	functionArgs = append(functionArgs, "vote_"+args[0]+"_"+args[1]+"_z")
	query, _ := GetRecordsByRangeByte(stub, functionArgs)

	arrayResult := Votes{}
	arrayBody := []byte(query)

	if err = json.Unmarshal(arrayBody, &arrayResult.Votes); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}
	CampaignJSONasBytes, _ := json.Marshal(arrayResult.Votes)
	return shim.Success(CampaignJSONasBytes)
}

func (t *ProxyChaincode) GetAllKeyValuesFromPrefix(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var functionArgs []string
	functionArgs = append(functionArgs, args[0]+"_0")
	functionArgs = append(functionArgs, args[0]+"_z")
	query, _ := GetRecordsByRangeByte(stub, functionArgs)
	return shim.Success([]byte(query))
}

func GetAllKeyValuesFromPrefixString(stub shim.ChaincodeStubInterface, arg string) string {
	var functionArgs []string
	functionArgs = append(functionArgs, arg+"_0")
	functionArgs = append(functionArgs, arg+"_z")
	query, _ := GetRecordsByRangeByte(stub, functionArgs)
	strQuery := string(query[:])
	return strQuery
}

func (t *ProxyChaincode) CreateCusipProposalMapping(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Issuer")

	arrayResult := CusipProposalMappings{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.CusipProposalMappings); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	// delete previous cusips begin
	var functionArgs []string
	functionArgs = append(functionArgs, "cusipproposalmapping_"+arrayResult.CusipProposalMappings[0].CampaignId+"_"+arrayResult.CusipProposalMappings[0].CusipId+"_0")
	functionArgs = append(functionArgs, "cusipproposalmapping_"+arrayResult.CusipProposalMappings[0].CampaignId+"_"+arrayResult.CusipProposalMappings[0].CusipId+"_z")

	query, err := GetRecordsByRangeByte(stub, functionArgs)

	DelCusipList := CusipProposalMappings{}
	arrayBody = []byte(query)

	if err = json.Unmarshal(arrayBody, &DelCusipList.CusipProposalMappings); err != nil {
		DelCusipList.BaseResult.Error = string(arrayBody)
	}
	for _, cpMap := range arrayResult.CusipProposalMappings {
		//stub.DelState("cusipproposalmapping_"+ cpMap.CampaignId +"_"+ cpMap.CusipId +"_" + cpMap.ProposalId )
		cpMap.IsDeleted = "true"
		ddd, _ := json.Marshal(cpMap)
		err = stub.PutState("cusipproposalmapping_"+cpMap.CampaignId+"_"+cpMap.CusipId+"_"+cpMap.ProposalId, ddd)
	}
	//delete previous cusips end

	for _, value := range arrayResult.CusipProposalMappings {
		cpMapping := value
		cpMapping.ObjectType = "cusipproposalmapping"
		cpMapping.IsDeleted = "false"
		key := "cusipproposalmapping_" + cpMapping.CampaignId + "_" + cpMapping.CusipId + "_" + cpMapping.ProposalId
		fmt.Println(key)
		cpMappingJSONasBytes, err := json.Marshal(cpMapping)
		err = stub.PutState(key, cpMappingJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) CreateCampaignCusipMapping(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Issuer")

	arrayResult := CampaignCusipMappings{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.CampaignCusipMappings); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.CampaignCusipMappings {
		ccMapping := value
		ccMapping.ObjectType = "getcampaigncusipmapping"
		key := "getcampaigncusipmapping_" + ccMapping.CampaignId + "_" + ccMapping.CusipId
		fmt.Println(key)
		ccMappingJSONasBytes, err := json.Marshal(ccMapping)
		err = stub.PutState(key, ccMappingJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
	//
}

func (t *ProxyChaincode) GetProposalMapping(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "proposalmapping_" + args[0] + "_" + args[1]
	ProposalAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(ProposalAsBytes)
}

func (t *ProxyChaincode) GetProposalById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "proposal_" + args[0] + "_" + args[1]
	ProposalAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(ProposalAsBytes)
}

func (t *ProxyChaincode) CreateFund(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("- start creating Issuer")

	arrayResult := Funds{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.Funds); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.Funds {
		fund := value
		fund.ObjectType = "fund"
		key := "fund_" + fund.IssuerId + "_" + fund.FundId
		fmt.Println(key)
		fundJSONasBytes, err := json.Marshal(fund)
		err = stub.PutState(key, fundJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
	//
}
func (t *ProxyChaincode) CreateCusip(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	fmt.Println("- start creating Issuer")

	arrayResult := Cusips{}
	arrayBody := []byte(args[0])
	if err = json.Unmarshal(arrayBody, &arrayResult.Cusips); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	for _, value := range arrayResult.Cusips {
		cusip := value
		cusip.ObjectType = "cusip"
		key := "cusip_" + cusip.CusipId
		fmt.Println(key)
		cusipJSONasBytes, err := json.Marshal(cusip)
		err = stub.PutState(key, cusipJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
	//
}
func (t *ProxyChaincode) GetVoteById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	objectType := "vote_" + args[0] + "_" + args[1] + "_" + args[2] + "_" + args[3]
	VoteAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(VoteAsBytes)
}
func (t *ProxyChaincode) GetIssuerById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "issuer_" + args[0]
	objAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(objAsBytes)
}
func (t *ProxyChaincode) GetCampaignById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	objectType := "campaign_" + args[0] + "_" + args[1]
	objAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(objAsBytes)
}

func GetCampaignByIdByte(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	objectType := "campaign_" + args[0] + "_" + args[1]
	objAsBytes, _ := stub.GetState(objectType)
	return objAsBytes, nil
}

func (t *ProxyChaincode) GetFundById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	objectType := "fund_" + args[0] + "_" + args[1]
	objAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(objAsBytes)
}
func (t *ProxyChaincode) GetCusipById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	objectType := "cusip_" + args[0]
	objAsBytes, err := stub.GetState(objectType)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(objAsBytes)
}

func GetRecordsByRangeByte(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	startKey := args[0]
	endKey := args[1]

	resultsIterator, _ := stub.GetStateByRange(startKey, endKey)

	var buffer bytes.Buffer

	defer resultsIterator.Close()

	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()

		// Add a comma before array members, suppress it for the first array member

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Printf("- Range queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *ProxyChaincode) GetRecordsByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)

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

		// Add a comma before array members, suppress it for the first array member

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Printf("- Range queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())

}

func (t *ProxyChaincode) GetHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {

		return shim.Error("Incorrect number of arguments. Expecting 1")

	}

	keyName := args[0]

	fmt.Printf("- start getHistory: %s\n", keyName)

	resultsIterator, err := stub.GetHistoryForKey(keyName)

	if err != nil {

		return shim.Error(err.Error())

	}

	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble

	var buffer bytes.Buffer

	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {

		response, err := resultsIterator.Next()

		if err != nil {

			return shim.Error(err.Error())

		}

		// Add a comma before array members, suppress it for the first array member

		if bArrayMemberAlreadyWritten == true {

			buffer.WriteString(",")

		}

		buffer.WriteString("{\"TxId\":")

		buffer.WriteString("\"")

		buffer.WriteString(response.TxId)

		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")

		// if it was a delete operation on given key, then we need to set the

		//corresponding value null. Else, we will write the response.Value

		//as-is (as the Value itself a JSON marble)

		if response.IsDelete {

			buffer.WriteString("null")

		} else {

			buffer.WriteString(string(response.Value))

		}

		buffer.WriteString(", \"Timestamp\":")

		buffer.WriteString("\"")

		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())

		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")

		buffer.WriteString("\"")

		buffer.WriteString(strconv.FormatBool(response.IsDelete))

		buffer.WriteString("\"")

		buffer.WriteString("}")

		bArrayMemberAlreadyWritten = true

	}

	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}
func main() {
	err := shim.Start(new(ProxyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
	//RichQuery("")
}

type ByShare []Shareholder

func (a ByShare) Len() int           { return len(a) }
func (a ByShare) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByShare) Less(i, j int) bool { return a[i].Shares > a[j].Shares }

func (t *ProxyChaincode) VotingReport(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//==================== vote data ===============================
	totalNumberofShare := make(map[string]float64)
	aggregateQuery := GetAllKeyValuesFromPrefixString(stub, "benificial_aggregate_"+args[0]+"_"+args[1])

	AggregateResult := BenificialAggregates{}
	arrayBody := []byte(aggregateQuery)
	if err := json.Unmarshal(arrayBody, &AggregateResult.BenificialAggregates); err != nil {
		AggregateResult.BaseResult.Error = string(arrayBody)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	//GET VOTE DATA FOR THIS CAMPAIGN
	voteQuery := GetAllKeyValuesFromPrefixString(stub, "vote_"+args[1])

	voteResult := Votes{}
	arrayBody = []byte(voteQuery)
	if err := json.Unmarshal(arrayBody, &voteResult.Votes); err != nil {
		voteResult.BaseResult.Error = string(arrayBody)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	//GET PROPOSALS FOR THIS CAMPAIGN

	proposalQuery := GetAllKeyValuesFromPrefixString(stub, "proposal_"+args[1])
	proposalResult := Proposals{}
	arrayBody = []byte(proposalQuery)
	if err := json.Unmarshal(arrayBody, &proposalResult.Proposals); err != nil {
		proposalResult.BaseResult.Error = string(arrayBody)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	//GET Campaign INFO

	campaignResult, _ := GetCampaignByIdByte(stub, args)
	Use(err)
	campaign := Campaign{}
	json.Unmarshal(campaignResult, &campaign)

	//CALCULATE FOR EACH CUSIP

	var campaignDetails []CampaignDetail
	cusipProposalResult := CusipProposalMappings{}

	cusipProposalQuery := GetAllKeyValuesFromPrefixString(stub, "cusipproposalmapping_"+args[1])

	arrayBody = []byte(cusipProposalQuery)
	if err := json.Unmarshal(arrayBody, &cusipProposalResult.CusipProposalMappings); err != nil {
		cusipProposalResult.BaseResult.Error = string(arrayBody)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	for _, cpMap := range cusipProposalResult.CusipProposalMappings {
		var aggFor float64 = 0
		var aggAgainst float64 = 0
		var aggAbstain float64 = 0
		for _, aggregate := range AggregateResult.BenificialAggregates {
			//strProsalNo
			strProposalId := strconv.Itoa(aggregate.ProposalNo)
			if cpMap.CusipId == aggregate.CusipId && cpMap.ProposalId == strProposalId {
				aggFor += aggregate.For
				aggAgainst += aggregate.Against
				aggAbstain += aggregate.Abstain
			}
		}
		var voteFor float64 = 0
		var voteAgainst float64 = 0
		var voteAbstain float64 = 0
		for _, voteData := range voteResult.Votes {
			if cpMap.CusipId == voteData.CusipId && voteData.CampaignId == cpMap.CampaignId {
				strSequenceWithPadding := cpMap.Sequence
				if len(cpMap.Sequence) == 1 {
					strSequenceWithPadding = "0" + cpMap.Sequence
				}
				var vote interface{} = voteData
				t := reflect.TypeOf(vote)
				v := reflect.ValueOf(vote)
				for i := 0; i < t.NumField(); i++ {
					if strings.Contains(t.Field(i).Tag.Get("json"), "PP"+strSequenceWithPadding) {

						switch os := v.Field(i).Interface().(string); os {
						case "F":
							voteFor += voteData.Shares
						case "A":
							voteAgainst += voteData.Shares
						case "B":
							voteAbstain += voteData.Shares
						default:
							fmt.Printf("invalid vote input")
						}
						//isloopbreak = true
						break
					}
				}
			}
		}
		campaignDetail := CampaignDetail{}
		campaignDetail.ForVoteCount = aggFor + voteFor
		campaignDetail.AgainstVoteCount = aggAgainst + voteAgainst
		campaignDetail.AbstainVoteCount = aggAbstain + voteAbstain
		campaignDetail.CampaignId = cpMap.CampaignId
		campaignDetail.Cusip = cpMap.CusipId
		campaignDetail.Sequence = cpMap.Sequence
		campaignDetail.TotalVoted = (campaignDetail.ForVoteCount + campaignDetail.AgainstVoteCount + campaignDetail.AbstainVoteCount)

		if _, ok := totalNumberofShare[campaignDetail.Cusip]; ok {
			if totalNumberofShare[campaignDetail.Cusip] < campaignDetail.TotalVoted {
				totalNumberofShare[campaignDetail.Cusip] = campaignDetail.TotalVoted
			}
		} else {
			totalNumberofShare[campaignDetail.Cusip] = campaignDetail.TotalVoted
		}

		proposalText := ""
		for _, propo := range proposalResult.Proposals {
			if cpMap.ProposalId == propo.ProposalId {
				proposalText = propo.ProposalText
			}
		}
		campaignDetail.ProposalDetail = proposalText
		campaignDetails = append(campaignDetails, campaignDetail)
	}

	for i, detail := range campaignDetails {
		if detail.TotalVoted > 0 {
			detail.ForVotePercent = ((detail.ForVoteCount / detail.TotalVoted) * 100)
			detail.AgainstVotePercent = ((detail.AgainstVoteCount / detail.TotalVoted) * 100)
			detail.AbstainVotePercent = ((detail.AbstainVoteCount / detail.TotalVoted) * 100)
		} else {
			detail.ForVotePercent = 0
			detail.AgainstVotePercent = 0
			detail.AbstainVotePercent = 0
		}
		campaignDetails[i] = detail
	}

	newCampaign := CampaignMaster{}
	newCampaign.CampaignId = campaign.CampaignId
	newCampaign.CampaignTitle = campaign.CampaignTitle
	newCampaign.CampaignStartDate = campaign.CampaignStartDate
	newCampaign.CampaignEndDate = campaign.CampaignEndDate
	newCampaign.TotalSharesVoted = 0
	newCampaign.CampaignDetails = campaignDetails

	for _, v := range totalNumberofShare {
		newCampaign.TotalSharesVoted += v
	}

	result, _ := json.Marshal(newCampaign)
	return shim.Success(result)
}

func (t *ProxyChaincode) TopNShareHolder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//query:= `{"selector":{"docType":{"$eq":"shareholder"},"campaign_id":{"$eq":"`+ args[0] +`"}},"fields":["docType","account_number","share_holder_id","campaign_id","control_number","number_of_vote_represent","shares"]}`
	var functionArgs []string
	functionArgs = append(functionArgs, "shareholder_"+args[1]+"_0")
	functionArgs = append(functionArgs, "shareholder_"+args[1]+"_z")

	var shareHolderN, _ = strconv.Atoi(args[2])

	query, err := GetRecordsByRangeByte(stub, functionArgs)

	arrayResult := Shareholders{}
	arrayBody := []byte(query)

	if err = json.Unmarshal(arrayBody, &arrayResult.Shareholders); err != nil {
		arrayResult.BaseResult.Error = string(arrayBody)
	}

	var TotalNumberOfShare float64
	TotalNumberOfShare = 0
	for _, value := range arrayResult.Shareholders {
		shareholder := value
		TotalNumberOfShare += shareholder.Shares
	}

	campaignResult, _ := GetCampaignByIdByte(stub, args)

	campaign := Campaign{}
	json.Unmarshal(campaignResult, &campaign)

	sort.Sort(ByShare(arrayResult.Shareholders))

	var finalResult []Shareholder
	intlen := len(arrayResult.Shareholders)

	if intlen > shareHolderN {
		intlen = shareHolderN
	}

	for _, val := range arrayResult.Shareholders[0:intlen] {
		finalResult = append(finalResult, val)
	}

	for i, val := range finalResult {
		finalResult[i].SharePercent = float64(val.Shares) / float64(TotalNumberOfShare)
		finalResult[i].SharePercent = Round(finalResult[i].SharePercent, .5, 2)
	}

	shareHolderJSONasBytes, err := json.Marshal(finalResult)
	if err != nil {
		return shim.Error(err.Error())
	}

	var buffer bytes.Buffer
	buffer.WriteString("{\"data\":")
	buffer.WriteString("\"")
	buffer.WriteString(string(shareHolderJSONasBytes))
	buffer.WriteString("\",")

	buffer.WriteString("\"meetingDate\":")
	buffer.WriteString("\"")
	buffer.WriteString(campaign.CampaignRunDate)
	buffer.WriteString("\"")

	buffer.WriteString("}")

	return shim.Success(buffer.Bytes())
}

func (t *ProxyChaincode) RichQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
func (t *ProxyChaincode) GetQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryAsBytes, err := stub.GetState(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryAsBytes)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		if err != nil {
			return nil, err
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf(buffer.String())

	return buffer.Bytes(), nil
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
