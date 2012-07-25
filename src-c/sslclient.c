#include <stdio.h>
#include <string.h>
#include <openssl/bio.h>
#include <openssl/ssl.h>
#include <openssl/err.h>

typedef unsigned long SSL_ERR;

int main(int argc, char *argv[]) {
   BIO *bio;
   SSL *ssl;
   SSL_ERR err;

   /*
    * TODO everything is static/hard coded right now
    */
   char *sendstr = "HEAD / HTTP/1.1\r\nHost: google.com\r\n\r\n";
   char recvbuf[8192];
   int nbytes;

   SSL_library_init();
   SSL_load_error_strings();
   ERR_load_BIO_strings();
   OpenSSL_add_all_algorithms();

   SSL_CTX *ctx = SSL_CTX_new(SSLv23_client_method());

   if(!ctx) {
      err = ERR_get_error();
      printf("%s\n", ERR_error_string(err, NULL));
      return 1;
   }

   bio = BIO_new_ssl_connect(ctx);
   BIO_get_ssl(bio, &ssl);
   SSL_set_mode(ssl, SSL_MODE_AUTO_RETRY);

   BIO_set_conn_hostname(bio, "google.com");
   BIO_set_conn_port(bio, "443");

   if(BIO_do_connect(bio) <= 0) {
      err = ERR_get_error();
      printf("%s\n", ERR_error_string(err, NULL));
      return 1;
   }

   /*
    * TODO flesh this out
    */
   BIO_write(bio, sendstr, strlen(sendstr));
   nbytes = BIO_read(bio, recvbuf, sizeof(recvbuf));
   printf("%s\n", recvbuf);

   BIO_free_all(bio);
   return 0;
}
