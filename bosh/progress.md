
## [Documentation](https://bosh.io/docs)

### Introduction

 - [x] What is BOSH?
	 - [x] What problems does BOSH solve?
		 - [x] Stemcell
		 - [x] Release
		 - [x] Deployment
	 - ~~Comparison to other tools~~
 - [x] Terminology

### Install BOSH

#### Install BOSH with bosh-init

 - [x] BOSH components
 - [ ] Install bosh-init
     - [ ] Using bosh-init
         - [ ] Migrating to bosh-init from the micro CLI plugin
 - [x] Install BOSH CLI
 - [x] Bootstrapping an environment
	 - [ ] On AWS
	 - ~~On OpenStack~~
	 - ~~On vSphere~~
	 - ~~On vCloud~~
	 - [x] On Local machine using BOSH Lite

 - [ ] Director SSL Certificate Configuration
 - [ ] Using bosh-init
 	 - [ ] Migrating to bosh-init from the micro CLI plugin
 - [ ] Install BOSH CLI

#### Install BOSH with the micro CLI plugin [deprecated, but supported]

 - [ ] Bootstrapping an environment
 	 - [ ] On AWS
 	 - ~~On OpenStack~~
 	 - ~~On vSphere~~

#### Advanced Director configuration

 - [ ] User management
	 - [ ] UAA Integration
 - [ ] Configuring External Database
 - [ ] Configuring External Blobstore
 - [ ] Troubleshooting

### Using BOSH to deploy software

 - [ ] Basic workflow
	 - [ ] Deployment basics
		 - [ ] Deployment manifest schema
	 - [ ] Uploading stemcells
	 - [ ] Uploading releases
	 - [ ] Deploying
	 - [ ] Running one off tasks
	 - [ ] Updating deployment to deal with security vulnerabilities
 - [ ] Deploying step-by-step
 - [ ] CLI Commands
 - [ ] Director tasks

#### Detailed Deployment Configuration

 - [ ] Deployment Jobs
 - [ ] Resource Pools
	 - [ ] VM anti-affinity
 - [ ] Networks
 - [ ] Persistent disks
 - [ ] Trusted certificates
 - [ ] IaaS specifics
	 - [ ] AWS
		 - [ ] Using IAM instance profiles
	 - [ ] Using instance storage
	 - ~~OpenStack~~
	 - ~~vSphere~~
	 - ~~vCloud~~
	 - ~~Azure [beta]~~
	 - [ ] Warden/Garden

#### Health Management of VMs and Processes

 - [ ] Monitoring
	 - [ ] Configuring Health Monitor
 - [ ] Process monitoring with Monit
 - [ ] Manual repair with Cloud Check
 - [ ] Automatic repair with Resurrector
 - [ ] Persistent disk snapshotting

#### VM Configuration (Looking inside a Deployment)

 - [ ] Structure of a BOSH VM
	 - [ ] VM Configuration Locations
 - [ ] Location and use of logs
 - [ ] Debugging issues with jobs

### Using BOSH to package and distribute software

 - [ ] What is a release?
 	 - [ ] Creating a release
	 - [ ] Testing with dev releases
	 - [ ] Cutting final releases
		 - [ ] Versioning of releases
 - [ ] What is a job?
	 - [ ] Creating a job
	 - [ ] Job properties
	 - [ ] Job lifecycle
		 - [ ] Pre-start script
		 - [ ] Drain script
 - [ ] What is a package?
	 - [ ] Creating a package
	 - [ ] Relationship to release blobs
 - [ ] How do releases, jobs, and packages interact?
 - [ ] Managing release repository
	 - [ ] Release blobstore
		 - [ ] Configuring S3 release blobstore

### Extending BOSH to support other IaaSs

 - [ ] What is a CPI?
 - [ ] Building a CPI
	 - [ ] CPI API v1
	 - [ ] Interactions between CPIs and BOSH Agent
 - [ ] Building a stemcell
