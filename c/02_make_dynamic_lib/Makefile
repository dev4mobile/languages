all:
	gcc -o main main.c -L. -ltest
lib:
	gcc -c -o test.o test.c -fPIC
	gcc -o libtest.so test.o -shared

run: all
	./main

clean:
	rm main

cleanall:
	rm main
	rm *.o *.so
