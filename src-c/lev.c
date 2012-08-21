/*******************************************************************************
* levenshtein (edit) distance; O(n^2) for this approach :C
*
* TODO - this code is a bit "magic numberey" in a few places, clean that up.
*******************************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int min3(int a, int b, int c) {
   return (a < b) ? (a < c) ? a : c : (b < c) ? b : c;
}

/*
 * this silly dump routine just prints the lev matrix
 */
void dump(int **mat, int m, int n, char *s1, char *s2) {
   int i,j;

   /*
    * print nice header
    */
   printf("% 4s", "");
   for(j=0; j<n; j++) {
      printf("% 3c", j ? s2[j-1] : ' ');
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
      printf("% 3c|", i ? s1[i-1] : ' ');
      for(j=0; j<n; j++) {
         printf("% 3d", mat[i][j]);
      }
      printf("\n");
   }
}

int lev(char *s1, char *s2) {
   int i,j;
   int m = strlen(s1) + 1;
   int n = strlen(s2) + 1;

   int **mat;
   mat = malloc(m * sizeof(int*));
   for(i=0; i<m; i++) {
      mat[i] = malloc(n * sizeof(int));
   }


   for(i=0; i<m; i++) {
      mat[i][0] = i;
   }
   for(j=0; j<n; j++) {
      mat[0][j] = j;
   }

   /*
    * dynamic-programming
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
   if(argc != 3)
      printf("Usage: lev string1 string2\n");
   else
      lev(argv[1], argv[2]);
   return 0;
}
