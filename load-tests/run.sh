export DOCKER_CONFIG_JSON=ewogICJhdXRocyI6IHsKICAgICJxdWF5LmlvIjogewogICAgICAiYXV0aCI6ICJiWE5oZDI5dlpEcHBVelo0T0doMlUyOXhhV3AwYW5SYVJsSlBjV1JwWkVzek16QjVObGR2UlRoSk1YZGpWbXRUVDJWbkswOHlkRk5OTVZGckswWjFkalJ1VTA1VVZYbEoiLAogICAgICAiZW1haWwiOiAiIgogICAgfQogIH0KfQ==

export GITHUB_E2E_ORGANIZATION=app-studio-test

export GITHUB_TOKEN=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


if [ -z ${DOCKER_CONFIG_JSON+x} ]; then echo "env DOCKER_CONFIG_JSON need to be defined"; exit 1;  else echo "DOCKER_CONFIG_JSON is set"; fi

go test -test.v -test.run ^TestFeatures$