lib:
	# gcc -fPIC -shared greeter.c -o libgreeter.dylib
	gcc -fPIC -shared greeter.c -o libgreeter.so
	# gcc -dynamiclib -o libgreeter.dylib greeter.c
	# gcc -c greeter.c
run:
	go run main.go

clean:
	rm *.o
	rm *.so
	rm *.dylib
