# 環境変数設定
ROOT=.
MOCK_DIR=./mock

# モック化したいファイルのあるディレクトリ一覧
SEARCH_DIR_LIST=\
internal/domain/repository \
internal/usecase \
internal/infrastructure/auth

# Firebase Auth Clientのインターフェースファイル
FIREBASE_AUTH_INTERFACE=./internal/infra/auth/firebase_auth_client.go

# モック生成のターゲット
.PHONY: mockgen
mockgen:
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
	@if [ -f $(FIREBASE_AUTH_INTERFACE) ]; then \
		mockgen -package mock -source=$(FIREBASE_AUTH_INTERFACE) -destination=$(MOCK_DIR)/mock_firebase_auth_client.go; \
		echo "Generated mock for Firebase Auth Client"; \
	else \
		echo "Firebase Auth Client interface file not found. Skipping."; \
	fi
	@echo "Mock generation completed."

# モックファイルのクリーンアップ
.PHONY: clean-mocks
clean-mocks:
	@echo "Cleaning up mock files..."
	@rm -rf $(MOCK_DIR)
	@echo "Mock cleanup completed."