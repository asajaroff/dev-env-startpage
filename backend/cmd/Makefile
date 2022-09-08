.PHONY= test

URL := localhost
PORT := 2525

test:
	# Create a cluster
	curl -X POST http://${URL}:${PORT}/v1/clusters/new -d 'vcluster-name=alejandro-devstarter'
	# List all clusters
	curl -X GET http://${URL}:${PORT}/v1/clusters/list 

create: 
	# Create a cluster
	curl -X POST http://${URL}:${PORT}/v1/clusters/new -d 'vcluster-name=cmake-automated-01'

destroy: 
	curl -X DELETE http://${URL}:${PORT}/v1/clusters/new -d 'vcluster-name=cmake-automated-01'
