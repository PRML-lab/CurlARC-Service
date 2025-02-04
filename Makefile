# 環境変数設定
ROOT=.
MOCK_DIR=./mock

# モック化したいファイルのあるディレクトリ一覧
SEARCH_DIR_LIST=\
internal/domain/repository \
internal/domain/auth \
internal/usecase

# モック生成のターゲット
.PHONY: mockgen
mockgen:
	make clean-mocks
	@echo "Starting mock generation..."
	@mkdir -p $(MOCK_DIR)
	@cd $(ROOT) && \
	for dir in $(SEARCH_DIR_LIST); do \
		file_path_list=$$(find ./$$dir -type f -not -name "*_test.go" -name "*.go"); \
		for file_path in $$file_path_list; do \
			base_name=$$(basename $$file_path .go); \
			package_name=$$(basename $$(dirname $$file_path)); \
			mockgen -package mock -source=$$file_path -destination=$(MOCK_DIR)/mock_$${package_name}_$${base_name}.go; \
			echo "Generated mock for $$file_path"; \
		done \
	done
	@echo "Mock generation completed."

# モックファイルのクリーンアップ
.PHONY: clean-mocks
clean-mocks:
	@echo "Cleaning up mock files..."
	@rm -rf $(MOCK_DIR)
	@echo "Mock cleanup completed."


test:
	gotest -v ./...

# Variables
DB_URL="postgres://postgres:EbmU6Q0LbRbe0LV@localhost:5432?sslmode=disable"
MIGRATIONS_DIR=migrations


# Apply the migration in remote to local database
migrate-apply:
	@migrate -database $(DB_URL) -path $(MIGRATIONS_DIR) up
