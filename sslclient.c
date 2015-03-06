/*******************************************************************************
* This is a basic ssl client example.  This program needs to be linked against
* the and libssl (available via openssl).  We send an HTTP HEAD request and
* display the headers sent back.
*
* gcc -Wall -o ex sslclient.c -lssl
*******************************************************************************/
#include <openssl/bio.h>
#include <openssl/ssl.h>
#include <openssl/err.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

typedef unsigned long SSL_ERR;
#define BUFFSZ 8192

int main(int argc, char **argv) {
   BIO *bio;
   SSL *ssl;
   SSL_ERR err;
   char c;
   char *host = "google.com";
   char *port = "443";

   char *sendtemplate = "HEAD / HTTP/1.1\r\nHost: %s\r\n\r\n";
   char sendbuf[BUFFSZ];
   char recvbuf[BUFFSZ];
   int nbytes;

   SSL_library_init();
   SSL_load_error_strings();
   ERR_load_BIO_strings();
   OpenSSL_add_all_algorithms();

   SSL_CTX *ctx = SSL_CTX_new(SSLv23_client_method());

   while((c = getopt(argc, argv, "h:p:")) != -1) {
      switch(c) {
      case 'h':
         host = optarg;
         break;
      case 'p':
         port = optarg;
         break;
      case '?':
         return 1;
         break;
      }
   }

   snprintf(sendbuf, BUFFSZ, sendtemplate, host);

   if(!ctx) {
      err = ERR_get_error();
      printf("%s\n", ERR_error_string(err, NULL));
      return 1;
   }

   bio = BIO_new_ssl_connect(ctx);
   BIO_get_ssl(bio, &ssl);
   SSL_set_mode(ssl, SSL_MODE_AUTO_RETRY);

   BIO_set_conn_hostname(bio, host);
   BIO_set_conn_port(bio, port);

   if(BIO_do_connect(bio) <= 0) {
      err = ERR_get_error();
      printf("%s\n", ERR_error_string(err, NULL));
      return 1;
   }

   /*
    * TODO flesh this out
    * send request, and print the response
    */
   BIO_write(bio, sendbuf, strlen(sendbuf));
   nbytes = BIO_read(bio, recvbuf, sizeof(recvbuf));
   printf("%s\n", recvbuf);

   BIO_free_all(bio);
   return 0;
}
