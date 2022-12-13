resource "rhoas_kafka" "ai-events-stream-stage" {
  name = "ai-events-stream-stage"
  plan = "standard.x1"
  billing_model = "standard"
}

resource "rhoas_service_account" "svc-account-stage" {
  name        = "svc-account-stage"
  description = "service account for stage usage"

  depends_on = [
    rhoas_kafka.ai-events-stream-stage
  ]
}

resource "rhoas_service_account" "svc-account-integration" {
  name        = "svc-account-integration"
  description = "service account for integration usage"

  depends_on = [
    rhoas_kafka.ai-events-stream-stage
  ]
}

resource "rhoas_topic" "events-stream" {
  name       = "events-stream"
  partitions = 6
  kafka_id   = rhoas_kafka.ai-events-stream-stage.id

  depends_on = [
    rhoas_kafka.ai-events-stream-stage
  ]
}

resource "rhoas_topic" "events-stream-integration" {
  name       = "events-stream-integration"
  partitions = 6
  kafka_id   = rhoas_kafka.ai-events-stream-stage.id

  depends_on = [
    rhoas_kafka.ai-events-stream-stage
  ]
}

output "bootstrap_server_stage" {
  value = rhoas_kafka.ai-events-stream-stage.bootstrap_server_host
}

output "stage_client_id" {
  value = rhoas_service_account.svc-account-stage.client_id
}

output "stage_client_secret" {
  value     = rhoas_service_account.svc-account-stage.client_secret
  sensitive = true
}

output "integration_client_id" {
  value = rhoas_service_account.svc-account-integration.client_id
}

output "integration_client_secret" {
  value     = rhoas_service_account.svc-account-integration.client_secret
  sensitive = true
}
