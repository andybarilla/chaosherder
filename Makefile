.PHONY: watch-css
watch-css:
	 npx @tailwindcss/cli -i tailwind.css -o internal/server/public/styles/app.css --watch