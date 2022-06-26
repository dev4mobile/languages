all:
	@cc -o main main.c -Wall -I/usr/include -L/usr/lib -lpcre2-8 

run: all
	@./main

clean:
	rm main

cleanall:
	rm main
	rm *.o *.so
