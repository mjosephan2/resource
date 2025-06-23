package hashtable

import (
	"fmt"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	ht := NewHashTable(10)

	if ht.capacity != 10 {
		t.Errorf("Expected capacity 10, got %d", ht.capacity)
	}

	if len(ht.buckets) != 10 {
		t.Errorf("Expected buckets length 10, got %d", len(ht.buckets))
	}

	if ht.size != 0 {
		t.Errorf("Expected initial size 0, got %d", ht.size)
	}

	if ht.alpha != 0.75 {
		t.Errorf("Expected alpha 0.75, got %f", ht.alpha)
	}
}

func TestHashTable_SetAndGet(t *testing.T) {
	ht := NewHashTable(10)

	// Test basic set and get
	ht.Set("key1", "value1")
	ht.Set("key2", "value2")

	if val := ht.Get("key1"); val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	if val := ht.Get("key2"); val != "value2" {
		t.Errorf("Expected 'value2', got '%s'", val)
	}

	// Test getting non-existent key
	if val := ht.Get("nonexistent"); val != "" {
		t.Errorf("Expected empty string for non-existent key, got '%s'", val)
	}
}

func TestHashTable_UpdateValue(t *testing.T) {
	ht := NewHashTable(10)

	// Set initial value
	ht.Set("key1", "value1")

	// Update the value
	ht.Set("key1", "updated_value")

	if val := ht.Get("key1"); val != "updated_value" {
		t.Errorf("Expected 'updated_value', got '%s'", val)
	}

	// Size should still be 1 since we updated, not added
	if ht.size != 1 {
		t.Errorf("Expected size 1 after update, got %d", ht.size)
	}
}

func TestHashTable_CollisionHandling(t *testing.T) {
	ht := NewHashTable(1) // Force collisions by using capacity 1

	ht.Set("key1", "value1")
	ht.Set("key2", "value2")
	ht.Set("key3", "value3")

	if val := ht.Get("key1"); val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	if val := ht.Get("key2"); val != "value2" {
		t.Errorf("Expected 'value2', got '%s'", val)
	}

	if val := ht.Get("key3"); val != "value3" {
		t.Errorf("Expected 'value3', got '%s'", val)
	}

	if ht.size != 3 {
		t.Errorf("Expected size 3, got %d", ht.size)
	}
}

func TestHashTable_Rehashing(t *testing.T) {
	ht := NewHashTable(4) // Small capacity to trigger rehashing

	// Add items to trigger rehashing (load factor > 0.75)
	ht.Set("key1", "value1")
	ht.Set("key2", "value2")
	ht.Set("key3", "value3")
	ht.Set("key4", "value4") // This should trigger rehashing

	// Verify all values are still accessible after rehashing
	if val := ht.Get("key1"); val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	if val := ht.Get("key2"); val != "value2" {
		t.Errorf("Expected 'value2', got '%s'", val)
	}

	if val := ht.Get("key3"); val != "value3" {
		t.Errorf("Expected 'value3', got '%s'", val)
	}

	if val := ht.Get("key4"); val != "value4" {
		t.Errorf("Expected 'value4', got '%s'", val)
	}

	// Capacity should have doubled
	if ht.capacity != 8 {
		t.Errorf("Expected capacity 8 after rehashing, got %d", ht.capacity)
	}
}

func TestHashTable_HashFunction(t *testing.T) {
	ht := NewHashTable(10)

	// Test that same key produces same hash
	hash1 := ht.hash("test_key")
	hash2 := ht.hash("test_key")

	if hash1 != hash2 {
		t.Errorf("Hash function not consistent for same key: %d != %d", hash1, hash2)
	}

	// Test that hash is within bounds
	if hash1 < 0 || hash1 >= ht.capacity {
		t.Errorf("Hash value %d is out of bounds [0, %d)", hash1, ht.capacity)
	}
}

func TestHashTable_EmptyTable(t *testing.T) {
	ht := NewHashTable(10)

	// Test getting from empty table
	if val := ht.Get("any_key"); val != "" {
		t.Errorf("Expected empty string from empty table, got '%s'", val)
	}

	if ht.size != 0 {
		t.Errorf("Expected size 0 for empty table, got %d", ht.size)
	}
}

func TestHashTable_LargeDataSet(t *testing.T) {
	ht := NewHashTable(100)

	// Add many items
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		ht.Set(key, value)
	}

	// Verify all items can be retrieved
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		expectedValue := fmt.Sprintf("value_%d", i)
		if val := ht.Get(key); val != expectedValue {
			t.Errorf("Expected '%s', got '%s' for key '%s'", expectedValue, val, key)
		}
	}

	if ht.size != 1000 {
		t.Errorf("Expected size 1000, got %d", ht.size)
	}
}

func TestHashTable_SpecialCharacters(t *testing.T) {
	ht := NewHashTable(10)

	// Test with special characters and unicode
	testCases := []struct {
		key   string
		value string
	}{
		{"key with spaces", "value with spaces"},
		{"key-with-dashes", "value-with-dashes"},
		{"key_with_underscores", "value_with_underscores"},
		{"key123", "value123"},
		{"中文键", "中文值"},
		{"key!@#$%", "value!@#$%"},
		{"", "empty_key_value"},
	}

	for _, tc := range testCases {
		ht.Set(tc.key, tc.value)
		if val := ht.Get(tc.key); val != tc.value {
			t.Errorf("Expected '%s', got '%s' for key '%s'", tc.value, val, tc.key)
		}
	}
}

func BenchmarkHashTable_Set(b *testing.B) {
	ht := NewHashTable(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		ht.Set(key, value)
	}
}

func BenchmarkHashTable_Get(b *testing.B) {
	ht := NewHashTable(1000)

	// Pre-populate the hash table
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		ht.Set(key, value)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i%1000)
		ht.Get(key)
	}
}
