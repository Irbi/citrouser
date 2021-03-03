PID_FILE = /tmp/citro-user-app.pid

serve: start
	fswatch -or --event=Updated . | xargs -n1 -I {} make restart

start:
	go run main.go & echo $$! > $(PID_FILE)

stop:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

before:
	@echo "STOPPED citro-user-app" && printf '%*s\n' "40" '' | tr ' ' -

restart: stop before start
	@echo "STARTED citro-user-app" && printf '%*s\n' "40" '' | tr ' ' -

.PHONY: start before stop restart serve