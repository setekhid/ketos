#!/bin/bash

TEST_ROOT="$( cd "$(dirname "${BASH_SOURCE[0]}")" && pwd )"

(cd "${TEST_ROOT}" && docker-compose up -d)
function clean_docker_ps {
	(cd "${TEST_ROOT}" && docker-compose down)
}
trap clean_docker_ps EXIT
sleep 5

docker exec -it test_ketos_1 bash /test/test_ketos.sh
