.PHONY: run-elastic run-kibana

run-elastic:
	@echo "  >  Create elasticsearch container...\n"
	@docker run -p 9200:9200 -d --name elasticsearch \
		-e "discovery.type=single-node" \
		-e "xpack.security.enabled=false" \
		-e "xpack.security.http.ssl.enabled=false" \
		-e "xpack.license.self_generated.type=trial" \
		docker.elastic.co/elasticsearch/elasticsearch:8.13.0

run-kibana:
	@echo "  >  Create kibana container...\n"
	@docker run -p 5601:5601 -d --name kib01 docker.elastic.co/kibana/kibana:8.13.4
