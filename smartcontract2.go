// smartcontract.go

// asset-transfer-basic/chaincode-go/chaincode/smartcontract.go

// Incase error occurs in creating item, whereas steel has already been reduced, then restore the quantity of steel back to the original quantity

package chaincode

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	// AppraisedValue int    `json:"AppraisedValue"`
	// Color          string `json:"Color"`
	Name string `json:"Name"`
	ID   string `json:"ID"`
	// Owner          string `json:"Owner"`
	Number int `json:"Number"`
}

var steelStrengthProbability float64 = 0.5
var dieCastingProbability float64 = 0.75

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		// {ID: "steel", Number: 1000},
		{Name: "steel", ID: "steel", Number: 400000},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, name string, id string, number int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		Name:   name,
		ID:     id,
		Number: number,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, name string, id string, number int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		Name:   name,
		ID:     id,
		Number: number,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, name string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(name)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

/*
// TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return "", err
	}

	oldOwner := asset.Owner
	asset.Owner = newOwner

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldOwner, nil
}

*/

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func randomBoolWithBias(trueProb float64) bool {
	// Seed the random number generator
	// rand.Seed(time.Now().UnixNano())

	// Generate a random float between 0 and 1
	randomNumber := rand.Float64()

	// Return true if the number is less than the true probability threshold
	return randomNumber < trueProb
}

// func checkSteelStrength() bool {
// 	// actually this will contain a Computer Vision based code to check steel strength

// 	// return randomBoolWithBias(0.5)
// 	// return randomBoolWithBias(1) // will surely return true only
// 	return true
// }

// func checkDieCasting() bool {
// 	// parameter inside randomBoolWithBias is the probability of returning true

// 	return randomBoolWithBias(0.75)
// 	// return true
// }

func checkSteelStrength() bool {
	return randomBoolWithBias(steelStrengthProbability)
}

func checkDieCasting() bool {
	return randomBoolWithBias(dieCastingProbability)
}

// /*
// Smart contract that converts one asset into another
// func (s *SmartContract) MakeItem(ctx contractapi.TransactionContextInterface, name string, id string, number int) error {
func (s *SmartContract) MakeDoor(ctx contractapi.TransactionContextInterface, steelStrengthProb float64, dieCastingProb float64) error {

	steelStrengthProbability = steelStrengthProb
	dieCastingProbability = dieCastingProb

	steelAsset, err := s.ReadAsset(ctx, "steel")
	if err != nil {
		return fmt.Errorf("failed to read steel asset: %w", err) // Wrap original error for context
	}

	needed_steel := 10

	if steelAsset.Number < needed_steel {
		return fmt.Errorf("insufficient steel: needed %d, only %d available", needed_steel, steelAsset.Number)
	} else if !checkSteelStrength() {
		return fmt.Errorf("steel is not of the required strength. Aborting")
	}

	steelAsset.Number -= needed_steel
	err = s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number)
	if err != nil {
		return fmt.Errorf("failed to update steel asset: %w", err)
	}

	if !checkDieCasting() {
		err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
		if err2 != nil {
			return fmt.Errorf("die Casting is not of proper strength. need to cast again\nAnother error - %w", err2)
		}
		return fmt.Errorf("die Casting is not of proper strength. need to cast again")
	}

	// below is the code to add door asset to the blockchain
	name := "door"
	exists, err := s.AssetExists(ctx, name)
	if err != nil {
		err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
		if err2 != nil {
			return fmt.Errorf("failed to check for %s asset: %w\nAnother error - %w", name, err, err2)
		}
		return fmt.Errorf("failed to check for %s asset: %w", name, err)
	}

	if !exists {
		// Create item asset if it doesn't exist
		err = s.CreateAsset(ctx, name, name, 1)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to create item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to create item asset - %s: %w", name, err)
		}
	} else {
		// Increment item asset count
		itemAsset, err := s.ReadAsset(ctx, name)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to read item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to read item asset - %s: %w", name, err)
		}
		itemAsset.Number += 1
		err = s.UpdateAsset(ctx, name, name, itemAsset.Number)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to update item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to update item asset - %s: %w", name, err)
		}
	}

	return nil
}

