OPENAPI_SPEC_OUTDIR=web-client/docs
OPENAPI_SPEC_PATH=${OPENAPI_SPEC_OUTDIR}/openapi3-spec.yml

.PHONY: help
help:
	echo "see Makefile"

.PHONY: generate_openapi_typescript
generate_openapi_typescript:
	mkdir -p ${OPENAPI_SPEC_OUTDIR}
	go run projects-app-main.go generate-openapi-spec --output ${OPENAPI_SPEC_PATH}
	sed -i 's/Domain//g' ${OPENAPI_SPEC_PATH}
	sed -i 's/Webservices//g' ${OPENAPI_SPEC_PATH}
	cd web-client && mkdir -p src/openapi/generated && rm -f src/openapi/generated/* && echo "running yarn generate-openapi" && yarn generate-openapi

.PHONY: run_dev_server
run_dev_server:
	go run projects-app-main.go serve ~/src

.PHONY: run_dev_client
run_dev_client:
	cd web-client && yarn dev