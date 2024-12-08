ENV_FILE=".production.env"

bun tailwindcss -i ./input.css -o ./static/css/style.css
templ generate

rsync -uvrP --exclude ".git/" --delete --delete-excluded ./ pictureframe@pictureframe://home/pictureframe/pictureframe
ssh pictureframe@pictureframe "cd /home/pictureframe/pictureframe && docker compose up --build -d && docker image prune -f"