MAIN_PKG = src/cmd/main.go

BUILD_DIR = dist

CONFIG_PATH = config.yaml

serve:
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}
	${BUILD_DIR}/main.exe -config=${CONFIG_PATH}

build: 
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}

swag:
	swag init --parseDependency --parseInternal -g src/internal/app/app.go -o src/docs

.PHONY: swag