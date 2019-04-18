// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Files", testFiles)
	t.Run("NginxConfigFiles", testNginxConfigFiles)
	t.Run("Services", testServices)
}

func TestDelete(t *testing.T) {
	t.Run("Files", testFilesDelete)
	t.Run("NginxConfigFiles", testNginxConfigFilesDelete)
	t.Run("Services", testServicesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Files", testFilesQueryDeleteAll)
	t.Run("NginxConfigFiles", testNginxConfigFilesQueryDeleteAll)
	t.Run("Services", testServicesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Files", testFilesSliceDeleteAll)
	t.Run("NginxConfigFiles", testNginxConfigFilesSliceDeleteAll)
	t.Run("Services", testServicesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Files", testFilesExists)
	t.Run("NginxConfigFiles", testNginxConfigFilesExists)
	t.Run("Services", testServicesExists)
}

func TestFind(t *testing.T) {
	t.Run("Files", testFilesFind)
	t.Run("NginxConfigFiles", testNginxConfigFilesFind)
	t.Run("Services", testServicesFind)
}

func TestBind(t *testing.T) {
	t.Run("Files", testFilesBind)
	t.Run("NginxConfigFiles", testNginxConfigFilesBind)
	t.Run("Services", testServicesBind)
}

func TestOne(t *testing.T) {
	t.Run("Files", testFilesOne)
	t.Run("NginxConfigFiles", testNginxConfigFilesOne)
	t.Run("Services", testServicesOne)
}

func TestAll(t *testing.T) {
	t.Run("Files", testFilesAll)
	t.Run("NginxConfigFiles", testNginxConfigFilesAll)
	t.Run("Services", testServicesAll)
}

func TestCount(t *testing.T) {
	t.Run("Files", testFilesCount)
	t.Run("NginxConfigFiles", testNginxConfigFilesCount)
	t.Run("Services", testServicesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Files", testFilesHooks)
	t.Run("NginxConfigFiles", testNginxConfigFilesHooks)
	t.Run("Services", testServicesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Files", testFilesInsert)
	t.Run("Files", testFilesInsertWhitelist)
	t.Run("NginxConfigFiles", testNginxConfigFilesInsert)
	t.Run("NginxConfigFiles", testNginxConfigFilesInsertWhitelist)
	t.Run("Services", testServicesInsert)
	t.Run("Services", testServicesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("NginxConfigFileToServiceUsingService", testNginxConfigFileToOneServiceUsingService)
	t.Run("ServiceToFileUsingFile", testServiceToOneFileUsingFile)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("FileToServices", testFileToManyServices)
	t.Run("ServiceToNginxConfigFiles", testServiceToManyNginxConfigFiles)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("NginxConfigFileToServiceUsingNginxConfigFiles", testNginxConfigFileToOneSetOpServiceUsingService)
	t.Run("ServiceToFileUsingServices", testServiceToOneSetOpFileUsingFile)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("NginxConfigFileToServiceUsingNginxConfigFiles", testNginxConfigFileToOneRemoveOpServiceUsingService)
	t.Run("ServiceToFileUsingServices", testServiceToOneRemoveOpFileUsingFile)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("FileToServices", testFileToManyAddOpServices)
	t.Run("ServiceToNginxConfigFiles", testServiceToManyAddOpNginxConfigFiles)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("FileToServices", testFileToManySetOpServices)
	t.Run("ServiceToNginxConfigFiles", testServiceToManySetOpNginxConfigFiles)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("FileToServices", testFileToManyRemoveOpServices)
	t.Run("ServiceToNginxConfigFiles", testServiceToManyRemoveOpNginxConfigFiles)
}

func TestReload(t *testing.T) {
	t.Run("Files", testFilesReload)
	t.Run("NginxConfigFiles", testNginxConfigFilesReload)
	t.Run("Services", testServicesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Files", testFilesReloadAll)
	t.Run("NginxConfigFiles", testNginxConfigFilesReloadAll)
	t.Run("Services", testServicesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Files", testFilesSelect)
	t.Run("NginxConfigFiles", testNginxConfigFilesSelect)
	t.Run("Services", testServicesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Files", testFilesUpdate)
	t.Run("NginxConfigFiles", testNginxConfigFilesUpdate)
	t.Run("Services", testServicesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Files", testFilesSliceUpdateAll)
	t.Run("NginxConfigFiles", testNginxConfigFilesSliceUpdateAll)
	t.Run("Services", testServicesSliceUpdateAll)
}
