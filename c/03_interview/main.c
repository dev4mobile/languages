#define PCRE2_STATIC
#define PCRE2_CODE_UNIT_WIDTH 8

#include <stdio.h>
#include <pcre2.h>
#include <string.h>
#include "pcre2posix.h"

int main() {
	int errornumber;
	PCRE2_SIZE erroroffset;
	char* str = "(?<=\\d{4})([^\\d\\s]){3,11}(?=\\S)";
	PCRE2_SPTR pattern = (PCRE2_SPTR)str; 
	pcre2_code *re = pcre2_compile(pattern, PCRE2_ZERO_TERMINATED, 0, &errornumber, &erroroffset, NULL);
	if (re == NULL) {
		PCRE2_UCHAR buffer[256];
  		pcre2_get_error_message(errornumber, buffer, sizeof(buffer));
		printf("PCRE2 compilation failed at offset %d: %s\n", (int)erroroffset, buffer);
		return 1;
	}
	pcre2_match_data *match_data = pcre2_match_data_create_from_pattern(re, NULL);
	PCRE2_SPTR subject = (PCRE2_SPTR)"a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999";
	// size_t subject_length;
	size_t subject_length = strlen((char *) subject);
// 	int rc = pcre2_match(
//   	re,                   /* the compiled pattern */
//   	subject,              /* the subject string */
//   	subject_length,       /* the length of the subject */
//   	0,                    /* start at offset 0 in the subject */
//   	0,                    /* default options */
//   	match_data,           /* block for storing the result */
//   	NULL);                /* use default match context */

// 	if (rc < 0)
//   	{
//   switch(rc)
//     {
//     case PCRE2_ERROR_NOMATCH: printf("No match\n"); break;
//     /*
//     Handle other special cases if you like
//     */
//     default: printf("Matching error %d\n", rc); break;
//     }
//   	pcre2_match_data_free(match_data);   /* Release memory used for the match */
//   	pcre2_code_free(re);                 /* data and the compiled pattern. */
//   return 1;
//   }
//   PCRE2_SIZE *ovector = pcre2_get_ovector_pointer(match_data);
// 	printf("\nMatch succeeded at offset %d\n", (int)ovector[0]);
// 	printf("success\n");
	int rc = 0;
	int start_offset = 0;
	while((rc = pcre2_match(re, subject, subject_length, start_offset, 0, match_data, NULL)) > 0) {
		PCRE2_SIZE *ovector = pcre2_get_ovector_pointer(match_data);
		int i = 0;
        for (i = 0; i < rc; i++) {
			PCRE2_SPTR substring_start = subject + ovector[2*i];
			size_t substring_length = ovector[2*i+1] - ovector[2*i];
			printf("matched %2d: %.*s\n", i, (int)substring_length, (char *)substring_start);
		}
		 start_offset = ovector[2*(i-1) + 1];
	}
	return 0;
}
