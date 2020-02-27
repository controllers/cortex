/*
Copyright 2020 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package userconfig

const (
	// API
	NameKey           = "name"
	EndpointKey       = "endpoint"
	PredictorKey      = "predictor"
	TrackerKey        = "tracker"
	ComputeKey        = "compute"
	AutoscalingKey    = "autoscaling"
	UpdateStrategyKey = "update_strategy"

	// Predictor
	TypeKey         = "type"
	PathKey         = "path"
	ModelKey        = "model"
	PythonPathKey   = "python_path"
	ConfigKey       = "config"
	EnvKey          = "env"
	SignatureKeyKey = "signature_key"

	// Tracker
	KeyKey       = "key"
	ModelTypeKey = "model_type"

	// Compute
	CPUKey = "cpu"
	MemKey = "mem"
	GPUKey = "gpu"

	// Autoscaling
	MinReplicasKey                  = "min_replicas"
	MaxReplicasKey                  = "max_replicas"
	InitReplicasKey                 = "init_replicas"
	WorkersPerReplicaKey            = "workers_per_replica"
	ThreadsPerWorkerKey             = "threads_per_worker"
	TargetQueueLengthKey            = "target_queue_length"
	WindowKey                       = "window"
	DownscaleStabilizationPeriodKey = "downscale_stabilization_period"
	UpscaleStabilizationPeriodKey   = "upscale_stabilization_period"
	MaxDownscaleFactorKey           = "max_downscale_factor"
	MaxUpscaleFactorKey             = "max_upscale_factor"
	DownscaleToleranceKey           = "downscale_tolerance"
	UpscaleToleranceKey             = "upscale_tolerance"

	// UpdateStrategy
	MaxSurgeKey       = "max_surge"
	MaxUnavailableKey = "max_unavailable"

	// K8s annotation
	MinReplicasAnnotationKey                  = "autoscaling.cortex.dev/min-replicas"
	MaxReplicasAnnotationKey                  = "autoscaling.cortex.dev/max-replicas"
	WorkersPerReplicaAnnotationKey            = "autoscaling.cortex.dev/workers-per-replica"
	ThreadsPerWorkerAnnotationKey             = "autoscaling.cortex.dev/threads-per-worker"
	TargetQueueLengthAnnotationKey            = "autoscaling.cortex.dev/target-queue-length"
	WindowAnnotationKey                       = "autoscaling.cortex.dev/window"
	DownscaleStabilizationPeriodAnnotationKey = "autoscaling.cortex.dev/downscale-stabilization-period"
	UpscaleStabilizationPeriodAnnotationKey   = "autoscaling.cortex.dev/upscale-stabilization-period"
	MaxDownscaleFactorAnnotationKey           = "autoscaling.cortex.dev/max-downscale-factor"
	MaxUpscaleFactorAnnotationKey             = "autoscaling.cortex.dev/max-upscale-factor"
	DownscaleToleranceAnnotationKey           = "autoscaling.cortex.dev/downscale-tolerance"
	UpscaleToleranceAnnotationKey             = "autoscaling.cortex.dev/upscale-tolerance"
)