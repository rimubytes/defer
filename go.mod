// Package defer provides a service with etcd backend and gin web framework
module github.com/rimubytes/defer

// Using Go 1.21 for latest language features and security updates
go 1.21

require (
	// Core etcd dependencies
	github.com/coreos/bbolt v1.3.7         // Embedded key-value database used by etcd
	github.com/coreos/etcd v3.3.27+incompatible       // Distributed key-value store for configuration and service discovery
	github.com/coreos/go-semver v0.3.1     // Semantic versioning for etcd
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf  // SystemD integration
	github.com/coreos/pkg v0.0.0-20230601102743-20bbbf26f4d8  // Common CoreOS packages

	// Authentication and security
	github.com/dgrijalva/jwt-go v3.2.0+incompatible  // JSON Web Token implementation

	// Web framework and routing
	github.com/gin-gonic/gin v1.9.1  // High-performance HTTP web framework

	// Protocol Buffers and gRPC related dependencies
	github.com/gogo/protobuf v1.3.2        // Protocol Buffers for Go with extensions
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect  // Caching library
	github.com/google/btree v1.1.2         // B-Tree implementation used by etcd
	github.com/google/uuid v1.4.0          // UUID generation
	github.com/gorilla/websocket v1.5.1    // WebSocket implementations
	
	// gRPC middleware and monitoring
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0  // gRPC middleware components
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0  // Prometheus monitoring for gRPC
	github.com/grpc-ecosystem/grpc-gateway v1.16.0       // gRPC to JSON proxy generator
	
	// Utility packages
	github.com/jonboulle/clockwork v0.4.0   // Time mocking for testing
	
	// Database drivers
	github.com/mattn/go-sqlite3 v1.14.18               // SQLite3 database driver

	// Monitoring and metrics
	github.com/prometheus/client_golang v1.17.0  // Prometheus monitoring client

	// Network and proxy utilities
	github.com/soheilhy/cmux v0.1.5         // Connection multiplexer
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75  // WebSocket proxy for gRPC
	github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510  // Network probing utility

	// Additional storage and logging
	go.etcd.io/bbolt v1.3.8               // Modern maintained fork of coreos/bbolt
	go.uber.org/zap v1.26.0               // High-performance logging library

	// Core Go packages and extensions
	golang.org/x/crypto v0.15.0           // Cryptographic packages
	golang.org/x/net v0.18.0              // Extended networking functionality
	golang.org/x/time v0.4.0              // Time-related utilities

	// Google API and Protocol Buffer support
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17  // Generated Protocol Buffer code
	google.golang.org/grpc v1.59.0        // gRPC framework
	
	// Kubernetes-related
	sigs.k8s.io/yaml v1.4.0              // YAML support for configuration    
)