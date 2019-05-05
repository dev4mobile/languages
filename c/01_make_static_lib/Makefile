all:
	gcc main.c -o main -L. -ltest

lib:
	gcc -o test.o test.c -c
	ar -rc libtest.a test.o

run: all
	./main

clean:
	rm main

cleanall:
	rm main
	rm *.o *.a
