/*******************************************************************************
 * So, a buddy of mine asked me how to write a c program to determine if a stack
 * grows up or down.  In a very hand wavey explanation I told him to "just
 * compare the address of a local variable to the stack address when the
 * variable is passed to some procedure as an argument".  This is a really basic
 * demonstration... nothing fancy.
 ******************************************************************************/
#include <stdio.h>
#include <sys/types.h>

void proc(int arg) {
   int localvar;

   /*
    * Simply print the address of some local (to me) variable along along with
    * the address of a passed parameter.  This will tell us, relative to main,
    * which direction the stack is growing.
    */
   printf("proc() local variable: 0x%08zx\n", (size_t)&localvar);
   printf("proc() stack variable: 0x%08zx\n", (size_t)&arg);
}

int main(int argc, char *argv[]) {
   int localvar;

   printf("main() local variable: 0x%08zx\n", (size_t)&localvar);

   proc(localvar);

   return 0;
}
