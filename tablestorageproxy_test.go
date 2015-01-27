package gohavestorage

import (
	"encoding/json"
	"fmt"
	"gohavestorage/gohavestoragecommon"
	"reflect"
	"testing"
)

var Key = ""
var Account = ""
var Table = "TestTable"

func TestTableMethods(t *testing.T) {
	goHaveStorage := NewWithDebug(Account, Key, false)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	httpStatusCode := tableStorageProxy.CreateTable("TableForTestingTableMethods")
	if httpStatusCode != 201 {
		fmt.Printf("Faild http code other than expected:%d", httpStatusCode)
		t.Fail()
	}

	body, httpStatusCode := tableStorageProxy.QueryTables()
	if httpStatusCode != 200 {
		fmt.Printf("Faild http code other than expected:%d", httpStatusCode)
		t.Fail()
	}
	if strings.Contains(string(body), "\"TableName\":\"TableForTestingTableMethods\"") != true {
		t.Fail()
	}

	httpStatusCode = tableStorageProxy.DeleteTable("TableForTestingTableMethods")
	if httpStatusCode != 204 {
		fmt.Printf("Faild http code other than expected:%d", httpStatusCode)
		t.Fail()
	}
}

func TestInsertEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	entity := &TestEntity{}
	entity.PartitionKey = "ABC"
	entity.RowKey = "123"
	entity.Property1 = "Value1"
	entity.Property2 = "Value2"
	entity.Property3 = "Value3"

	json1, _ := json.Marshal(entity)
	tableStorageProxy.InsertEntity(Table, json1)

	entity.RowKey = "456"
	json2, _ := json.Marshal(entity)
	tableStorageProxy.InsertEntity(Table, json2)

	entity.RowKey = "789"
	json3, _ := json.Marshal(entity)
	tableStorageProxy.InsertEntity(Table, json3)
}

func TestQueryEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntity(Table, "ABC", "123", "")
}

func TestQueryEntityWithSelect(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntity(Table, "ABC", "123", "RowKey,Property1,Property3")
}

func TestQueryEntities(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntities(Table, "", "", "")
}

func TestQueryEntitiesWithSelect(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntities(Table, "RowKey,Property1,Property3", "", "")
}

func TestQueryEntitiesWithTop(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntities(Table, "", "", "1")
}

func TestQueryEntitiesWithSelectAndTop(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntities(Table, "RowKey,Property1,Property3", "", "1")
}

func TestQueryEntitiesWithSelectAndFilterAndTop(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.QueryEntities(Table, "RowKey,Property1,Property3", "RowKey%20gt%20'123'", "1")
}

func TestDeleteEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	tableStorageProxy.DeleteEntity(Table, "ABC", "123")
}

func TestUpdateEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	entity := &TestEntity{}
	entity.PartitionKey = "ABC"
	entity.RowKey = "123"
	entity.Property1 = "Value1"
	entity.Property2 = "Value2"
	entity.Property3 = "Value3"

	json, _ := json.Marshal(entity)
	tableStorageProxy.UpdateEntity(Table, "ABC", "456", json)
}

func TestMergeEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	entity := &TestEntity{}
	entity.PartitionKey = "ABC"
	entity.RowKey = "123"
	entity.Property1 = "Value1"
	entity.Property2 = "Value2"
	entity.Property3 = "Value3"

	json, _ := json.Marshal(entity)
	tableStorageProxy.MergeEntity(Table, "ABC", "456", json)
}

func TestInsertOrMergeEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	entity := &TestEntity{}
	entity.PartitionKey = "ABC"
	entity.RowKey = "123"
	entity.Property1 = "Value1"
	entity.Property2 = "Value2"
	entity.Property3 = "Value3"

	json, _ := json.Marshal(entity)
	tableStorageProxy.InsertOrMergeEntity(Table, "ABC", "456", json)
}

func TestInsertOrReplaceEntity(t *testing.T) {
	goHaveStorage := New(Account, Key)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	entity := &TestEntity{}
	entity.PartitionKey = "ABC"
	entity.RowKey = "123"
	entity.Property1 = "Value1"
	entity.Property2 = "Value2"
	entity.Property3 = "Value3"

	json, _ := json.Marshal(entity)
	tableStorageProxy.InsertOrReplaceEntity(Table, "ABC", "456", json)
}

func TestTableServiceProperties(t *testing.T) {
	goHaveStorage := NewWithDebug(Account, Key, false)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()
	properties, _ := tableStorageProxy.GetTableServiceProperties()
	httpStatusCode := tableStorageProxy.SetTableServiceProperties(properties)

	lastestProperties, _ := tableStorageProxy.GetTableServiceProperties()

	if httpStatusCode != 202 {
		fmt.Printf("Faild http code other than expected:%d", httpStatusCode)
		t.Fail()
	}
	if reflect.DeepEqual(properties, lastestProperties) == false {
		fmt.Printf("Dump:\n%+v\n\nvs\n\n%+v", properties, lastestProperties)
		t.Fail()
	}
}

func TestGetTableServiceStats(t *testing.T) {
	goHaveStorage := NewWithDebug(Account, Key, false)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()
	stats, httpStatusCode := tableStorageProxy.GetTableServiceStats()

	if httpStatusCode != 200 {
		fmt.Printf("Faild http code other than expected:%d", httpStatusCode)
		t.Fail()
	}
	if stats.GeoReplication.Status == "" || stats.GeoReplication.LastSyncTime == "" {
		t.Fail()
	}
}

func TestTableACL(t *testing.T) {
	goHaveStorage := NewWithDebug(Account, Key, false)
	tableStorageProxy := goHaveStorage.NewTableStorageProxy()

	accessPolicy := gohavestoragecommon.AccessPolicy{Start: "2014-12-31T00:00:00.0000000Z", Expiry: "2114-12-31T00:00:00.0000000Z", Permission: "raud"}
	signedIdentifier := gohavestoragecommon.SignedIdentifier{Id: "b54df8ab0e2d52759110f48c8d0c19e2", AccessPolicy: accessPolicy}
	signedIdentifiers := &gohavestoragecommon.SignedIdentifiers{[]gohavestoragecommon.SignedIdentifier{signedIdentifier}}
	tableStorageProxy.SetTableACL(Table, signedIdentifiers)

	acl, httpStatusCode := tableStorageProxy.GetTableACL(Table)

	if httpStatusCode != 200 {
		fmt.Printf("Faild http code other than expected:%d", httpStatusCode)
		t.Fail()
	}
	if reflect.DeepEqual(signedIdentifiers, acl) == false {
		fmt.Printf("Dump:\n%+v\n\nvs\n\n%+v", signedIdentifiers, acl)
		t.Fail()
	}
}

type TestEntity struct {
	PartitionKey string
	RowKey       string
	Property1    string
	Property2    string
	Property3    string
}
