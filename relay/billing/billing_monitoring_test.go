package billing

import (
	"testing"
	"time"

	"github.com/songquanpeng/one-api/common/metrics"
)

// MockMetricsRecorder for testing billing monitoring
type MockMetricsRecorder struct {
	BillingOperations []BillingOperationRecord
	BillingTimeouts   []BillingTimeoutRecord
	BillingErrors     []BillingErrorRecord
}

type BillingOperationRecord struct {
	StartTime   time.Time
	Operation   string
	Success     bool
	UserId      int
	ChannelId   int
	ModelName   string
	QuotaAmount float64
}

type BillingTimeoutRecord struct {
	UserId         int
	ChannelId      int
	ModelName      string
	EstimatedQuota float64
	ElapsedTime    time.Duration
}

type BillingErrorRecord struct {
	ErrorType string
	Operation string
	UserId    int
	ChannelId int
	ModelName string
}

func (m *MockMetricsRecorder) RecordBillingOperation(startTime time.Time, operation string, success bool, userId int, channelId int, modelName string, quotaAmount float64) {
	m.BillingOperations = append(m.BillingOperations, BillingOperationRecord{
		StartTime:   startTime,
		Operation:   operation,
		Success:     success,
		UserId:      userId,
		ChannelId:   channelId,
		ModelName:   modelName,
		QuotaAmount: quotaAmount,
	})
}

func (m *MockMetricsRecorder) RecordBillingTimeout(userId int, channelId int, modelName string, estimatedQuota float64, elapsedTime time.Duration) {
	m.BillingTimeouts = append(m.BillingTimeouts, BillingTimeoutRecord{
		UserId:         userId,
		ChannelId:      channelId,
		ModelName:      modelName,
		EstimatedQuota: estimatedQuota,
		ElapsedTime:    elapsedTime,
	})
}

func (m *MockMetricsRecorder) RecordBillingError(errorType, operation string, userId int, channelId int, modelName string) {
	m.BillingErrors = append(m.BillingErrors, BillingErrorRecord{
		ErrorType: errorType,
		Operation: operation,
		UserId:    userId,
		ChannelId: channelId,
		ModelName: modelName,
	})
}

// Implement other required methods as no-ops for testing
func (m *MockMetricsRecorder) RecordHTTPRequest(startTime time.Time, path, method, statusCode string) {
}
func (m *MockMetricsRecorder) RecordHTTPActiveRequest(path, method string, delta float64) {}
func (m *MockMetricsRecorder) RecordRelayRequest(startTime time.Time, channelId int, channelType, model, userId string, success bool, promptTokens, completionTokens int, quotaUsed float64) {
}
func (m *MockMetricsRecorder) UpdateChannelMetrics(channelId int, channelName, channelType string, status int, balance float64, responseTimeMs int, successRate float64) {
}
func (m *MockMetricsRecorder) UpdateChannelRequestsInFlight(channelId int, channelName, channelType string, delta float64) {
}
func (m *MockMetricsRecorder) RecordUserMetrics(userId, username, group string, quotaUsed float64, promptTokens, completionTokens int, balance float64) {
}
func (m *MockMetricsRecorder) RecordDBQuery(startTime time.Time, operation, table string, success bool) {
}
func (m *MockMetricsRecorder) UpdateDBConnectionMetrics(inUse, idle int)                            {}
func (m *MockMetricsRecorder) RecordRedisCommand(startTime time.Time, command string, success bool) {}
func (m *MockMetricsRecorder) UpdateRedisConnectionMetrics(active int)                              {}
func (m *MockMetricsRecorder) RecordRateLimitHit(limitType, identifier string)                      {}
func (m *MockMetricsRecorder) UpdateRateLimitRemaining(limitType, identifier string, remaining int) {}
func (m *MockMetricsRecorder) RecordTokenAuth(success bool)                                         {}
func (m *MockMetricsRecorder) UpdateActiveTokens(userId, tokenName string, count int)               {}
func (m *MockMetricsRecorder) RecordError(errorType, component string)                              {}
func (m *MockMetricsRecorder) RecordModelUsage(modelName, channelType string, latency time.Duration) {
}
func (m *MockMetricsRecorder) UpdateBillingStats(totalBillingOperations, successfulBillingOperations, failedBillingOperations int64) {
}
func (m *MockMetricsRecorder) InitSystemMetrics(version, buildTime, goVersion string, startTime time.Time) {
}

