// Advent of Code 2023, Day 05
//
// Given a list of "seeds" (numbers), go through a series of transformations
// based on tables. Each table is a tuple (dst, src, n). If the number is
// within src .. src+n,  it is assigned to dst. If no dst is found, the dst
// is the same as src for that transformation. For Part 1, find the lowest
// final transformation. Part 2 is the same, but treat each pair of "seeds"
// as start and length of a range. Brute force approach, acceptably fast
// in Go, Rust, Zig, and C.
//
// To build:
//      gcc day05.c -o day05      (debug mode)
//      gcc -O3 day05.c -o day05  (fast mode, takes 2:40 mins)
//
// To run:
//      ./day05
//
// AK, 12 Oct 2024 for C version

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <stdint.h>

// These must be set big enough for the data, or strange things will occur
#define MAX_SEEDS 128   // set to maximum number of seeds
#define MAX_SECTIONS 32 // set to maximum number of sections
#define MAX_RANGES 128  // set to maximum number of ranges within a sections

// Data structures for a Section containing multiple Ranges
typedef struct {
    long a, b, c;
} Range;

// A section is a sequence of ranges
typedef struct {
    char *name;
    int nranges;
    Range *ranges[MAX_RANGES];
} Section;

// Functions declared below
int find_char(char *s, char c);
void rtrim(char *s);
int parse_nums(char *s, long *nums);
long process_seed(long seed, Section *sections[], int nsections);

// Main execution
void main()
{
    // Seeds data, list of numbers (initialize to zero)
    long seeds[MAX_SEEDS];
    int nseeds = 0;
    for ( int i = 0; i < MAX_SEEDS; ++i )
        seeds[i] = 0;

    // Sections data
    Section *sections[MAX_SECTIONS];
    int nsections = 0;

    // Read data: list of numbers on first line,
    // the "sections" with 3 numbers on each line
    FILE *f = fopen("input.txt", "r");  // open text file
    char buf[512];                      // buffer for reading each line
    long nums[32];                      // only 3 numbers per range
    Section *curSection = NULL;         // section currently being parsed
    int line_no = 0;                    // file line number for messages
    while ( fgets(buf, 512, f) ) {      // read each line

        // Seeds on first line
        line_no += 1;
        rtrim(buf);                     // remove trailing spaces
        if ( line_no == 1 ) {
            int colon = find_char(buf, ':');
            nseeds = parse_nums(buf+colon+2, seeds);
        }

        // Line starting a new section has a colon
        else if ( find_char(buf, ':') > 0 ) {
            curSection = malloc(sizeof(Section));
            curSection->name = strdup(buf);
            curSection->nranges = 0;
            sections[nsections++] = curSection;
        }

        // Otherwise a range of numbers within the current section
        else if ( strlen(buf) > 0 ) {
            if ( parse_nums(buf, nums) == 3 ) {
                Range *r = malloc(sizeof(Range));
                r->a = nums[0];
                r->b = nums[1];
                r->c = nums[2];
                curSection->ranges[curSection->nranges] = r;
                ++curSection->nranges;
            } else {
                printf("Line %d does not contain 3 numbers: %s\n", line_no, buf);
            }
        }
    }

    // Show seeds found
    printf("Seeds:");
    for ( int i = 0; seeds[i] != 0; ++i )
        printf(" %ld", seeds[i]);
    puts("");

    // Show sections
    printf("Sections:\n");
    for ( int i = 0; i < nsections; ++i ) {
        Section *s = sections[i];
        printf("  %s (%d ranges)\n", s->name, s->nranges);
    }

	// Part 1: process each seed, find the lowest location
	long lowest = -1;
	for ( int i = 0; i < nseeds; i++ ) {
		long loc = process_seed(seeds[i], sections, nsections);
		if ( lowest == -1 || loc < lowest ) {
			lowest = loc;
		}
	}
	printf("Part 1 (s/b 424490994): %ld\n", lowest);

	// Part 2: same, but treat seed numbers as ranges
	lowest = -1;
	for ( int i = 0; i < nseeds; i += 2 ) {
		long s1 = seeds[i],     // start of range
		    n = seeds[i+1];     // length of range
		printf("Seed %ld for length %ld\n", s1, n);
		for ( long j = 0; j < n; ++j ) {
			long loc = process_seed(s1 + j, sections, nsections);
			if ( lowest == -1 || loc < lowest ) {
				lowest = loc;
			}
		}
	}
	printf("Part 2 (s/b 15290096): %ld\n", lowest);

    // TODO: free memory
}

// Process a seed, by traversing the transformations and getting
// the end location
long process_seed(long seed, Section *sections[], int nsections)  {

    // Go through each section
    long src = seed;
    for ( int i = 0; i < nsections; ++i ) {

        // Find the destination corresponding to the current source
        Section *sec = sections[i];
        long dst = 0;
        int found = 0; // false
        for ( int j = 0; j < sec->nranges; ++j ) {
            Range *r = sec->ranges[j];
            if ( src >= r->b && src <= r->b + r->c ) {
                dst = r->a + src - r->b;
                found = 1;
                break;
            }
        }

        // If not found, use the source value
        if ( !found ) {
            dst = src;
        }

        // Source of next transformation is the destination of this one
        src = dst;
    }

    // Return the final destination
    return src;
}

// Find character in a string, return index or -1
int find_char(char *s, char c) {
    for ( int i = 0; i < strlen(s); ++i ) {
        if ( s[i] == c )
            return i;
    }
    return -1;
}

// Right-trim a string
void rtrim(char *s) {
    char *p;
    for ( p = s + strlen(s) - 1; (p >= s) && isspace(*p); --p ) {
        *p = '\0';
    }
}

// Parse a line of space-separated of long integers, return count found
int parse_nums(char *s, long *nums) {

    int ni = 0;         // the current number being created
    int i = 0;          // index within the string
    while ( 1 ) {       //true

        // Skip leading spaces, exit if end of string or not a digit
        while ( s[i] == ' ' )
            ++i;
        if ( !isdigit(s[i]) )
            return ni;

        // parse digits of a number
        nums[ni] = 0LL;
        while ( isdigit(s[i]) ) {
            nums[ni] *= 10LL;
            nums[ni] += s[i] - '0';
            ++i;
        }

        // Ready for next number
        ++ni;
    }
}

