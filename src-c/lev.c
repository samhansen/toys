/*******************************************************************************
* Levenshtein (edit) distance; O(n^2) for this approach :C  This was mostly an
* exercise to solidify pointer manipulation in c.
*
* Notes -
*  (1) the min function is a bit needlessly complicated.  I found some puzzle
*      regarding implementing min/max without the use of if statements and
*      decided to take a whack at that here.
*  (2) the Levenshtein distance considers increasingly long string prefixes.
*      The null string is considered a valid prefix (represented by a space).
*
*******************************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*
 * min function - in practice this should be a macro or something less obscure,
 * see notes.
 */
int min3(int a, int b, int c) {
   return (a < b) ? (a < c) ? a : c : (b < c) ? b : c;
}



/*
 * this silly dump routine just prints the lev matrix
 */
void dump(int **mat, int m, int n, char *s1, char *s2) {
   int i,j;

   /*
    * print a nice header
    */
   printf("% 4s", "");
   for(j=0; j<n; j++) {
      printf("% 3c", s2[j]);
   }
   printf("\n");
   printf("% 4s", "");
   for(j=0; j<n; j++) {
      printf("---");
   }
   printf("\n");

   /*
    * the meat of dump - what you might expect
    */
   for(i=0; i<m; i++) {
      printf("% 3c|", s1[i]);
      for(j=0; j<n; j++) {
         printf("% 3d", mat[i][j]);
      }
      printf("\n");
   }
}



/*
 * compute the Levenshtein distance
 */
int lev(char *s1, char *s2) {
   int i,j;
   int m = strlen(s1);
   int n = strlen(s2);

   /*
    * standard cruft to dynamically allocate a 2d array
    */
   int **mat;
   mat = malloc(m * sizeof(int*));
   for(i=0; i<m; i++) {
      mat[i] = malloc(n * sizeof(int));
   }

   /*
    * pre-compute what we already know to be true
    */
   for(i=0; i<m; i++) {
      mat[i][0] = i;
   }
   for(j=0; j<n; j++) {
      mat[0][j] = j;
   }

   /*
    * the meat of computing the Levenshtein distance
    */
   for(j=1; j<n; j++) {
      for(i=1; i<m; i++) {
         if(s1[i-1] == s2[j-1])
            mat[i][j] = mat[i-1][j-1];
         else
            mat[i][j] = min3(mat[i-1][j] + 1,
                             mat[i][j-1] + 1,
                             mat[i-1][j-1] + 1);
      }
   }

   dump(mat,m,n,s1,s2);

   return mat[m-1][n-1];
}


int main(int argc, char **argv) {
   char *s1, *s2;
   int lens1, lens2;
   int distance;

   if(argc != 3)
      printf("Usage: lev string1 string2\n");
   else {
      /*
       * we need 2 additional characters - one for the terminating null, and
       * one for the leading blank (actually a space) character.  See notes.
       */
      lens1 = 2 + strlen(argv[1]);
      lens2 = 2 + strlen(argv[2]);
      s1 = malloc(lens1 * sizeof(char));
      s2 = malloc(lens2 * sizeof(char));

      /*
       * generate the strings actually used in the lev computation
       */
      snprintf(s1, lens1, " %s", argv[1]);
      snprintf(s2, lens2, " %s", argv[2]);

      distance = lev(s1, s2);
      printf("Levenshtein distance: %d\n", distance);
   }

   return 0;
}