func TestBillingMonitoring(t *testing.T) {
	// Setup mock metrics recorder
	mockRecorder := &MockMetricsRecorder{}
	originalRecorder := metrics.GlobalRecorder
	metrics.GlobalRecorder = mockRecorder
	defer func() {
		metrics.GlobalRecorder = originalRecorder
	}()

	// Test direct metrics recording (without database operations)
	startTime := time.Now()
	userId := 123
	channelId := 456
	modelName := "gpt-4.1"
	quotaAmount := 1000.0

	// Record a successful billing operation
	metrics.GlobalRecorder.RecordBillingOperation(startTime, "post_consume_detailed", true, userId, channelId, modelName, quotaAmount)

	// Verify billing operation was recorded
	if len(mockRecorder.BillingOperations) != 1 {
		t.Errorf("Expected 1 billing operation record, got %d", len(mockRecorder.BillingOperations))
	}

	operation := mockRecorder.BillingOperations[0]
	if operation.Operation != "post_consume_detailed" {
		t.Errorf("Expected operation 'post_consume_detailed', got '%s'", operation.Operation)
	}
	if operation.Success != true {
		t.Errorf("Expected successful operation, got %v", operation.Success)
	}
	if operation.UserId != userId {
		t.Errorf("Expected userId %d, got %d", userId, operation.UserId)
	}
	if operation.ChannelId != channelId {
		t.Errorf("Expected channelId %d, got %d", channelId, operation.ChannelId)
	}
	if operation.ModelName != modelName {
		t.Errorf("Expected modelName '%s', got '%s'", modelName, operation.ModelName)
	}
	if operation.QuotaAmount != quotaAmount {
		t.Errorf("Expected quotaAmount %f, got %f", quotaAmount, operation.QuotaAmount)
	}
}

func TestBillingErrorMonitoring(t *testing.T) {
	// Setup mock metrics recorder
	mockRecorder := &MockMetricsRecorder{}
	originalRecorder := metrics.GlobalRecorder
	metrics.GlobalRecorder = mockRecorder
	defer func() {
		metrics.GlobalRecorder = originalRecorder
	}()

	// Test direct error recording
	userId := 123
	channelId := 456
	modelName := "gpt-4.1"

	metrics.GlobalRecorder.RecordBillingError("validation_error", "post_consume_detailed", userId, channelId, modelName)

	// Verify billing error was recorded
	if len(mockRecorder.BillingErrors) != 1 {
		t.Errorf("Expected 1 billing error record, got %d", len(mockRecorder.BillingErrors))
	}

	error := mockRecorder.BillingErrors[0]
	if error.ErrorType != "validation_error" {
		t.Errorf("Expected error type 'validation_error', got '%s'", error.ErrorType)
	}
	if error.Operation != "post_consume_detailed" {
		t.Errorf("Expected operation 'post_consume_detailed', got '%s'", error.Operation)
	}
	if error.UserId != userId {
		t.Errorf("Expected userId %d, got %d", userId, error.UserId)
	}
	if error.ChannelId != channelId {
		t.Errorf("Expected channelId %d, got %d", channelId, error.ChannelId)
	}
	if error.ModelName != modelName {
		t.Errorf("Expected modelName '%s', got '%s'", modelName, error.ModelName)
	}
}

func TestBillingTimeoutMonitoring(t *testing.T) {
	// Setup mock metrics recorder
	mockRecorder := &MockMetricsRecorder{}
	originalRecorder := metrics.GlobalRecorder
	metrics.GlobalRecorder = mockRecorder
	defer func() {
		metrics.GlobalRecorder = originalRecorder
	}()

	// Test billing timeout recording
	userId := 123
	channelId := 456
	modelName := "gpt-4.1"
	estimatedQuota := 1500.0
	elapsedTime := 35 * time.Second

	metrics.GlobalRecorder.RecordBillingTimeout(userId, channelId, modelName, estimatedQuota, elapsedTime)

	// Verify billing timeout was recorded
	if len(mockRecorder.BillingTimeouts) != 1 {
		t.Errorf("Expected 1 billing timeout record, got %d", len(mockRecorder.BillingTimeouts))
	}

	timeout := mockRecorder.BillingTimeouts[0]
	if timeout.UserId != userId {
		t.Errorf("Expected userId %d, got %d", userId, timeout.UserId)
	}
	if timeout.ChannelId != channelId {
		t.Errorf("Expected channelId %d, got %d", channelId, timeout.ChannelId)
	}
	if timeout.ModelName != modelName {
		t.Errorf("Expected modelName '%s', got '%s'", modelName, timeout.ModelName)
	}
	if timeout.EstimatedQuota != estimatedQuota {
		t.Errorf("Expected estimatedQuota %f, got %f", estimatedQuota, timeout.EstimatedQuota)
	}
	if timeout.ElapsedTime != elapsedTime {
		t.Errorf("Expected elapsedTime %v, got %v", elapsedTime, timeout.ElapsedTime)
	}
}
