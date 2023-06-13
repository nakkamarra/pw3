INSTALL_PATH := "/usr/local/bin"
pw3: clean
	go build
clean:
	rm -f pw3
uninstall:
	sudo rm -f $(INSTALL_PATH)/pw3
install: uninstall pw3
	sudo install -s pw3 $(INSTALL_PATH)