func (s *SmartContract) MakeItem(ctx contractapi.TransactionContextInterface, name string, id string, string_number string) error {

	number, err := strconv.Atoi(string_number)

	if err != nil {
		// ... handle error
		return fmt.Errorf("%s", err)
	}

	steelRequired := make(map[string]int)
	steelRequired["body"] = 50
	steelRequired["door"] = 5
	steelRequired["chassis"] = 100
	steelRequired["engine"] = 100
	steelRequired["transmission"] = 100
	steelRequired["suspension"] = 100
	steelRequired["wheels"] = 2

	_, exists := steelRequired[name]

	if !exists {
		return fmt.Errorf("cannot create item with name - %s", name)
	}

	// Check steel asset availability
	steelAsset, err := s.ReadAsset(ctx, "steel")
	if err != nil {
		return fmt.Errorf("failed to read steel asset: %w", err) // Wrap original error for context
	}

	needed_steel := steelRequired[name] * number

	if steelAsset.Number < needed_steel {
		return fmt.Errorf("insufficient steel: needed %d, only %d available", needed_steel, steelAsset.Number)
	}

	// Update steel asset quantity (assuming it's decremented after use)
	steelAsset.Number -= needed_steel
	err = s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number)
	if err != nil {
		return fmt.Errorf("failed to update steel asset: %w", err)
	}

	// check for FT asset existence here ---CONTINUE HERE -------
	if name == "engine" || name == "chassis" {
		err = s.CreateAsset(ctx, name, id, number)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to create item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to create item asset - %s: %w", name, err)
		}

		return nil
	}

	// Check if the asset already exists
	exists, err = s.AssetExists(ctx, name)
	if err != nil {
		err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
		if err2 != nil {
			return fmt.Errorf("failed to check for %s asset: %w\nAnother error - %w", name, err, err2)
		}
		return fmt.Errorf("failed to check for %s asset: %w", name, err)
	}

	if exists {
		// Increment item asset count
		itemAsset, err := s.ReadAsset(ctx, name)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to read item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to read item asset - %s: %w", name, err)
		}
		itemAsset.Number += number
		err = s.UpdateAsset(ctx, name, name, itemAsset.Number)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to update item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to update item asset - %s: %w", name, err)
		}
	} else {
		// Create item asset if it doesn't exist
		err = s.CreateAsset(ctx, name, name, number)
		if err != nil {
			err2 := s.UpdateAsset(ctx, "steel", "steel", steelAsset.Number+needed_steel)
			if err2 != nil {
				return fmt.Errorf("failed to create item asset - %s: %w\nAnother error - %w", name, err, err2)
			}
			return fmt.Errorf("failed to create item asset - %s: %w", name, err)
		}
	}

	return nil // Indicate successful item creation
}

// */

/*
// Smart contract that converts one asset into another
func (s *SmartContract) MakeBody(ctx contractapi.TransactionContextInterface, id string, id2 string) error {

	// Check steel asset availability
	steelAsset, err := s.ReadAsset(ctx, "steel")
	if err != nil {
		return fmt.Errorf("failed to read steel asset: %w", err) // Wrap original error for context
	}

	if steelAsset.Number < 50 {
		return fmt.Errorf("insufficient steel: need at least 50, only %d available", steelAsset.Number)
	}

	// Update steel asset quantity (assuming it's decremented after use)
	steelAsset.Number -= 50
	err = s.UpdateAsset(ctx, "steel", steelAsset.Number)
	if err != nil {
		return fmt.Errorf("failed to update steel asset: %w", err)
	}

	// Check if body asset already exists
	exists, err := s.AssetExists(ctx, "body")
	if err != nil {CONTINUEent body asset count
		bodyAsset, err := s.ReadAsset(ctx, "body")
		if err != nil {
			return fmt.Errorf("failed to read body asset: %w", err)
		}
		bodyAsset.Number++
		err = s.UpdateAsset(ctx, "body", bodyAsset.Number)
		if err != nil {
			return fmt.Errorf("failed to update body asset: %w", err)
		}
	} else {
		// Create body asset if it doesn't exist
		err = s.CreateAsset(ctx, "body", 1)
		if err != nil {
			return fmt.Errorf("failed to create body asset: %w", err)
		}
	}

	return nil // Indicate successful body creation
}

*/
