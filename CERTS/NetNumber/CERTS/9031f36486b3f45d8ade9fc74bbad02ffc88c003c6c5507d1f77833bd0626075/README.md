# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Baltimore-Washington Telephone Company SHAKEN cert 8697

Tested At: 22 Aug 24 15:16 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 18 day(s)\
Subject: O=Baltimore-Washington Telephone Company, C=US, CN=Baltimore-Washington Telephone Company SHAKEN cert 8697\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://bwt.io/8697a2.crt

[View certificate details](https://x509.io/?cert=MIICvjCCAkOgAwIBAgIJAKuTbaA2ti9gMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yNDA4MDkxNTAwMDBaFw0yNDA5MDkxNDU5NTlaMIGAMUAwPgYDVQQDDDdCYWx0aW1vcmUtV2FzaGluZ3RvbiBUZWxlcGhvbmUgQ29tcGFueSBTSEFLRU4gY2VydCA4Njk3MQswCQYDVQQGEwJVUzEvMC0GA1UECgwmQmFsdGltb3JlLVdhc2hpbmd0b24gVGVsZXBob25lIENvbXBhbnkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR7PqgstTJ0GphOuUb64MPJkWWX6Nd%2FNz9kYzpj%2BjtDTPIPk2EMAzGCSNNBFuRF0obuipQkJf6mYYHRZPSkS%2BqBo4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQ4Njk3MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBS8CFGQU1KbdE7xcbVIUHRr4NPURTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaQAwZgIxANbBM0kfVhd%2Bugdd5CbMxELMkdlp0U%2BV4%2FqzZPfYcudAsZHfl2xSh6Hyatevb5GcfQIxAJbCzNMYSglN8ZVCCorEHHO26TGUK7W9SSVkgi8ZzRUKGy%2BsblSdwiKey4anfVIHhQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8697', but common name is 'Baltimore-Washington Telephone Company SHAKEN cert 8697' |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |


Generated: 22 Aug 24 15:44 UTC