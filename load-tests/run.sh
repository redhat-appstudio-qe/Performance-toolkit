export DOCKER_CONFIG_JSON=

export GITHUB_E2E_ORGANIZATION=app-studio-test

export GITHUB_TOKEN=




if [ -z ${DOCKER_CONFIG_JSON+x} ]; then echo "env DOCKER_CONFIG_JSON need to be defined"; exit 1;  else echo "DOCKER_CONFIG_JSON is set"; fi

go test -test.v -test.run ^TestFeatures$