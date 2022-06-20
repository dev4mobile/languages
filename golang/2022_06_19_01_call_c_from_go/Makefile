lib:
	# gcc -fPIC -shared greeter.c -o libgreeter.dylib
	gcc -dynamiclib -o libgreeter.dylib greeter.c
run:
	go run main.go

clean:
	rm *.o
	rm *.so
	rm *.dylib
