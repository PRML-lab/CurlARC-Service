# 環境変数設定
ROOT=.
MOCK_DATA_ROOT=./mock

# ROOTのモック化したいファイルのあるディレクトリ一覧
SEARCH_DIR_LIST=\
internal/domain/repository \
internal/usecase \

# モック生成のターゲット
mockgen:
	@echo "Start..."
	@cd $(ROOT) && \
	for dir in $(SEARCH_DIR_LIST); do \
	  file_path_list=$$(find ./$$dir -type f -not -name "*_test.go" -name "*.go"); \
	  for file_path in $$file_path_list; do \
	    base_name=$$(basename $$file_path); \
	    mockgen -package mock -source=./$$file_path -destination=./$(MOCK_DATA_ROOT)/$$dir/mock_$$base_name; \
	  done \
	done
	@echo "Done."