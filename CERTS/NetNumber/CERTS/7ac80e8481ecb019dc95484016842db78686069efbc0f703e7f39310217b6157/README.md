# STIR/SHAKEN CA Ecosystem Compliance

## Certificate NetNumber SHAKEN Root CA

Tested At: 01 Nov 22 10:05 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 6824 day(s)\
Subject: L=Lowell, ST=Massachusettes, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA\
Issuer: L=Lowell, ST=Massachusettes, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA

View: [Click to view](https://understandingwebpki.com/?cert=MIICqzCCAlGgAwIBAgIIJfiyoyDN6lQwCgYIKoZIzj0EAwIwcjEhMB8GA1UEAwwYTmV0TnVtYmVyIFNIQUtFTiBSb290IENBMQswCQYDVQQGEwJVUzEWMBQGA1UECgwNTmV0TnVtYmVyIEluYzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMTA3MTIyMzI1MTdaFw00MTA3MDcyMzI1MTdaMHIxITAfBgNVBAMMGE5ldE51bWJlciBTSEFLRU4gUm9vdCBDQTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxFzAVBgNVBAgMDk1hc3NhY2h1c2V0dGVzMQ8wDQYDVQQHDAZMb3dlbGwwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2FqKl0J%2FwfOHoXROKL9%2FSewIzTxTdfrm1t6sGL6N%2Bn6RbHRbCVKXeJzqdoSQCgTirWy8cTqcM50eRKV%2FW48QCEo4HQMIHNMBIGA1UdEwEB%2FwQIMAYBAf8CAQMwDgYDVR0PAQH%2FBAQDAgIEMB8GA1UdIwQYMBaAFI9DG8y%2BjTlzGNbIaOoUvosF1SQ3MB0GA1UdDgQWBBSPQxvMvo05cxjWyGjqFL6LBdUkNzBLBgNVHR8ERDBCMECgPqA8hjpodHRwczovL2F1dGhlbnRpY2F0ZS1hcGktc3RnLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNIADBFAiEAwP6jpZnv%2FKgyBqjayLOwa4tZ3yCxYjOZvZ0VKuzd85MCIFlgIiYiEnUzHRNOpG7eIrPsxYeRw7TYL%2Bj%2BIONUGa2m)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_ca_subject_rdn_unknown](../../ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_serial_number
- e_sti_ca_subject_cn
- e_sti_root_certificate_policies
- e_sti_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 10:05:32